package google

import (
	"github.com/advanced-go/stdlib/controller"
	"net/url"
	"time"
)

const (
	PkgPath              = "github/advanced-go/search/google"
	searchUri            = "https://www.google.com/search"
	searchControllerName = "google-search"
)

// Controllers - egress traffic controllers
var (
	Controllers = []controller.Config{
		{searchControllerName, "www.google.com", "", "", time.Second * 2},
	}
)

func buildURL(url *url.URL) string {
	if url == nil || url.Query() == nil {
		return searchUri
	}
	return searchUri + "?" + url.Query().Encode()
}
