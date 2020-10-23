package consumer

import (
	"github.com/streadway/amqp"
	"log"
	"os"
	"time"
)

type payload struct {
	TimeStamp int64 `json:"time_stamp"`
}

func Consumer() {
	log.Println("Consuming")
	channel := connectToQueue()
	consumeFromQueue(channel)
	defer channel.Close()
}


func failOnError(err error){
	if err != nil {
		log.Fatal(err)
	}
}

func connectToQueue() *amqp.Channel {
	rabbitmqURL := os.Getenv("RABBITMQ_URL")
	log.Println("Connecting to rabbit mq")
	conn, err := amqp.Dial(rabbitmqURL)
	failOnError(err)

	log.Println("Creating a channel")
	ch, err := conn.Channel()
	failOnError(err)

	return ch
}

func consumeFromQueue(ch *amqp.Channel) {
	queue := os.Getenv("QUEUE_NAME")
	msgs, err := ch.Consume(queue, "consumer-1", false, false, false, false, nil)
	failOnError(err)

	consumerChan := make(chan bool)
	go func() {
		for msg := range msgs {
			msgBytes := msg.Body
			log.Println("Received ", time.Now().Nanosecond(), string(msgBytes))
			_ = msg.Ack(false)
		}
	}()
	<- consumerChan

}
