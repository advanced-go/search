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
	RouteName      = "yahoo-search"
)

var resolver = uri.NewResolver(nil)

// EgressRoute - egress traffic route configuration
func EgressRoute(routeName string) (*controller.Config, bool) {
	switch routeName {
	case RouteName:
		return &controller.Config{RouteName: RouteName, Host: searchHost, Authority: "", LivenessPath: "", Duration: time.Second * 2}, true
	default:
		return nil, false
	}
}
