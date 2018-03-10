// UNCLASSIFIED
// (C) 2017 Altamira Technologies Corporation

package catalog // import "bitbucket.di2e.net/scm/pir/go-catalog-types.git"

import (
	"fmt"
	"strings"
	"time"
)

// Customized catalog time.
type Time struct {
	time.Time
}

func (ct *Time) MarshalJSON() ([]byte, error) {
	got := time.Time(ct.Time)
	stamp := fmt.Sprintf("\"%s\"", got.In(time.UTC).Format("2006-01-02T15:04:05.000Z"))
	return []byte(stamp), nil
}

func (ct *Time) UnmarshalJSON(buf []byte) (err error) {
	ct.Time, err = time.Parse("2006-01-02T15:04:05.000Z", strings.Trim(string(buf), `"`))
	// if cannot parse using above format, attempt parsing with RFC3339
	if err != nil {
		ct.Time, err = time.Parse(time.RFC3339, strings.Trim(string(buf), `"`))
		if err != nil {
			return err
		}
	}
	return
}

// UNCLASSIFIED
