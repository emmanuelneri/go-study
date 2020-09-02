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
	consumedChan  chan *sarama.ConsumerMessage
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

	return &Consumer{consumerGroup: consumer, consumedChan: make(chan *sarama.ConsumerMessage)}, nil
}

func (consumer *Consumer) ConsumerMessage() chan *sarama.ConsumerMessage {
	return consumer.consumedChan
}

func (consumer *Consumer) Subscribe(ctx context.Context, topic string) {
	go logErrors(consumer.consumerGroup.Errors())

	handler := ConsumerHandler{consumedChan: make(chan *sarama.ConsumerMessage)}
	go consumer.consume(ctx, topic, &handler)

	for {
		message := <-handler.consumedChan
		consumer.consumedChan <- message
	}
}

func (consumer *Consumer) consume(ctx context.Context, topic string, handler *ConsumerHandler) {
	err := consumer.consumerGroup.Consume(ctx, []string{topic}, handler)
	if err != nil {
		log.Fatalf("[consume groupd error. %v", err)
	}
}

func logErrors(errorsChan <-chan error) {
	log.Printf("Error: %v", <-errorsChan)
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
	log.Println("ConsumeClaim")
	for message := range claim.Messages() {
		session.MarkMessage(message, "")
		ch.consumedChan <- message
	}

	return nil
}
