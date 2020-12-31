package main

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
		"group.id":          "my-01-group",
		"auto.offset.reset": "earliest",
		// "session.timeout.ms": 6000,
	})

	if err != nil {
		panic(err)
	}

	c.SubscribeTopics([]string{"TutorialTopic"}, nil)

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			fmt.Println(msg.Value, string(msg.Value))
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		} else {
			// The client will automatically try to recover from all errors.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}

	// c.Close()
}
