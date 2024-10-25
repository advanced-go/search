package yahoo

import (
	"github.com/advanced-go/common/uri"
)

// DEBUG : https://search.yahoo.com/search?p=golang

const (
	PkgPath = "github/advanced-go/search/yahoo"

	SearchHost  = "www.search.yahoo.com"
	SearchPath  = "search"
	SearchRoute = "yahoo-search"
)

var (
	resolver = uri.NewResolver("localhost:8081")
)
