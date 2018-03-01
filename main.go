package main

import (
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
		log.Printf("BlockTimestamp=%s\n", m.BlockTimestamp)
		log.Printf("Headers=%s\n", m.Headers)
		log.Printf("Key=%s\n", m.Key)
		log.Printf("Offset=%d\n", m.Offset)
		log.Printf("Partition=%v\n", m.Partition)
		log.Printf("Timestamp=%s\n", m.Timestamp)
		log.Printf("Topic=%s\n", m.Topic)
		log.Printf("Value.typeOf()=%T\n", m.Value)
		log.Printf("Value=%s\n", m.Value)
		log.Printf("\n\n")
		return nil
	}
}
