package main

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func insertRow(p TFRMCatalogEnvlope, db *sql.DB) error {
	insertSQL := `
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

	p_source_filename := source_filename
	p_class := "1000"
	p_state := "PROCESSED"
	p_ifl_id := 1
	p_file_origin := "APX"
	p_checksum := md5
	p_file_size := file_size

	var p_uri_location string
	archiveFound := false
	for _, val := range loc {
		if val.Name == "archive" {
			p_uri_location = val.Uri
			archiveFound = true
		}
	}

	if archiveFound == false {
		log.Printf("No archive location found for: %s\n", p_source_filename)
		return nil
	}

	colonPos := strings.Index(p_uri_location, ":")
	fileFQN := p_uri_location[colonPos+1:] //this removes the file:/
	filePath := filepath.Dir(fileFQN)
	log.Printf("filePath=%s\n", filePath)

	base := "/c2s/prod"
	lenBase := len(base)
	var fileLocation string
	if filePath[0:lenBase] == base {
		fileLocation = filePath[lenBase:]
	} else {
		log.Printf("Bad file location.  base=%s\t\t filePath=%s\n", base, filePath)
		log.Print(filePath[0:lenBase])
		return nil
	}
	log.Print(fileLocation)

	//
	//   This section ensures the file hasn't already been processed.  That is an error.
	//
	selectSQL, err := db.Prepare("select nvl(sum(1),0) from iots_file_master where source_filename=:1")
	if err != nil {
		log.Fatalf("Failed to Prepare: \"select sum(1) ROWCOUNT from iots_file_master where source_filename=:1\": %s\n", err)
	}
	defer selectSQL.Close()
	row := selectSQL.QueryRow(p_source_filename)
	var rowCount int
	err = row.Scan(&rowCount)
	if err != nil {
		log.Fatalf("Failed to retrieve rowcount for %s. err=%s\n", p_source_filename, err)
	}
	if rowCount > 0 {
		log.Printf("%s has been previously processed", p_source_filename)
		return nil
	}

	//
	// This section inserts the record into the database
	//
	_, err = db.Exec(insertSQL, p_source_filename, p_class, p_state, p_ifl_id, p_file_origin, p_checksum, p_file_size, fileLocation)
	if err != nil {
		log.Fatalf("Execution of register_file failed: %s\n", err)
	}

	//
	//  This section retrieves the file name with the version number prepended.
	//    The value is necessary to create the OS link
	//
	selectSQL, err = db.Prepare("select filename from iots_file_master where source_filename=:1")
	defer selectSQL.Close()
	if err != nil {
		log.Fatalf("Failed to query the version number. %s\n", err)
	}

	row = selectSQL.QueryRow(p_source_filename)
	var filename string
	err = row.Scan(&filename)
	if err != nil {
		log.Fatalf("Failed to scan the version number. %s\n", err)
	}
	log.Printf("The new filename is %s\n", filename)

	//
	//  This section creates the hard link
	//
	os.Link(filePath+"/"+source_filename, filePath+"/"+filename)
	return nil
}

// pSQL, err := db.Prepare(sqlStrg)
// defer pSQL.Close()
// rows, err := pSQL.Exec(p_source_filename, p_class, p_state, p_ifl_id, p_file_origin, p_checksum, p_file_size, p_uri_location)
// defer rows.Close()
// fmt.Printf("rows=%v\n", rows)

//for rows.Next() {
//var i int
//err = rows.Scan(&i)
// if err != nil {
// 	log.Fatal(err)
//}

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
