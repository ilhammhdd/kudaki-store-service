package event_driven

import (
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/ilhammhdd/kudaki-store-service/externals/mysql"

	"github.com/ilhammhdd/go-toolkit/safekit"

	"github.com/ilhammhdd/go-toolkit/errorkit"

	"github.com/ilhammhdd/kudaki-store-service/adapters"

	"github.com/google/uuid"
	"github.com/ilhammhdd/kudaki-entities/events"
	"github.com/ilhammhdd/kudaki-store-service/externals/kafka"
	"gopkg.in/Shopify/sarama.v1"
)

const TOTAL_CONSUMER_MEMBER = 5

func AddStorefrontItem() {
	topic := []string{events.StoreTopic_name[int32(events.StoreTopic_ADD_STOREFRONT_ITEM_REQUESTED)]}
	groupID := uuid.New().String()

	for i := 0; i < TOTAL_CONSUMER_MEMBER; i++ {
		consMember := kafka.NewConsumptionMember(groupID, topic, sarama.OffsetNewest, "AddStorefrontItemRequested", i)
		signals := make(chan os.Signal)
		signal.Notify(signals)

		safekit.Do(func() {
			defer close(consMember.Close)
		ConsLoop:
			for {
				select {
				case msg := <-consMember.Messages:
					key, value, err := adapters.AddStorefrontItem(msg.Partition, msg.Offset, string(msg.Key), msg.Value, mysql.NewDBOperation())
					if err == nil {
						prod := kafka.NewProduction()
						prod.Set(events.StoreTopic_name[int32(events.StoreTopic_STOREFRONT_ITEM_ADDED)])
						start := time.Now()
						prodPart, prodOffset, err := prod.SyncProduce(key, value)
						duration := time.Since(start)
						errorkit.ErrorHandled(err)
						log.Printf("produced StorefrontItemAdded : partition = %d, offset = %d, duration = %f seconds, key = %s", prodPart, prodOffset, duration.Seconds(), key)
					}
				case errs := <-consMember.Errs:
					errorkit.ErrorHandled(errs)
				case <-signals:
					break ConsLoop
				}
			}
		})
	}
}

func DeleteStorefrontItem() {

	groupID := uuid.New().String()
	topics := []string{events.StoreTopic_name[int32(events.StoreTopic_DELETE_STOREFRONT_ITEM_REQUESTED)]}
	memberName := "DeleteStorefrontItemRequested"

	for i := 0; i < TOTAL_CONSUMER_MEMBER; i++ {
		consMember := kafka.NewConsumptionMember(groupID, topics, sarama.OffsetNewest, memberName, i)
		sig := make(chan os.Signal)
		signal.Notify(sig, os.Interrupt)

		safekit.Do(func() {
			defer close(consMember.Close)
			for {
				select {
				case msg := <-consMember.Messages:
					key, value, err := adapters.DeleteStorefrontItem(msg.Partition, msg.Offset, string(msg.Key), msg.Value, mysql.NewDBOperation())
					if err == nil {
						prod := kafka.NewProduction()
						prod.Set(events.StoreTopic_name[int32(events.StoreTopic_STOREFRONT_ITEM_DELETED)])
						start := time.Now()
						prodPart, prodOffset, err := prod.SyncProduce(key, value)
						duration := time.Since(start)
						errorkit.ErrorHandled(err)
						log.Printf("produced StorefrontItemDeleted : partition = %d, offset = %d, key = %s, duration = %f seconds", prodPart, prodOffset, key, duration.Seconds())
					}
				case errs := <-consMember.Errs:
					errorkit.ErrorHandled(errs)
				case <-sig:
					return
				}
			}
		})
	}
}

type StorefrontItemsRetrieval struct{}

func (s *StorefrontItemsRetrieval) Handle(interface{}) {}

func (s *StorefrontItemsRetrieval) Work() interface{} {

	groupID := uuid.New().String()
	topics := []string{events.StoreTopic_name[int32(events.StoreTopic_RETRIEVE_STOREFRONT_ITEMS_REQUESTED)]}

	for i := 0; i < TOTAL_CONSUMER_MEMBER; i++ {
		consMember := kafka.NewConsumptionMember(groupID, topics, sarama.OffsetNewest, "RetrieveStorefrontItemsRequested", i)
		sig := make(chan os.Signal)
		signal.Notify(sig, os.Interrupt)

		safekit.Do(func() {
			defer close(consMember.Close)
			for {
				select {
				case msg := <-consMember.Messages:
					sirr := adapters.StorefrontItemsRetrieval{
						Key:       string(msg.Key),
						Offset:    msg.Offset,
						Partition: msg.Partition,
						Value:     msg.Value,
					}

					key, value := sirr.Retrieve(mysql.NewDBOperation())
					s.produce(key, value)
				case errs := <-consMember.Errs:
					errorkit.ErrorHandled(errs)
				case <-sig:
					return
				}
			}
		})
	}

	return nil
}

