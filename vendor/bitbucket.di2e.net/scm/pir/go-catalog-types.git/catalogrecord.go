// UNCLASSIFIED
// (C) 2017 Altamira Technologies Corporation

package catalog // import "bitbucket.di2e.net/scm/pir/go-catalog-types.git"

type CatalogRecord struct {
	Id             string       `json:"id"`
	SchemaVersion  string       `json:"schemaVersion"`
	DataType       string       `json:"dataType,omitempty"`
	SearchDataType []string     `json:"searchDataType,omitempty"`
	Permissions    *Permissions `json:"permissions,omitempty"`
	GspReleasable  bool         `json:"gspReleasable,omitempty"`
	Meta           *Meta        `json:"meta,omitempty"`
	Locations      []*Location  `json:"locations,omitempty"`
	History        []*History   `json:"history,omitempty"`
	Errors         []*Error     `json:"errors,omitempty"`
}

// UNCLASSIFIED
