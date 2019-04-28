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

func AddStorefrontItem() {
	topic := []string{events.StoreTopic_name[int32(events.StoreTopic_ADD_STOREFRONT_ITEM_REQUESTED)]}
	groupID := uuid.New().String()

	for i := 0; i < 10; i++ {
		consMember := kafka.NewConsumptionMember(groupID, topic, sarama.OffsetNewest, "AddStorefrontItemRequested", i)
		signals := make(chan os.Signal)
		signal.Notify(signals)

		safekit.Do(func() {
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
