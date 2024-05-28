package google

import (
	"net/url"
)

const (
	PkgPath   = "github/advanced-go/search/google"
	searchUri = "https://www.google.com/search"
)

func buildURL(url *url.URL) string {
	if url == nil || url.Query() == nil {
		return searchUri
	}
	return searchUri + "?" + url.Query().Encode()
}
