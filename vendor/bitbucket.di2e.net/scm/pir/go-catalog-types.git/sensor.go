// UNCLASSIFIED
// (C) 2017 Altamira Technologies Corporation

package catalog // import "bitbucket.di2e.net/scm/pir/go-catalog-types.git"

type Sensor struct {
	Name  string   `json:"name,omitempty"`
	Class string   `json:"class,omitempty"`
	Mode  string   `json:"mode,omitempty"`
	Band  []string `json:"band,omitempty"`
}

// UNCLASSIFIED
