package google

import (
	"github.com/advanced-go/stdlib/controller"
	"time"
)

const (
	PkgPath        = "github/advanced-go/search/google"
	searchHost     = "www.google.com"
	searchResource = "search"
	RouteName      = "google-search"
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
