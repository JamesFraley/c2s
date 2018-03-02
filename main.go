package main

import (
	"log"
	"os"
	"time"

	"github.com/Shopify/sarama"
)

var topicName string

func main() {
	topicName = os.Getenv("topicName")
	zookeeperAddr := os.Getenv("zookeeperAddr")
	log.Print(zookeeperAddr)

	producer, err := createKafkaProducer(zookeeperAddr)
	if err != nil {
		log.Fatal("Failed to connect to Kafka")
	}

	//Ensures that the topic has been created in kafka
	log.Println("Checking Topic...")
	producer.Input() <- &sarama.ProducerMessage{
		Key:       sarama.StringEncoder("init"),
		Topic:     topicName,
		Timestamp: time.Now(),
	}

	for {
		log.Print("Start consume.")
		consumeMessages("127.0.0.1:2181", msgHandler())
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
