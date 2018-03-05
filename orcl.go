package main

import (
	"database/sql"
	"log"
)

func insertRow(p place, db *sql.DB) error {
	stmt, err := db.Prepare("INSERT INTO place(name, streetaddress, city, state, zipcode, latitude, longitude) VALUES(:1, :2, :3, :4, :5, :6, :7)")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(
		p.Name,
		p.Addr.StreetAddr,
		p.Addr.City,
		p.Addr.State,
		p.Addr.Zipcode,
		p.Point.Latitude,
		p.Point.Longitude)
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
