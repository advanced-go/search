package yahoo

import (
	"net/url"
)

const (
	PkgPath   = "github/advanced-go/search/yahoo"
	searchUri = "https://search.yahoo.com/search"
)

func buildURL(url *url.URL) string {
	if url == nil || url.Query() == nil {
		return searchUri
	}
	return searchUri + "?" + url.Query().Encode()
}
