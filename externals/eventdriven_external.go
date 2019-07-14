package externals

import (
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/golang/protobuf/proto"

	"github.com/google/uuid"
	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/go-toolkit/safekit"
	"github.com/ilhammhdd/kudaki-storefront-service/adapters"
	"github.com/ilhammhdd/kudaki-storefront-service/externals/kafka"
	"github.com/ilhammhdd/kudaki-storefront-service/usecases"
	"gopkg.in/Shopify/sarama.v1"
)

const TOTAL_CONSUMER_MEMBER = 3

type PostUsecaseExecutor interface {
	ExecutePostUsecase(inEvent proto.Message, outEvent proto.Message)
}

type PostDownstreamUsecaseExecutor interface {
	ExecutePostDownstreamUsecase(inEvent proto.Message, usecaseRes *usecases.UsecaseHandlerResponse)
}

type EventDrivenExternal struct {
	inTopics            []string
	eventName           string
	eventDrivenAdapter  adapters.EventDrivenAdapter
	eventDrivenUsecase  usecases.EventDrivenUsecase
	outTopic            string
	PostUsecaseExecutor PostUsecaseExecutor
}

func (edc *EventDrivenExternal) produce(key string, msg []byte) {
	prod := kafka.NewProduction()
	prod.Set(edc.outTopic)
	start := time.Now()
	partition, offset, err := prod.SyncProduce(key, msg)
	duration := time.Since(start)
	errorkit.ErrorHandled(err)

	log.Printf("produced %s : partition = %d, offset = %d, key = %s, duration = %f seconds", edc.outTopic, partition, offset, key, duration.Seconds())
}

func (edc *EventDrivenExternal) handleSingleConsumer() {
	cons := kafka.NewConsumption()
	cons.Set(edc.inTopics[0], 0, sarama.OffsetNewest)
	partCons, sig, closeChan := cons.Consume()
	cl := adapters.ConsumerLog{EventName: edc.eventName}
	defer close(closeChan)

	for {
		select {
		case msg := <-partCons.Messages():
			if inEvent, ok := edc.eventDrivenAdapter.ParseIn(msg.Value); ok {
				cl.Log(msg.Partition, msg.Offset, string(msg.Key))
				outEvent := edc.eventDrivenUsecase.Handle(inEvent)

				if edc.PostUsecaseExecutor != nil {
					edc.PostUsecaseExecutor.ExecutePostUsecase(inEvent, outEvent)
				}

				outKey, outMsg := edc.eventDrivenAdapter.ParseOut(outEvent)
				edc.produce(outKey, outMsg)
			}
		case errs := <-partCons.Errors():
			log.Printf("error while consuming %s : %s", edc.inTopics[0], errs.Err.Error())
		case <-sig:
			return
		}
	}
}

func (edc *EventDrivenExternal) handle() {
	groupID := uuid.New().String()
	cl := adapters.ConsumerLog{EventName: edc.eventName}

	for i := 0; i < TOTAL_CONSUMER_MEMBER; i++ {
		consMember := kafka.NewConsumptionMember(groupID, edc.inTopics, sarama.OffsetNewest, edc.eventName, i)
		signals := make(chan os.Signal)
		signal.Notify(signals)

		safekit.Do(func() {
			defer close(consMember.Close)
		ConsLoop:
			for {
				select {
				case msg := <-consMember.Messages:
					if inEvent, ok := edc.eventDrivenAdapter.ParseIn(msg.Value); ok {
						cl.Log(msg.Partition, msg.Offset, string(msg.Key))
						outEvent := edc.eventDrivenUsecase.Handle(inEvent)

						if edc.PostUsecaseExecutor != nil {
							edc.PostUsecaseExecutor.ExecutePostUsecase(inEvent, outEvent)
						}

						outKey, outMsg := edc.eventDrivenAdapter.ParseOut(outEvent)
						edc.produce(outKey, outMsg)
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

type EventDrivenDownstreamExternal struct {
	inTopics            []string
	eventName           string
	eventDrivenAdapter  adapters.EventDrivenDownstreamAdapter
	eventDrivenUsecase  usecases.EventDrivenDownstreamUsecase
	PostUsecaseExecutor PostDownstreamUsecaseExecutor
}

func (edde *EventDrivenDownstreamExternal) handleSingleConsumer() {
	cons := kafka.NewConsumption()
	cons.Set(edde.inTopics[0], 0, sarama.OffsetNewest)
	partCons, sig, closeChan := cons.Consume()
	cl := adapters.ConsumerLog{EventName: edde.eventName}
	defer close(closeChan)

	for {
		select {
		case msg := <-partCons.Messages():
			if inEvent, ok := edde.eventDrivenAdapter.ParseIn(msg.Value); ok {
				cl.Log(msg.Partition, msg.Offset, string(msg.Key))
				stat := edde.eventDrivenUsecase.Handle(inEvent)

				if edde.PostUsecaseExecutor != nil {
					edde.PostUsecaseExecutor.ExecutePostDownstreamUsecase(inEvent, stat)
				}
			}
		case errs := <-partCons.Errors():
			log.Printf("error while consuming %s : %s", edde.inTopics[0], errs.Err.Error())
		case <-sig:
			return
		}
	}
}
func (edde *EventDrivenDownstreamExternal) handle() {
	groupID := uuid.New().String()
	cl := adapters.ConsumerLog{EventName: edde.eventName}

	for i := 0; i < TOTAL_CONSUMER_MEMBER; i++ {
		consMember := kafka.NewConsumptionMember(groupID, edde.inTopics, sarama.OffsetNewest, edde.eventName, i)
		signals := make(chan os.Signal)
		signal.Notify(signals)

		safekit.Do(func() {
			defer close(consMember.Close)
		ConsLoop:
			for {
				select {
				case msg := <-consMember.Messages:
					if inEvent, ok := edde.eventDrivenAdapter.ParseIn(msg.Value); ok {
						cl.Log(msg.Partition, msg.Offset, string(msg.Key))
						stat := edde.eventDrivenUsecase.Handle(inEvent)

						if edde.PostUsecaseExecutor != nil {
							edde.PostUsecaseExecutor.ExecutePostDownstreamUsecase(inEvent, stat)
						}
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
