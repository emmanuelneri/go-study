package kafka

import (
	"context"
	"github.com/Shopify/sarama"
	"kafka_app/config"
	"log"
	"time"
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
	for {
		select {
		case message := <-consumer.consumedChan[topic]:
			return message
		case <-time.After(time.Millisecond):
		}
	}
}

func (consumer *Consumer) Subscribe(ctx context.Context, topics []string) {
	log.Printf("topic subscribed %s", topics)
	go logErrors(consumer.consumerGroup.Errors())

	for _, topic := range topics {
		consumer.consumedChan[topic] = make(chan *sarama.ConsumerMessage)
	}

	handler := ConsumerHandler{consumedChan: make(chan *sarama.ConsumerMessage)}
	go consumer.consume(ctx, topics, &handler)

	for {
		message := <-handler.consumedChan
		topicChan := consumer.consumedChan[message.Topic]
		topicChan <- message
	}
}

func (consumer *Consumer) consume(ctx context.Context, topics []string, handler *ConsumerHandler) {
	for {
		err := consumer.consumerGroup.Consume(ctx, topics, handler)
		if err != nil {
			log.Fatalf("consume group error. %v", err)
		}
	}
}

func logErrors(errorsChan <-chan error) {
	for err := range errorsChan {
		log.Printf("Error: %v", err)
	}
}

type ConsumerHandler struct {
	consumedChan chan *sarama.ConsumerMessage
}

func (ch *ConsumerHandler) Setup(sarama.ConsumerGroupSession) error {
	log.Println("ConsumerHandler setup")
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
