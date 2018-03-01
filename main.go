package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Shopify/sarama"
)

const (
	hostPort  = "3000"
	topicName = "kafka-example-topic"
)

func main() {
	producer, err := createKafkaProducer("127.0.0.1:9092")

	if err != nil {
		log.Fatal("Failed to connect to Kafka")
	}

	//Ensures that the topic has been created in kafka
	producer.Input() <- &sarama.ProducerMessage{
		Key:       sarama.StringEncoder("init"),
		Topic:     topicName,
		Timestamp: time.Now(),
	}

	log.Println("Creating Topic...")

	for {
		log.Print("----------------")
		consumeMessages("127.0.0.1:2181", msgHandler())
		log.Print("~~~~~~~~~~~~~~~~")
	}

}

func msgHandler() func(m *sarama.ConsumerMessage) error {
	return func(m *sarama.ConsumerMessage) error {
		// Empty body means it is an init message
		if len(m.Value) == 0 {
			return nil
		}
		log.Print(m)
		fmt.Printf("%s\n", m.BlockTimestamp)
		fmt.Printf("%s\n", m.Headers)
		fmt.Printf("Key=%s\n", m.Key)
		fmt.Printf("Offset=%d\n", m.Offset)
		fmt.Printf("Partition=%v\n", m.Partition)
		fmt.Printf("%s\n", m.Timestamp)
		fmt.Printf("%s\n", m.Topic)
		fmt.Printf("%T\n", m.Value)
		fmt.Printf("%s\n", m.Value)
		fmt.Printf("\n\n")
		return nil
	}
}
