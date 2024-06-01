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
	YahooControllerName  = "yahoo"
	GoogleControllerName = "google"
)

// config - egress traffic controllers configuration
var (
	config = []controller.Config{
		{YahooControllerName, "www.search.yahoo.com", "", "", time.Second * 2},
		{GoogleControllerName, "www.google.com", "", "", time.Second * 2},
	}
)

// ControllerConfig - get the controller configuration
func ControllerConfig(ctrlName string) (controller.Config, bool) {
	return controller.GetConfig(ctrlName, config)
}
