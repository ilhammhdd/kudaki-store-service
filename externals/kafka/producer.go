package kafka

import (
	"os"
	"strings"
	"time"

	"github.com/ilhammhdd/go-toolkit/errorkit"
	sarama "gopkg.in/Shopify/sarama.v1"
)

type Production struct {
	topic string
}

func NewProduction() *Production { return &Production{} }

func (p *Production) Set(topic string) {
	p.topic = topic
}

func (p *Production) Get() (topic string) {
	return p.topic
}

func (p *Production) SyncProduce(key string, value []byte) (partition int32, offset int64, err error) {
	type rtrn struct {
		Successes bool
		Errors    bool
	}

	config := sarama.NewConfig()
	config.Producer.Return = rtrn{
		Successes: true,
		Errors:    true}
	config.Producer.Compression = sarama.CompressionNone
	config.Producer.Flush.Frequency = time.Duration(0)
	config.Producer.RequiredAcks = sarama.WaitForLocal

	prod, err := sarama.NewSyncProducer(strings.Split(os.Getenv("KAFKA_BROKERS"), ","), config)
	errorkit.ErrorHandled(err)

	defer prod.Close()

	topic := p.Get()

	msg := sarama.ProducerMessage{
		Topic:     topic,
		Offset:    sarama.OffsetNewest,
		Key:       sarama.StringEncoder(key),
		Value:     sarama.ByteEncoder(value),
		Timestamp: time.Now(),
	}

	return prod.SendMessage(&msg)
}
