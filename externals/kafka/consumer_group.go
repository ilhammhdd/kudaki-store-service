package kafka

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/ilhammhdd/go-toolkit/errorkit"
	"github.com/ilhammhdd/go-toolkit/safekit"
	sarama "gopkg.in/Shopify/sarama.v1"
)

type ConsumptionMember struct {
	Ready    chan bool
	Messages chan *sarama.ConsumerMessage
	Errs     chan error
	Close    chan bool
	Name     string
	Number   int
}

func NewConsumptionMember(groupID string, topics []string, offset int64, name string, number int) *ConsumptionMember {

	version, err := sarama.ParseKafkaVersion(os.Getenv("KAFKA_VERSION"))
	errorkit.ErrorHandled(err)

	config := sarama.NewConfig()
	config.Version = version
	config.Consumer.Offsets.Initial = offset

	kafkaBrokers := strings.Split(os.Getenv("KAFKA_BROKERS"), ",")
	client, _ := sarama.NewConsumerGroup(kafkaBrokers, groupID, config)

	readyChan := make(chan bool)
	messagesChan := make(chan *sarama.ConsumerMessage)
	errsChan := make(chan error)
	closeChan := make(chan bool)

	member := &ConsumptionMember{
		Close:    closeChan,
		Errs:     errsChan,
		Messages: messagesChan,
		Ready:    readyChan,
		Name:     name,
		Number:   number,
	}

	safekit.Do(func() {
		member.Errs <- <-client.Errors()
	})

	safekit.Do(func() {
		<-member.Close
		errorkit.ErrorHandled(client.Close())
	})

	safekit.Do(func() {
		consErr := client.Consume(context.Background(), topics, member)
		errorkit.ErrorHandled(consErr)
	})

	return member
}

func (cm *ConsumptionMember) Setup(session sarama.ConsumerGroupSession) error {
	close(cm.Ready)
	log.Printf("%s consumer member no = %d ready...", cm.Name, cm.Number)
	return nil
}

func (cm *ConsumptionMember) Cleanup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (cm *ConsumptionMember) ConsumeClaim(session sarama.ConsumerGroupSession, consClaim sarama.ConsumerGroupClaim) error {

	for message := range consClaim.Messages() {
		cm.Messages <- message
		session.MarkMessage(message, "consumed")
	}

	return nil
}
