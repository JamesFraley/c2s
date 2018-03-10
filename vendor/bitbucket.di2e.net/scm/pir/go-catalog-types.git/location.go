// UNCLASSIFIED
// (C) 2017 Altamira Technologies Corporation

package catalog // import "bitbucket.di2e.net/scm/pir/go-catalog-types.git"

type Location struct {
	Name        string `json:"name"`
	Uri         string `json:"uri"`
	FileSize    uint64 `json:"fileSize,omitempty"`
	ValidStart  *Time  `json:"validStart,omitempty"`
	ValidEnd    *Time  `json:"validEnd,omitempty"`
	Compression string `json:"compression,omitempty"`
	Md5         string `json:"md5,omitempty"`
	Sha1        string `json:"sha1,omitempty"`
}

// UNCLASSIFIED
