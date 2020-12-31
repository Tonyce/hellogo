package main

import (
	"fmt"

	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

func main() {

	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers":  "localhost",
		"session.timeout.ms": 6000,
		"auto.offset.reset":  "latest"})
	if err != nil {
		panic(err)
	}

	defer p.Close()

	// Delivery report handler for produced messages
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	// Produce messages to topic (asynchronously)
	topic := "TutorialTopic"
	// s1 := rand.NewSource(time.Now().UnixNano())
	// r1 := rand.New(s1)
	for _, word := range []string{"Welcome", "to", "the", "Confluent", "Kafka", "Golang", "client", "Welcome", "to", "the", "Confluent", "Kafka", "Golang", "client"} {
		// partition := r1.Int31n(5)
		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic:     &topic,
				Partition: -1,
			},
			Value: []byte(word),
		}, nil)
	}

	// Wait for message deliveries before shutting down
	p.Flush(15 * 1000)
}
