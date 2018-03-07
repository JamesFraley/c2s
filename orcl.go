package main

import (
	"database/sql"
	"fmt"
	"log"
)

func insertRow(p TFRMCatalogEnvlope, db *sql.DB) error {
	sqlStrg := `
declare
   a_row iots_file_master%ROWTYPE;
   p_row iots_file_master%ROWTYPE;
begin
   p_row.source_filename := :1;
   p_row.classification := :2;
   p_row.state := :3;
   p_row.ifl_id := :4;
   p_row.file_origin := :5;
   p_row.checksum := :6;
   p_row.file_size := :7;
   p_row.uri_location := :8;
   a_row := file_master_interface.register_file(p_row);
end;`

	meta := p.Catalog.Meta
	source_filename := (*meta).Source.FileName
	file_size := (*meta).Source.FileSize
	md5 := (*meta).Source.Md5
	//class := (*meta).Classification
	//marking := (*class).Marking
	loc := p.Catalog.Locations
	log.Print(loc[0].Uri)

	p_source_filename := source_filename
	p_class := "1000"
	p_state := "PROCESSED"
	p_ifl_id := 1
	p_file_origin := "APX"
	p_checksum := md5
	p_file_size := file_size
	p_uri_location := loc[0].Uri

	pSQL, err := db.Prepare(sqlStrg)
	defer pSQL.Close()
	rows, err := pSQL.Query(p_source_filename, p_class, p_state, p_ifl_id, p_file_origin, p_checksum, p_file_size, p_uri_location)
	defer rows.Close()

	for rows.Next() {
		//var i int
		fmt.Print("1")
		//err = rows.Scan(&i)
		// if err != nil {
		// 	log.Fatal(err)
	}

	fmt.Printf("err=%v\n", err)

	// ret, err := db.Exec(sqlStrg, p_source_filename, p_class, p_state, p_ifl_id, p_file_origin, p_checksum, p_file_size, p_uri_location)
	// log.Print(ret)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	return nil
}

// var rowCount int
// err = sql.QueryRow(p.Name).Scan(&rowCount)
// if rowCount > 0 && err != nil {
// 	log.Printf("rowCount=%d\n", rowCount)
// 	log.Fatalf("Error checking for existing file.  %s\n", err)
// }

// if rowCount > 0 {
// 	log.Printf("Attempting to ingest a duplicate file %s.  The database already reports %d rows.", p.Name, rowCount)
// 	return nil
// }

// stmt, err := db.Prepare("INSERT INTO place(name, streetaddress, city, state, zipcode, latitude, longitude) VALUES(:1, :2, :3, :4, :5, :6, :7)")
// if err != nil {
// 	log.Fatal(err)
// }
// res, err := stmt.Exec(
// 	p.Name,
// 	p.Addr.StreetAddr,
// 	p.Addr.City,
// 	p.Addr.State,
// 	p.Addr.Zipcode,
// 	p.Point.Latitude,
// 	p.Point.Longitude)
// if err != nil {
// 	log.Fatal(err)
// }
// lastId, err := res.LastInsertId()
// if err != nil {
// 	log.Fatal(err)
// }
// rowCnt, err := res.RowsAffected()
// if err != nil {
// 	log.Fatal(err)
// }
// log.Printf("ID = %d, affected = %d\n", lastId, rowCnt)
