package yahoo

import (
	"github.com/advanced-go/stdlib/controller"
	"github.com/advanced-go/stdlib/uri"
	"time"
)

const (
	PkgPath        = "github/advanced-go/search/yahoo"
	searchHost     = "www.search.yahoo.com"
	searchResource = "search"
	routeName      = "yahoo-search"
)

var resolver = uri.NewResolver(nil)

// EgressRoute - egress traffic route configuration
func EgressRoute() *controller.Config {
	return &controller.Config{RouteName: routeName, Host: searchHost, Authority: "", LivenessPath: "", Duration: time.Second * 2}
}
