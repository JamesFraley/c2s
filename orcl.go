package main

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func insertRow(catalogEnvlope TFRMCatalogEnvlope, db *sql.DB, fileLocations []iflLocation) error {
	insertSQL := `
declare
   a_row iots_file_master%ROWTYPE;
   p_row iots_file_master%ROWTYPE;
begin
   p_row.source_filename := :1;
   p_row.classification_text := :2;
   p_row.state := :3;
   p_row.ifl_id := :4;
   p_row.file_origin := :5;
   p_row.checksum := :6;
   p_row.file_size := :7;
   p_row.uri_location := :8;
	a_row := file_master_interface.register_file(p_row);
end;`

	var newFile ifmRecord

	meta := catalogEnvlope.Catalog.Meta
	class := (*meta).Classification
	newFile.classificationText = (*class).Marking
	newFile.sourceFilename = meta.Source.FileName
	newFile.processingState = "PROCESSED"
	newFile.fileOrigin = "APX"
	newFile.checksum = meta.Source.Md5
	newFile.fileSize = (*meta).Source.FileSize

	//
	//  Compare the filepath to the ifl table to see
	//    if there is an existing IFL_ID
	//
	loc := catalogEnvlope.Catalog.Locations
	var uriLocation string
	archiveFound := false
	for _, val := range loc {
		if val.Name == "archive" { //
			uriLocation = val.Uri
			archiveFound = true
		}
	}

	if archiveFound == false {
		log.Printf("No archive location found for: %s\n", newFile.sourceFilename)
		return nil
	}

	//
	// The full file path comes in with file:/<directory>/filename
	//    This little bit of code get it down to just the <directory>
	//
	colonPos := strings.Index(uriLocation, ":")
	fileFQN := uriLocation[colonPos+1:]          //this removes the file:/
	newFile.fullFilePath = filepath.Dir(fileFQN) //this removes the filename

	iflIDFound := false
	for _, dl := range fileLocations {
		lenBase := len(dl.absolutePathUnix)
		if newFile.fullFilePath[0:lenBase] == dl.absolutePathUnix {
			newFile.uriLocation = newFile.fullFilePath[lenBase:]
			newFile.iflID = dl.iflID
			iflIDFound = true
			break
		}
	}

	if iflIDFound == false {
		log.Printf("Could not find the ifl_id. %s\n", newFile.fullFilePath)
		return nil
	}

	//
	//   This section ensures the file hasn't already been processed.  That is an error.
	//
	selectSQL, err := db.Prepare("select nvl(sum(1),0) from iots_file_master where source_filename=:1")
	if err != nil {
		log.Fatalf("Failed to Prepare: \"select sum(1) ROWCOUNT from iots_file_master where source_filename=:1\": %s\n", err)
	}
	row := selectSQL.QueryRow(newFile.sourceFilename)
	selectSQL.Close()
	var rowCount int
	err = row.Scan(&rowCount)
	if err != nil {
		log.Fatalf("Failed to retrieve rowcount for %s. err=%s\n", newFile.sourceFilename, err)
	}
	if rowCount > 0 {
		log.Printf("%s has been previously processed", newFile.sourceFilename)
		return nil
	}

	//
	// This section inserts the record into the database
	//
	//_, err = db.Exec(insertSQL, p_source_filename,      p_classification_text,      p_state,                 p_ifl_id,      p_file_origin,       p_checksum,       p_file_size,      fileLocation)
	_, err = db.Exec(insertSQL, newFile.sourceFilename, newFile.classificationText, newFile.processingState, newFile.iflID, "APX", newFile.checksum, newFile.fileSize, newFile.uriLocation)
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

	row = selectSQL.QueryRow(newFile.sourceFilename)
	var filenameWithVersion string
	err = row.Scan(&filenameWithVersion)
	if err != nil {
		selectSQL.Close()
		log.Fatalf("Failed to scan the version number. %s\n", err)
	}
	selectSQL.Close()
	newFile.fnWithVersion = filenameWithVersion
	log.Printf("The new filename is %s\n", filenameWithVersion)

	//
	//  This section creates the hard link
	//
	err = os.Link(newFile.fullFilePath+"/"+newFile.sourceFilename, newFile.fullFilePath+"/"+filenameWithVersion)
	if err != nil {
		log.Printf("Failed to link to new filename")
	}
	log.Print(newFile)
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
