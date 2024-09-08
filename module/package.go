package module

const (
	Authority    = "github/advanced-go/search"
	RouteName    = "search"
	Version      = "1.1.1"
	VersionRoute = "version"
)

// Configuration keys used on startup for map values
const (
	PackageNameUserKey     = "user"    // type:text, empty:false
	PackageNamePasswordKey = "pswd"    // type:text, empty:false
	PackageNameRetriesKey  = "retries" // type:int, default:-1, range:0-100
)
