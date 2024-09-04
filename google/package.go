package google

import (
	"github.com/advanced-go/stdlib/uri"
	"net/http"
)

// STAGE : https://www.google.com/search?q=C%2B%2B

const (
	PkgPath     = "github/advanced-go/search/google"
	SearchHost  = "www.google.com"
	SearchPath  = "search"
	EgressRoute = "google-search"
)

var (
	resolver = uri.NewResolver("localhost:8081")
)

// EgressRoute - egress traffic route configuration
/*
func EgressRoute() *controller2.Config {
	return &controller2.Config{RouteName: routeName, Host: searchHost, Authority: "", Duration: time.Second * 2}
}
*/

// Url - egress URLs
func Url(host, path string, query any, h http.Header) string {
	return resolver.Url(host, path, query, h)
}
