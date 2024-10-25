# Gotenberg Go Client

The Go client for interacting with a Gotenberg API. This project is a further development of the client 
github.com/thecodingmachine/gotenberg-go-client, which does not support the functionality of version 8 and 
newer.

## Install

```zsh
$ go get -u github.com/runatal/gotenberg-go-client
```

## First steps

### Create the client and prepare files

```go
package main

import (
    "os"
	
    "github.com/dcaraxes/gotenberg-go-client"
)

func main() {
	// Create the Gotenberg client.
    client, err := gotenberg.NewClient("localhost:3000", http.DefaultClient)

    // There are several ways to create documentsю
    f1, err := gotenberg.NewDocumentFromPath("data.pdf", "/path/to/file")
    f2, err := gotenberg.NewDocumentFromString("index.html", "<html>Foo</html>")
    f3, err := gotenberg.NewDocumentFromBytes("index.html", []byte("<html>Foo</html>"))
    r, err := os.Open("index.html")
    f4, err := gotenberg.NewDocumentFromReader("index.html", r)
}
```

### Generating PDF from HTML

> [!TIP]
> Head to the [documentation](https://gotenberg.dev/) to learn about all request parameters.

```go
package main

import (
	"net/http"

	"github.com/dcaraxes/gotenberg-go-client"
)

func main() {
	client, err := gotenberg.NewClient("localhost:3000", http.DefaultClient)

	// Creates the Gotenberg documents from a files paths.
	index, err := gotenberg.NewDocumentFromPath("index.html", "/path/to/file")
	style, err := gotenberg.NewDocumentFromPath("style.css", "/path/to/file")
	img, err := gotenberg.NewDocumentFromPath("img.png", "/path/to/file")

	// Create the HTML request.
	req := gotenberg.NewHTMLRequest(index)

	// Setting up basic auth (if needed).
	req.UseBasicAuth("username", "password")

	// Set the document parameters to request (optional).
	req.Assets(style, img)
	req.Margins(gotenberg.NoMargins)
	req.Scale(0.75)
	req.PaperSize(gotenberg.A4)

	// Skips the IDLE events for faster PDF conversion.
	req.SkipNetworkIdleEvent()

	// Store method allows you to store the resulting PDF in a particular destination.
	client.Store(req, "path/to/store.pdf")

	// If you wish to redirect the response directly to the browser, you may also use:
	resp, err := client.Post(req)
}

```

### Read and write EXIF metadata
Reading metadata available only for PDF files, but you can write metadata to all Gotenberg supporting files.

#### Write
> [!TIP]
> You can write metadata to PDF for any request using the Metadata method.

```go
package main

import (
	"net/http"

	"github.com/dcaraxes/gotenberg-go-client"
)

func main() {
	client, err := gotenberg.NewClient("localhost:3000", http.DefaultClient)
	
	// Prepare the files required for your conversion.
	doc, err := NewDocumentFromPath("filename.ext", "/path/to/file")
	req := gotenberg.NewWriteMetadataRequest(doc)

	// Sets result file name.
	req.OutputFilename("foo.pdf")

	data := struct {
		Author    string `json:"Author"`
		Copyright string `json:"Copyright"`
	}{
		Author:    "Author name",
		Copyright: "Copyright",
	}

	jsonMetadata, err := json.Marshal(data)
	req.Metadata(jsonMetadata)
	err = client.Store(req, "path/to/store.pdf")

	resp, err := client.Post(req)
}
```

#### Read

```go
package main

import (
	"encoding/json"
	"net/http"

	"github.com/dcaraxes/gotenberg-go-client"
)

func main() {
	client, err := gotenberg.NewClient("localhost:3000", http.DefaultClient)

	// Prepare the files required for your conversion.
	doc, err := gotenberg.NewDocumentFromPath("filename.ext", "/path/to/file")
	req := gotenberg.NewReadMetadataRequest(doc)

	// This response body contains JSON-formatted EXIF metadata.
	resp, err := client.Post(req)

	var data = struct {
		FooPdf struct {
			Author    string `json:"Author"`
			Copyright string `json:"Copyright"`
		} `json:"foo.pdf"`
	}

	// Marshal metadata into a struct.
	err = json.NewDecoder(resp.Body).Decode(&data)
}

```

### Making screenshots
Making screenshots only available for HTML, URL and Markdown requests.

```go
package main

import (
	"net/http"

	"github.com/dcaraxes/gotenberg-go-client"
)

func main() {
	c, err := gotenberg.NewClient("localhost:3000", http.DefaultClient)

	index, _ := gotenberg.NewDocumentFromPath("index.html", "/path/to/file")

	// Create the HTML request.
	req := gotenberg.NewHTMLRequest(index)

	// Set image format.
	req.Format(gotenberg.JPEG)

	// Store to path.
	client.StoreScreenshot(req, "path/to/store.jpeg")
	// Or get response directly.
	resp, err := client.Screenshot(req)
}

```

---

**For more complete usages, head to the [documentation](https://gotenberg.dev/).**
