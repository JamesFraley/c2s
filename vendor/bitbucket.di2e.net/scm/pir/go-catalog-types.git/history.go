// UNCLASSIFIED
// (C) 2017 Altamira Technologies Corporation

package catalog // import "bitbucket.di2e.net/scm/pir/go-catalog-types.git"

type History struct {
	Process     string `json:"process"`
	Version     string `json:"version"`
	ProcessTime *Time  `json:"processTime,omitempty"`
	Status      string `json:"status,omitempty"`
}

// UNCLASSIFIED
