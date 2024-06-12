package google

import (
	"github.com/advanced-go/stdlib/controller2"
	"github.com/advanced-go/stdlib/uri"
	"time"
)

const (
	PkgPath        = "github/advanced-go/search/google"
	searchHost     = "www.google.com"
	searchResource = "search"
	routeName      = "google-search"
)

var resolver = uri.NewResolver(nil)

// EgressRoute - egress traffic route configuration
func EgressRoute() *controller2.Config {
	return &controller2.Config{RouteName: routeName, Host: searchHost, Authority: "", Duration: time.Second * 2}
}
