package producer

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
	"os"
	"time"
)

type payload struct {
	TimeStamp int64 `json:"time_stamp"`
}


func Producer() {
	var data payload
	log.Println("Starting producer...")
	channel := connectToQueue()

	for {
		data.TimeStamp = time.Now().Unix()
		log.Println("Publishing to queue")
		go publishToExchange(data, channel)
		time.Sleep(time.Millisecond * 10)
	}
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

func publishToExchange(data payload, ch *amqp.Channel) (err error) {
	jsonBytes, _ := json.Marshal(data)
	exchange := os.Getenv("QUEUE_EXCHANGE")
	routingKey := os.Getenv("ROUTING_KEY")

	err = ch.Publish(exchange, routingKey, false, false, amqp.Publishing{
		ContentType:     "application/json",
		Body:            jsonBytes,
	})
	failOnError(err)
	log.Println("Published to queue")
	return
}