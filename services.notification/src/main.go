package main

import (
	"wptest/services.notification/src/kafka"
	"wptest/shared/config"
)

func main() {
	config := config.LoadConfig(".")
	kafka.RegisterConsumer(config.KafkaOrderTopic)
}
