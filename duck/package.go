package duck

import (
	"github.com/advanced-go/stdlib/uri"
	"net/http"
)

// PROD  : https://duckduckgo.com/?q=Pascal

const (
	PkgPath     = "github/advanced-go/search/duck"
	SearchHost  = "duckduckgo.com"
	SearchPath  = "search"
	EgressRoute = "duckduckgo-search"
)

var (
	resolver = uri.NewResolver("localhost:8081")
)

// Url - egress URLs
func Url(host, path string, query any, h http.Header) string {
	return resolver.Url(host, path, query, h)
}
