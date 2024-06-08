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
var Route = controller.Config{RouteName: RouteName, Host: searchHost, Authority: "", LivenessPath: "", Duration: time.Second * 2}
