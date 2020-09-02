package kafka

import (
	"github.com/Shopify/sarama"
	"kafka_app/config"
	"log"
)

type Producer struct {
	async sarama.AsyncProducer
}

func NewProducer() (*Producer, error) {
	c := sarama.NewConfig()
	c.Version = config.KafkaVersion
	c.Producer.Return.Successes = true
	c.Producer.Return.Errors = true
	asyncProducer, err := sarama.NewAsyncProducer(config.KafkaBrokers(), c)
	if err != nil {
		return nil, err
	}

	return &Producer{async: asyncProducer}, nil
}

func (producer *Producer) Produce(message *sarama.ProducerMessage) {
	producer.async.Input() <- message
	go producer.logErrors()
	go producer.logSuccess()
}

func (producer *Producer) logSuccess() {
	for {
		produced := <-producer.async.Successes()
		log.Printf("Produced offset: %d", produced.Offset)
	}
}

func (producer *Producer) logErrors() {
	for {
		pe := <-producer.async.Errors()
		log.Printf("[ERROR] - producer error. %v", pe.Err)
	}
}
