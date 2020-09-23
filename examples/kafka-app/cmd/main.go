package main

import (
	"context"
	"github.com/Shopify/sarama"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"kafka_app/internal/kafka"
	"log"
	"net/http"
	"strconv"
	"time"
)

const (
	nameTopic  = "names"
	colorTopic = "colors"
)

func main() {
	log.Println("### starting kafka demo ###")

	go consume([]string{nameTopic, colorTopic})

	ticker := time.NewTicker(3 * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				go produce(nameTopic, 20)
				go produce(colorTopic, 10)
			}
		}
	}()

	http.Handle("/metrics", promhttp.Handler())
	log.Panicln(http.ListenAndServe(":8082", nil))
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

	consumerReady := make(chan bool)
	go consumer.Subscribe(context.Background(), topics, consumerReady)
	<-consumerReady

	for _, topic := range topics {
		t := topic
		go func() {
			for {
				msg := consumer.FetchMessage(t)
				log.Printf("topic: %s - Message: %s", msg.Topic, string(msg.Value))
			}
		}()
	}
}
