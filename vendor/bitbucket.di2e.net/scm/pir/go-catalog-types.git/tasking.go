// UNCLASSIFIED
// (C) 2017 Altamira Technologies Corporation

package catalog // import "bitbucket.di2e.net/scm/pir/go-catalog-types.git"

type Tasking struct {
	TaskName     string        `json:"taskName,omitempty"`
	Requirements []*Requirement `json:"requirements,omitempty"`
}

// UNCLASSIFIED
