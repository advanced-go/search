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

