package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Shopify/sarama"
	_ "github.com/mattn/go-oci8"
)

var topicName string

func main() {
	topicName = os.Getenv("topicName")
	//zookeeperAddr := os.Getenv("zookeeperAddr")
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

	log.Print("Start consume.")
	consumeMessages(kafkaConsumerAddr, msgHandler(), db)
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

		tfrmEnvlope := TFRMCatalogEnvlope{}
		reader := strings.NewReader(string(m.Value))
		decoder := json.NewDecoder(reader)
		err := decoder.Decode(&tfrmEnvlope)
		if err != nil {
			log.Printf("err=%s\n", err)
		}

		log.Printf("%s", string(m.Value))
		log.Print("------------------------------------------------------------------------------")
		meta := tfrmEnvlope.Catalog.Meta
		log.Printf("SOURCE_FILENAME=%s\n", (*meta).Source.FileName)
		log.Printf("FILESIZE=%d\n", (*meta).Source.FileSize)
		log.Printf("MD5=%s\n", (*meta).Source.Md5)

		class := (*meta).Classification
		marking := (*class).Marking
		log.Printf("CLASSIFICATION=%s\n", marking)

		loc := tfrmEnvlope.Catalog.Locations
		log.Print(loc[0].Uri)
		log.Print("------------------------------------------------------------------------------")
		insertRow(tfrmEnvlope, db)

		//		json.Unmarshal([]byte(m.Value), &record)
		return nil
	}
}
