package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Shopify/sarama"
	_ "github.com/mattn/go-oci8"
)

var topicName string

func main() {
	topicName = os.Getenv("topicName")
	zookeeperAddr := os.Getenv("zookeeperAddr")
	oracleUser := os.Getenv("oracleUser")
	oraclePassword := os.Getenv("oraclePassword")
	oracleService := os.Getenv("oracleService")

	oracleConnectString := oracleUser + "/" + oraclePassword + "@" + oracleService
	db, err := sql.Open("oci8", oracleConnectString)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		fmt.Printf("Error connecting to the database: %s\n", err)
		return
	}

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
		consumeMessages(zookeeperAddr, msgHandler(), db)
	}
}

func msgHandler() func(m *sarama.ConsumerMessage, db *sql.DB) error {
	return func(m *sarama.ConsumerMessage, db *sql.DB) error {
		// Empty body means it is an init message
		if len(m.Value) == 0 {
			return nil
		}

		if err := db.Ping(); err != nil {
			fmt.Printf("Error connecting to the database: %s\n", err)
			return err
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
