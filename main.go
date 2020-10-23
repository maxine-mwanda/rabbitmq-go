package main

import (
	"github.com/joho/godotenv"
	"rabbitMq/consumer"
	"rabbitMq/producer"
)

func main() {
	_ = godotenv.Load()
	qChan := make(chan bool)
	go consumer.Consumer()
	go producer.Producer()
	<- qChan
}
