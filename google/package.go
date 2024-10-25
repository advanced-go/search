package google

import (
	"github.com/advanced-go/common/uri"
)

// STAGE : https://www.google.com/search?q=C%2B%2B

const (
	PkgPath = "github/advanced-go/search/google"

	SearchHost  = "www.google.com"
	SearchPath  = "search"
	SearchRoute = "google-search"
)

var (
	resolver = uri.NewResolver("localhost:8081")
)
