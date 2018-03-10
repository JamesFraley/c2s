# IFM Catalog Types for Go

This library contains the New IFM Catalog Data Types and customized json output
tags. The only external dependency is on a GeoJSON library, which is included
as a vendored library.

### Code Example

```go
package main

import (
    "encoding/json"
    "fmt"
    "bitbucket.di2e.net/scm/pir/go-catalog-types.git"
    "time"
)

func main() {
    var record catalog.CatalogRecord
    now := catalog.Time(time.Now())
    record.Id = catalog.JustString("abc123")
    record.Meta = &catalog.Meta{}
    record.Meta.Updated = &now
    record.Meta.Classification = &catalog.Classification{
        Marking:        catalog.JustString("UNCLASSIFIED"),
        Classification: catalog.JustString("U"),
    }
    tmp, err := json.MarshalIndent(record, "", "  ")
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(string(tmp))
}
```
