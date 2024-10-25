package duck

import (
	"github.com/advanced-go/common/uri"
)

// PROD  : https://duckduckgo.com/?q=Pascal

const (
	PkgPath     = "github/advanced-go/search/duck"
	SearchHost  = "duckduckgo.com"
	SearchPath  = "search"
	SearchRoute = "duckduckgo-search"
)

var (
	resolver = uri.NewResolver("localhost:8081")
)
