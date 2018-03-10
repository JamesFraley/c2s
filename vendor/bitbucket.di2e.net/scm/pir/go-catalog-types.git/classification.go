// UNCLASSIFIED
// (C) 2017 Altamira Technologies Corporation

package catalog // import "bitbucket.di2e.net/scm/pir/go-catalog-types.git"

type Classification struct {
	Marking               string   `json:"marking"`
	Classification        string   `json:"classification,omitempty"`
	OwnerProducer         string   `json:"ownerProducer,omitempty"`
	ClassificationReason  []string `json:"classificationReason,omitempty"`
	ClassifiedBy          []string `json:"classifiedBy,omitempty"`
	DeclassDate           *Time    `json:"declassDate,omitempty"`
	SciControls           []string `json:"sciControls,omitempty"`
	DisseminationControls []string `json:"disseminationControls,omitempty"`
	FgiSourceOpen         []string `json:"fgiSourceOpen,omitempty"`
	ReleasableTo          []string `json:"releasableTo,omitempty"`
}

// UNCLASSIFIED
