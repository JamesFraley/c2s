// UNCLASSIFIED
// (C) 2017 Altamira Technologies Corporation

package catalog // import "bitbucket.di2e.net/scm/pir/go-catalog-types.git"

type Error struct {
	Process   string `json:"process"`
	Version   string `json:"version"`
	ErrorTime *Time  `json:"errorTime,omitempty"`
	Message   string `json:"message,omitempty"`
}

// UNCLASSIFIED
