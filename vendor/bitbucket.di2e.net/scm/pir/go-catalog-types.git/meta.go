// UNCLASSIFIED
// (C) 2017 Altamira Technologies Corporation

package catalog // import "bitbucket.di2e.net/scm/pir/go-catalog-types.git"

import "github.com/paulmach/go.geojson"

type Meta struct {
	Targets        []*Target         `json:"targets,omitempty"`
	Site           []*Site           `json:"site,omitempty"`
	Source         *Source           `json:"source"`
	MetaProducer   *MetaProducer     `json:"metaProducer,omitempty"`
	Classification *Classification   `json:"classification,omitempty"`
	Tasking        *Tasking          `json:"tasking,omitempty"`
	Collection     *Collection       `json:"collection,omitempty"`
	Tags           []string          `json:"tags,omitempty"`
	Extensible     string            `json:"extensible,omitempty"`
	Start          *Time             `json:"start,omitempty"`
	End            *Time             `json:"end,omitempty"`
	Created        *Time             `json:"created,omitempty"`
	Updated        *Time             `json:"updated,omitempty"`
	Distribution   string            `json:"distribution,omitempty"`
	Dissemintation string            `json:"dissemintation,omitempty"`
	Comments       []string          `json:"comments,omitempty"`
	Description    string            `json:"description,omitempty"`
	Geom           *geojson.Geometry `json:"geom,omitempty"`
}

// UNCLASSIFIED
