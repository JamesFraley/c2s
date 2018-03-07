package main

import (
	"database/sql"
	"log"
)

func insertRow(p TFRMCatalogEnvlope, db *sql.DB) error {
	// meta := p.Catalog.Meta
	// source_filename := (*meta).Source.FileName
	// file_size := (*meta).Source.FileSize
	// md5 := (*meta).Source.Md5
	// class := (*meta).Classification
	// marking := (*class).Marking
	// loc := p.Catalog.Locations
	// log.Print(loc[0].Uri)

	sqlStrg := `
	declare
   a_row iots_file_master%ROWTYPE;
   p_row iots_file_master%ROWTYPE;
begin
   p_row.source_filename := 'jims_file2';
   p_row.classification := '1000';
   p_row.state := 'PROCESSED';
   p_row.ifl_id := 1;
   p_row.file_origin := 'APX';
   p_row.checksum := '10923847';
   p_row.file_size := 120;
   p_row.uri_location := '/iots/prod/somewhere';
   a_row := file_master_interface.register_file(p_row);
end;`

	ret, err := db.Exec(sqlStrg)
	log.Print(ret)
	if err != nil {
		log.Fatal(err)
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
	return nil
}
