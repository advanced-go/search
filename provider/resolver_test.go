package provider

import (
	"fmt"
	"github.com/advanced-go/stdlib/uri"
	"net/url"
)

const (
	queryArg = "q"
)

func ExampleBuild() {
	v := make(url.Values)
	v.Add(queryArg, "golang")

	uri := resolver.Build(searchPath, v.Encode())
	fmt.Printf("test: resolver.Build-Debug(\"%v\") -> [uri:%v]\n", searchPath, uri)

	//Output:
	//test: resolver.Build-Debug("/search?%v") -> [uri:https://www.google.com/search?q=golang]

}

func ExampleBuild_Override() {
	//runtime.SetProdEnvironment()

	v := make(url.Values)
	v.Add(queryArg, "golang")

	uri1 := resolver.Build(searchPath, v.Encode())
	fmt.Printf("test: resolver.Build(\"%v\") -> [uri:%v]\n", searchPath, uri1)

	resolver.SetTemplates([]uri.Attr{{searchPath, "https://www.google.com/search?q=Pascal"}})
	s := v.Encode()
	uri1 = resolver.Build(searchPath, s)
	fmt.Printf("test: resolver.Build(\"%v\") -> [uri:%v]\n", searchPath, uri1)

	resolver.SetTemplates([]uri.Attr{{searchPath, "file://[cwd]/providertest/resource/query-result.txt"}})
	s = v.Encode()
	uri1 = resolver.Build(searchPath, s)
	fmt.Printf("test: resolver.Build(\"%v\") -> [uri:%v]\n", searchPath, uri1)

	//Output:
	//test: resolver.Build("/search?%v") -> [uri:https://www.google.com/search?q=golang]
	//test: resolver.Build("/search?%v") -> [uri:https://www.google.com/search?q=Pascal]
	//test: resolver.Build("/search?%v") -> [uri:file://[cwd]/providertest/resource/query-result.txt]

}
