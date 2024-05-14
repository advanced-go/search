package google

import (
	"github.com/advanced-go/stdlib/uri"
)

const (
	searchPath = "/search?%v"
)

var (
	resolver = uri.NewResolver()
)

func init() {
	resolver.SetTemplates([]uri.Attr{{searchPath, "https://www.google.com/search?%v"}})
}
