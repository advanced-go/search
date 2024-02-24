package provider

import (
	"fmt"
	"github.com/advanced-go/core/messaging"
	"net/http"
)

func ExampleStartupPing() {
	r, _ := http.NewRequest("", "github/advanced-go/search/provider:ping", nil)
	status := messaging.Ping(nil, r.URL)

	fmt.Printf("test: Ping() -> [status:%v]\n", status)

	//Output:
	//test: Ping() -> [status:OK]

}
