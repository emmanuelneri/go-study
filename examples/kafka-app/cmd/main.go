package main

import (
	"context"
	"github.com/Shopify/sarama"
	"kafka_app/internal/kafka"
	"log"
	"strconv"
)

const (
	topic = "names"
)

func main() {
	log.Println("### starting kafka demo ###")

	go consume()
	go produce(10)

	select {}
}

func produce(quantity int) {
	producer, err := kafka.NewProducer()
	if err != nil {
		panic(err)
	}

	for i := 0; i < quantity; i++ {
		message := &sarama.ProducerMessage{
			Topic: topic,
			Key:   sarama.ByteEncoder(strconv.Itoa(i)),
			Value: sarama.ByteEncoder("Name " + strconv.Itoa(i)),
		}

		producer.Produce(message)
	}
}

func consume() {
	consumer, err := kafka.NewConsumer()
	if err != nil {
		panic(err)
	}

	go func() {
		for {
			msg := <-consumer.ConsumerMessage()
			log.Printf("Message: topic = %s - value = %s", msg.Topic, string(msg.Value))
		}
	}()

	consumer.Subscribe(context.Background(), topic)
}
