package module

import (
	"github.com/advanced-go/stdlib/controller"
	"time"
)

const (
	Authority = "github/advanced-go/search"
	Name      = "search"
	Version   = "1.1.1"
)

// Configuration keys used on startup for map values
const (
	PackageNameUserKey     = "user"    // type:text, empty:false
	PackageNamePasswordKey = "pswd"    // type:text, empty:false
	PackageNameRetriesKey  = "retries" // type:int, default:-1, range:0-100
)

const (
	YahooRouteName  = "yahoo-search"
	GoogleRouteName = "google-search"
)

// Routes - egress traffic route configuration
var (
	Routes = []controller.Config{
		{YahooRouteName, "www.search.yahoo.com", "", "", time.Second * 2},
		{GoogleRouteName, "www.google.com", "", "", time.Second * 2},
	}
)

// GetRoute - get the route configuration
func GetRoute(routeName string) (controller.Config, bool) {
	return controller.GetRoute(routeName, Routes)
}
