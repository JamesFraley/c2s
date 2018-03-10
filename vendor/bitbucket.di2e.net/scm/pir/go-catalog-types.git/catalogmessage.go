// UNCLASSIFIED
// (C) 2017 Altamira Technologies Corporation

package catalog // import "bitbucket.di2e.net/scm/pir/go-catalog-types.git"

type CatalogMessage struct {
	Id        string      `json:"id"`
	Version   string      `json:"string"`
	DataType  string      `json:"dataType"`
	Meta      *Meta       `json:"meta"`
	Locations []*Location `json:"locations"`
}

// UNCLASSIFIED
