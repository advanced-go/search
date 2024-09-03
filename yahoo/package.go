package yahoo

import (
	"github.com/advanced-go/stdlib/uri"
	"net/http"
)

const (
	PkgPath     = "github/advanced-go/search/yahoo"
	EgressRoute = "yahoo-search"
	searchHost  = "www.search.yahoo.com"
	searchPath  = "search"
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
