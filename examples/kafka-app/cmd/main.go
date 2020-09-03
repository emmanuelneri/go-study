package main

import (
	"context"
	"github.com/Shopify/sarama"
	"kafka_app/internal/kafka"
	"log"
	"strconv"
)

const (
	nameTopic  = "names"
	colorTopic = "colors"
)

func main() {
	log.Println("### starting kafka demo ###")

	go consume([]string{nameTopic, colorTopic})
	go produce(nameTopic, 5)
	go produce(colorTopic, 5)

	select {}
}

func produce(topic string, quantity int) {
	producer, err := kafka.NewProducer()
	if err != nil {
		panic(err)
	}

	for i := 0; i < quantity; i++ {
		message := &sarama.ProducerMessage{
			Topic: topic,
			Key:   sarama.ByteEncoder(strconv.Itoa(i)),
			Value: sarama.ByteEncoder(topic + " " + strconv.Itoa(i)),
		}

		producer.Produce(message)
	}
}

func consume(topics []string) {
	consumer, err := kafka.NewConsumer()
	if err != nil {
		panic(err)
	}

	for _, topic := range topics {
		t := topic
		go func() {
			for {
				msg := consumer.FetchMessage(t)
				log.Printf("topic: %s - Message: %s", msg.Topic, string(msg.Value))
			}
		}()
	}

	go consumer.Subscribe(context.Background(), topics)
}
