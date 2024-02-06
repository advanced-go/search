package provider

import "fmt"

func ExampleLookupController() {
	key := googleControllerName
	c, err := cm.Get(key)

	fmt.Printf("test: LookupController(\"%v\") -> [duration:%v] [err:%v]\n", key, c.Duration, err)

	//Output:
	//test: LookupController("google-search") -> [duration:2s] [err:<nil>]

}
