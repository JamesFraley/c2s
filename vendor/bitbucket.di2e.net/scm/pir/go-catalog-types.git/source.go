// UNCLASSIFIED
// (C) 2017 Altamira Technologies Corporation

package catalog // import "bitbucket.di2e.net/scm/pir/go-catalog-types.git"

type Source struct {
	FileName       string        `json:"fileName,omitempty"`
	FileFormat     string        `json:"fileFormat,omitempty"`
	FileSource     string        `json:"fileSource,omitempty"`
	ProductId      string        `json:"productId,omitempty"`
	Countries      []string      `json:"countries,omitempty"`
	Sensors        []string      `json:"sensors,omitempty"`
	Requirements   []*Requirement `json:"requirements,omitempty"`
	Agency         string        `json:"agency,omitempty"`
	Provider       string        `json:"provider,omitempty"`
	DeliveryTime   *Time         `json:"deliveryTime,omitempty"`
	DiscoveredTime *Time         `json:"discoveredTime,omitempty"`
	Uri            string        `json:"uri,omitempty"`
	FileSize       uint64        `json:"fileSize,omitempty"`
	Sha1Hash       string        `json:"sha1Hash,omitempty"`
	Md5            string        `json:"md5,omitempty"`
	Name           string        `json:"name,omitempty"`
}

// UNCLASSIFIED
