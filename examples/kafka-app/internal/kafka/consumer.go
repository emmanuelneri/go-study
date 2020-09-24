package kafka

import (
	"context"
	"github.com/Shopify/sarama"
	"kafka_app/config"
	"log"
)

const consumerGroupId = "kafka-app"

type Consumer struct {
	consumerGroup sarama.ConsumerGroup
	consumedChan  map[string]chan *sarama.ConsumerMessage
}

func NewConsumer() (*Consumer, error) {
	c := sarama.NewConfig()
	c.Version = config.KafkaVersion
	c.Consumer.Offsets.Initial = sarama.OffsetOldest
	c.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin

	consumer, err := sarama.NewConsumerGroup(config.KafkaBrokers(), consumerGroupId, c)
	if err != nil {
		return nil, err
	}

	return &Consumer{consumerGroup: consumer,
		consumedChan: make(map[string]chan *sarama.ConsumerMessage),
	}, nil
}

func (consumer *Consumer) FetchMessage(topic string) *sarama.ConsumerMessage {
	return <-consumer.consumedChan[topic]
}

func (consumer *Consumer) Subscribe(ctx context.Context, topics []string, ready chan bool) {
	log.Printf("topic subscribed %s", topics)
	go logErrors(consumer.consumerGroup.Errors())

	for _, topic := range topics {
		consumer.consumedChan[topic] = make(chan *sarama.ConsumerMessage)
	}

	handler := newConsumerHandler()
	go func() {
		for {
			err := consumer.consumerGroup.Consume(ctx, topics, handler)
			if err != nil {
				log.Fatalf("consume group error. %v", err)
			}

			handler.ready = make(chan bool)
		}
	}()

	<-handler.ready
	ready <- true
	for {
		message := <-handler.consumedChan
		topicChan := consumer.consumedChan[message.Topic]
		topicChan <- message
	}
}

func logErrors(errorsChan <-chan error) {
	for err := range errorsChan {
		log.Printf("Error: %v", err)
	}
}

type ConsumerHandler struct {
	ready        chan bool
	consumedChan chan *sarama.ConsumerMessage
}

func newConsumerHandler() *ConsumerHandler {
	return &ConsumerHandler{
		consumedChan: make(chan *sarama.ConsumerMessage),
		ready:        make(chan bool),
	}
}

func (ch *ConsumerHandler) Setup(sarama.ConsumerGroupSession) error {
	log.Println("ConsumerHandler setup")
	close(ch.ready)
	return nil
}

func (ch *ConsumerHandler) Cleanup(sarama.ConsumerGroupSession) error {
	log.Println("ConsumerHandler Cleanup")
	return nil
}

func (ch *ConsumerHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	log.Printf("ConsumeClaim: %s", claim.Topic())
	for message := range claim.Messages() {
		session.MarkMessage(message, "")
		ch.consumedChan <- message
	}

	return nil
}
