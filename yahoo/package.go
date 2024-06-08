package yahoo

import (
	"github.com/advanced-go/stdlib/controller"
	"time"
)

const (
	PkgPath        = "github/advanced-go/search/yahoo"
	searchHost     = "www.search.yahoo.com"
	searchResource = "search"
	RouteName      = "yahoo-search"
)

// Route - egress traffic route configuration
func Route(routeName string) (*controller.Config, bool) {
	switch routeName {
	case RouteName:
		return &controller.Config{RouteName: RouteName, Host: searchHost, Authority: "", LivenessPath: "", Duration: time.Second * 2}, true
	default:
		return nil, false
	}
}
