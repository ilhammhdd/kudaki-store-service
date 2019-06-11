package usecases

import (
	"errors"
	"os"

	"gopkg.in/Shopify/sarama.v1"
)

var ErrProtoUnmarshalType = errors.New("proto: unmarshal: wire type and go type doesn't match")

type EventDrivenConsumer interface {
	Set(topic string, partition int32, offset int64)
	Get() (topic string, partition int32, offset int64)
	Consume() (partCons sarama.PartitionConsumer, signals chan os.Signal, close chan bool)
}

type EventDrivenProducer interface {
	Set(topic string)
	Get() (topic string)
	SyncProduce(key string, value []byte) (producedPartition int32, producedOffset int64, err error)
}

type EventDrivenConsumerGroup interface {
	Set(groupID string, topics []string, offset int64)
	Messages() chan *sarama.ConsumerMessage
	Errors() chan error
	Close()
}
