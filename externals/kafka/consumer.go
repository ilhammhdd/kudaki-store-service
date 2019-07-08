package kafka

import (
	"os"
	"os/signal"
	"strings"

	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/go-toolkit/safekit"
	sarama "gopkg.in/Shopify/sarama.v1"
)

type Consumption struct {
	Topic     string
	Partition int32
	Offset    int64
}

func NewConsumption() *Consumption { return &Consumption{} }

func (c *Consumption) Set(topic string, partition int32, offset int64) {
	c.Topic = topic
	c.Partition = partition
	c.Offset = offset
}

func (c *Consumption) Get() (string, int32, int64) {
	return c.Topic, c.Partition, c.Offset
}

// this function still tightly coupled to sarama by PartitionConsumer return value
func (c *Consumption) Consume() (sarama.PartitionConsumer, chan os.Signal, chan bool) {
	closeConsumer := make(chan bool)

	cons, err := sarama.NewConsumer(strings.Split(os.Getenv("KAFKA_BROKERS"), ","), nil)
	errorkit.ErrorHandled(err)

	topic, partition, offset := c.Get()

	partCons, err := cons.ConsumePartition(topic, partition, offset)
	errorkit.ErrorHandled(err)

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	safekit.Do(func() {
		_, ok := <-closeConsumer
		if !ok {
			cons.Close()
			partCons.Close()
		}
	})

	return partCons, signals, closeConsumer
}
