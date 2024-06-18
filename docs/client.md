# Create the client
```golang
package main

import (
	"net/http"
	"time"

	"github.com/dcaraxes/gotenberg-go-client/v8"
)

func main() {
    // Create the HTTP-client.
    httpClient := &http.Client{
		Timeout: 5*time.Second,
    }
    // Create the Gotenberg client
    client := &gotenberg.Client{Hostname: "http://localhost:3000", HTTPClient: httpClient}
}
```