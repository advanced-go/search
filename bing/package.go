package bing

import (
	"github.com/advanced-go/stdlib/uri"
	"net/http"
)

// TEST  : https://www.bing.com/search?q=C+Language

const (
	PkgPath     = "github/advanced-go/search/bing"
	SearchHost  = "www.bing.com"
	SearchPath  = "search"
	EgressRoute = "bing-search"
)

var (
	resolver = uri.NewResolver("localhost:8081")
)

// Url - egress URLs
func Url(host, path string, query any, h http.Header) string {
	return resolver.Url(host, path, query, h)
}