func (s StorefrontItemsRetrieval) produce(key string, value []byte) {
	prod := kafka.NewProduction()
	prod.Set(events.StoreTopic_name[int32(events.StoreTopic_STOREFRONT_ITEMS_RETRIEVED)])

	start := time.Now()
	part, offset, err := prod.SyncProduce(key, value)
	duration := time.Since(start)

	errorkit.ErrorHandled(err)
	log.Printf("produced StorefrontItemsRetrieved : partition = %d, offset = %d, key = %s, duration = %f seconds", part, offset, key, duration.Seconds())
}

type KafkaRecord struct {
	Key     string
	Message []byte
}

type StorefrontItemUpdate struct{}

func (s StorefrontItemUpdate) Handle(interface{}) {}

func (s StorefrontItemUpdate) Work() interface{} {

	groupID := uuid.New().String()
	topics := []string{events.StoreTopic_name[int32(events.StoreTopic_UPDATE_STOREFRONT_ITEM_REQUESTED)]}

	for i := 0; i < TOTAL_CONSUMER_MEMBER; i++ {
		consMember := kafka.NewConsumptionMember(groupID, topics, sarama.OffsetNewest, "UpdateStorefrontItemRequested", i)

		safekit.Do(func() {
			defer close(consMember.Close)
			sig := make(chan os.Signal)
			signal.Notify(sig, os.Interrupt)

			for {
				select {
				case msg := <-consMember.Messages:
					siu := adapters.StorefrontItemUpdate{
						Key:       string(msg.Key),
						Message:   msg.Value,
						Offset:    msg.Offset,
						Partition: msg.Partition,
					}
					key, siuMsg, err := siu.Update(mysql.NewDBOperation())
					if err == nil {
						s.produce(key, siuMsg)
					}
				case errs := <-consMember.Errs:
					errorkit.ErrorHandled(errs)
				case <-sig:
					return
				}
			}
		})
	}

	return nil
}

func (s StorefrontItemUpdate) produce(key string, msg []byte) {
	prod := kafka.NewProduction()
	prod.Set(events.StoreTopic_name[int32(events.StoreTopic_STOREFRONT_ITEM_UPDATED)])

	start := time.Now()
	partition, offset, err := prod.SyncProduce(key, msg)
	duration := time.Since(start)

	errorkit.ErrorHandled(err)

	log.Printf("produced StorefrontItemUpdated : partition = %d, offset = %d, key = %s, duration = %f seconds", partition, offset, key, duration.Seconds())
}

type ItemsRetrieval struct{}

func (ir ItemsRetrieval) Handle(interface{}) {}

func (ir ItemsRetrieval) Work() interface{} {

	groupID := uuid.New().String()
	topics := []string{events.StoreTopic_name[int32(events.StoreTopic_RETRIEVE_ITEMS_REQUESTED)]}

	for i := 0; i < TOTAL_CONSUMER_MEMBER; i++ {
		consMember := kafka.NewConsumptionMember(groupID, topics, sarama.OffsetNewest, "RetrieveItemsRequested", i)

		safekit.Do(func() {
			defer close(consMember.Close)
			sig := make(chan os.Signal)
			signal.Notify(sig, os.Interrupt)

			for {
				select {
				case msg := <-consMember.Messages:
					irAdapter := adapters.ItemsRetrieval{
						Key:       string(msg.Key),
						Message:   msg.Value,
						Offset:    msg.Offset,
						Partition: msg.Partition,
					}
					key, irdMsg, err := irAdapter.Retrieve(mysql.NewDBOperation())
					if err == nil {
						ir.produce(key, irdMsg)
					}
				case errs := <-consMember.Errs:
					errorkit.ErrorHandled(errs)
				case <-sig:
					return
				}
			}
		})
	}

	return nil
}

func (ir ItemsRetrieval) produce(key string, msg []byte) {
	prod := kafka.NewProduction()
	prod.Set(events.StoreTopic_name[int32(events.StoreTopic_ITEMS_RETRIEVED)])

	start := time.Now()
	partition, offset, err := prod.SyncProduce(key, msg)
	duration := time.Since(start)

	errorkit.ErrorHandled(err)

	log.Printf("produced ItemsRetrieved : partition = %d, offset = %d, key = %s, duration = %f seconds", partition, offset, key, duration.Seconds())
}
