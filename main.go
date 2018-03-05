package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	catalog "bitbucket.di2e.net/scm/pir/go-catalog-types.git"
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
	kafkaConsumerAddr := os.Getenv("kafkaConsumerAddr")

	var db *sql.DB

	oracleConnectString := oracleUser + "/" + oraclePassword + "@" + oracleService
	db, err := sql.Open("oci8", oracleConnectString)
	if err != nil {
		log.Fatalf("Error opening oracle connection: %s\n", err)
	} else {
		defer db.Close()
	}

	//http://go-database-sql.org/index.html
	//Strangely enough, sql.open doesn't create a connection
	//The db.ping will do it along with other sql & dml commands
	if err := db.Ping(); err != nil {
		db.Close()
		log.Fatalf("Error connecting to the database: %s\n", err)
	}

	producer, err := createKafkaProducer(zookeeperAddr)
	if err != nil {
		log.Fatalf("Failed to connect to Kafka")
	}

	//Ensures that the topic has been created in kafka
	log.Println("Checking Topic...")
	producer.Input() <- &sarama.ProducerMessage{
		Key:       sarama.StringEncoder("init"),
		Topic:     topicName,
		Timestamp: time.Now(),
	}

	//	for {
	log.Print("Start consume.")
	consumeMessages(kafkaConsumerAddr, msgHandler(), db)
	//	}
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

		var record catalog.CatalogRecord
		json.Unmarshal([]byte(m.Value), &record)

		log.Printf("%s", string(m.Value))
		//insertRow(p, db)

		return nil
	}
}

// p.Name = "James Fraley"
// p.Addr.StreetAddr = "1823 Andrea Circle"
// p.Addr.City = "Beavercreek"
// p.Addr.State = "OH"
// p.Addr.Zipcode = 45432
// p.Point.Latitude = 45.1
// p.Point.Longitude = 90.2
// p.FavColors = []string{"a", "b", "c", "d"}
// b, _ := json.Marshal(p)
// log.Print(string(b[:]))

//b, _ := json.Marshal(m.Value)

// log.Panicf("P=%v", p)
// log.Printf("BlockTimestamp=%s\n", m.BlockTimestamp)
// log.Printf("Headers=%s\n", m.Headers)
// log.Printf("Key=%s\n", m.Key)
// log.Printf("Offset=%d\n", m.Offset)
// log.Printf("Partition=%v\n", m.Partition)
// log.Printf("Timestamp=%s\n", m.Timestamp)
// log.Printf("Topic=%s\n", m.Topic)
// log.Printf("Value.typeOf()=%T\n", m.Value)
// log.Printf("Value=%s\n", m.Value)
// log.Printf("\n\n")
