# search

Search serves as a proxy for Google search, and contains the following uniform interfaces in the package.go file:
~~~
// Identifier
const (
    PkgPath = "github/advanced-go/search/provider"
)

// HttpHandler - HTTP handler endpoint
func HttpHandler(w http.ResponseWriter, r *http.Request) {
 // implementation details	
}
~~~

Search also utilizes a resolver for expanding URI templates, and a controller for implementing timeouts and access logging.

~~~
func init() {
  resolver.SetTemplates([]uri2.Pair{{searchPath, "https://www.google.com/search?%v"}})
}

uri := resolver.Build(searchPath, values.Encode())
defer apply(ctx, &newCtx, access.NewRequest(h, http.MethodGet, uri), &resp, googleControllerName, access.StatusCode(&status))()
	
~~~

