package config

import (
	"github.com/Shopify/sarama"
	"strings"
)

var (
	kafkaBrokers = "localhost:9092"
	KafkaVersion = sarama.V2_5_0_0
)

func KafkaBrokers() []string {
	return strings.Split(kafkaBrokers, ",")
}
