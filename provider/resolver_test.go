package provider

import (
	"fmt"
	"github.com/advanced-go/core/runtime"
	"net/url"
)

func Example_Init() {
	v := make(url.Values)
	v.Add(queryArg, "golang")

	uri := resolver.Build(searchTag, searchPath, v.Encode())
	fmt.Printf("test: resolver.Build-Debug(\"%v\") -> [uri:%v]\n", searchTag, uri)

	runtime.SetTestEnvironment()
	initResolver()
	v = make(url.Values)
	v.Add(queryArg, "C Language")
	uri = resolver.Build(searchTag, searchPath, v.Encode())
	fmt.Printf("test: resolver.Build-Test-(\"%v\") -> [uri:%v]\n", searchTag, uri)

	runtime.SetStageEnvironment()
	initResolver()
	v = make(url.Values)
	v.Add(queryArg, "C++")
	uri = resolver.Build(searchTag, searchPath, v.Encode())
	fmt.Printf("test: resolver.Build-Stage(\"%v\") -> [uri:%v]\n", searchTag, uri)

	runtime.SetProdEnvironment()
	initResolver()
	v = make(url.Values)
	v.Add(queryArg, "Pascal")
	uri = resolver.Build(searchTag, searchPath, v.Encode())
	fmt.Printf("test: resolver.Build-Prod-(\"%v\") -> [uri:%v]\n", searchTag, uri)

	//Output:
	//test: resolver.Build-Debug("{SEARCH}") -> [uri:https://search.yahoo.com/search?p=golang]
	//test: resolver.Build-Test-("{SEARCH}") -> [uri:https://www.bing.com/search?q=C+Language]
	//test: resolver.Build-Stage("{SEARCH}") -> [uri:https://www.google.com/search?q=C%2B%2B]
	//test: resolver.Build-Prod-("{SEARCH}") -> [uri:https://duckduckgo.com/?q=Pascal]

}

func Example_Override() {
	runtime.SetProdEnvironment()

	v := make(url.Values)
	v.Add(queryArg, "golang")

	uri := resolver.Build(searchTag, searchPath, v.Encode())
	fmt.Printf("test: resolver.Build(\"%v\") -> [uri:%v]\n", searchTag, uri)

	resolver.SetOverrides([]runtime.Pair{{searchTag, "https://www.google.com/search?q=Pascal"}})
	s := v.Encode()
	uri = resolver.Build(searchTag, searchPath, s)
	fmt.Printf("test: resolver.Build(\"%v\") -> [uri:%v]\n", searchTag, uri)

	//Output:
	//test: resolver.Build("{SEARCH}") -> [uri:https://duckduckgo.com/?q=golang]
	//test: resolver.Build("{SEARCH}") -> [uri:https://www.google.com/search?q=Pascal]

}

func Example_NewValues() {
	values := newValues(nil)

	fmt.Printf("test: newValues(nil) -> [values:%v]\n", values)

	values = make(url.Values)
	values.Add(defaultQueryArg, "golang")
	fmt.Printf("test: newValues(%v) -> [values:%v]\n", values, newValues(values))

	queryArg = yahooQueryArg
	fmt.Printf("test: newValues(%v) -> [values:%v]\n", values, newValues(values))

	//Output:
	//test: newValues(nil) -> [values:map[q:[]]]
	//test: newValues(map[q:[golang]]) -> [values:map[q:[golang]]]
	//test: newValues(map[p:[golang]]) -> [values:map[p:[golang]]]

}
