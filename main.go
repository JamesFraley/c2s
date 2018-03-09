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

	fileLocations, _ := loadIFL(db)
	log.Print(fileLocations)

	//http://go-database-sql.org/index.html
	//Strangely enough, sql.open doesn't create a connection
	//The db.ping will do it along with other sql & dml commands
	if err := db.Ping(); err != nil {
		db.Close()
		log.Fatalf("Error connecting to the database: %s\n", err)
	}

	log.Print("Start consume.")
	consumeMessages(kafkaConsumerAddr, msgHandler(), db, fileLocations)
}

func msgHandler() func(m *sarama.ConsumerMessage, db *sql.DB, fileLocations []iflLocation) error {
	return func(m *sarama.ConsumerMessage, db *sql.DB, fileLocations []iflLocation) error {
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

		log.Print("--Message Value----------------------------------------------------------------------------")
		log.Printf("%s", string(m.Value))
		log.Print("-----------------------------------------------")
		insertRow(tfrmEnvlope, db, fileLocations)
		return nil
	}
}

func loadIFL(db *sql.DB) ([]iflLocation, error) {
	log.Print("Starting loadIFL")
	var retVal []iflLocation

	selectSQL, err := db.Prepare("select ifl_id, absolute_path_unix from iots_file_locations")
	defer selectSQL.Close()
	if err != nil {
		log.Fatalf("Unable to prepare iots_file_locations query. %s\n", err)
	}

	rows, err := selectSQL.Query()
	if err != nil {
		log.Fatalf("Failed to query iots_file_locations. %s\n", err)
	}
	defer rows.Close()

	for rows.Next() {
		var iflID int
		var absolutePathUnix string
		err = rows.Scan(&iflID, &absolutePathUnix)
		if err != nil {
			log.Fatalf("Unable to scan iots_file_locations. %s\n", err)
		}
		dl := iflLocation{iflID: iflID, absolutePathUnix: absolutePathUnix}
		retVal = append(retVal, dl)
	}
	log.Print("Exiting loadIFL")
	return retVal, nil
}
