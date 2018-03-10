// UNCLASSIFIED
// (C) 2017 Altamira Technologies Corporation

package catalog // import "bitbucket.di2e.net/scm/pir/go-catalog-types.git"

type Collection struct {
	Key     string   `json:"key,omitempty"`
	Type    string   `json:"type,omitempty"`
	SubType string   `json:"subType,omitempty"`
	Sensors []*Sensor `json:"sensors,omitempty"`
}

// UNCLASSIFIED
