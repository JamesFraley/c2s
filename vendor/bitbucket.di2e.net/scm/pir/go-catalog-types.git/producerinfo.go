// UNCLASSIFIED
// (C) 2017 Altamira Technologies Corporation

package catalog // import "bitbucket.di2e.net/scm/pir/go-catalog-types.git"

type ProducerInfo struct {
	Name string `json:"name"`
	// not clear whether required or not
	Node    string `json:"node,omitempty"`
	Version string `json:"version,omitempty"`
}

// UNCLASSIFIED
