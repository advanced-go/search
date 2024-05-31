package yahoo

import (
	"github.com/advanced-go/stdlib/controller"
	"net/url"
	"time"
)

const (
	PkgPath              = "github/advanced-go/search/yahoo"
	searchUri            = "https://search.yahoo.com/search"
	searchControllerName = "yahoo-search"
)

// Controllers - egress traffic controllers
var (
	Controllers = []controller.Config{
		{searchControllerName, "www.search.yahoo.com", "", "", time.Second * 2},
	}
)

func buildURL(url *url.URL) string {
	if url == nil || url.Query() == nil {
		return searchUri
	}
	return searchUri + "?" + url.Query().Encode()
}
