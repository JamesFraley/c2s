package main

import (
	"database/sql"
	"log"
)

func insertRow(p place, db *sql.DB) error {
	stmt, err := db.Prepare("INSERT INTO place(name, address, city, population, latitiude, longitude) VALUES(?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(p.Name, p.Addr.StreetAddr, p.Addr.City)
	if err != nil {
		log.Fatal(err)
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ID = %d, affected = %d\n", lastId, rowCnt)
	return nil
}
