package main

import (
	"bitbucket.di2e.net/scm/pir/go-catalog-types.git"
)

type iflLocation struct {
	iflID            int
	absolutePathUnix string
}

type ifmRecord struct {
	sourceFilename     string //20180308_01_000001_FRM.h5
	classificationText string //UNCLASSIFIED//NK//99X9
	processingState    string //PROCESSED
	iflID              int    //1,2,3...
	fileOrigin         string //'APX'
	checksum           string //'e777d74706da00e6b0b1016e08a99896'
	fileSize           uint64 //9079400
	uriLocation        string //  /data/cbf/tfrm/2018/03/08
	fullFilePath       string // /c2s/prod/data/cbf/tfrm/2018/03/08/20180308_0953_9481185.txt
	fnWithVersion      string // 999.20180308_0953_9481185.txt
}

type (
	TFRMCatalogRecord struct {
		catalog.CatalogRecord
	}

	TFRMCatalogEnvlope struct {
		Catalog TFRMCatalogRecord `json:"catalog,omitempty"`
	}
)

type place struct {
	Name      string   `json:"name"`
	Addr      address  `json:"address,omitempty"`
	Point     point    `json:"point"`
	FavColors []string `json:"favColors,string,omitempty"`
}

type address struct {
	StreetAddr string `json:"street_address"`
	City       string `json:"city"`
	State      string `json:"state"`
	Zipcode    int    `json:"zipcode"`
}

type point struct {
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}
