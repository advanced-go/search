package yahoo

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
	resolver.SetTemplates([]uri.Attr{{searchPath, "https://search.yahoo.com/search?%v"}})
}
