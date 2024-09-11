package bing

import (
	"github.com/advanced-go/stdlib/uri"
)

// TEST  : https://www.bing.com/search?q=C+Language

const (
	PkgPath = "github/advanced-go/search/bing"

	SearchHost  = "www.bing.com"
	SearchPath  = "search"
	SearchRoute = "bing-search"
)

var (
	resolver = uri.NewResolver("localhost:8081")
)
