<p align="center">
    <h1 align="center">Gotenberg Go Client</h1>
    <p align="center">The Go client for interacting with a Gotenberg API. This project is a further development of the client 
        github.com/thecodingmachine/gotenberg-go-client, which does not support the functionality of version 8 and 
        newer.
    </p>
</p>

---

## Installation

> [!IMPORTANT]
> Sometimes beta versions are released, which may contain both new functionality and changes incompatible with previous versions of the client. When installing beta versions, read the release notes carefully!

To get the latest version of the client:

```zsh
$ go get github.com/dcaraxes/gotenberg-go-client@latest
```

## Usage examples

### Create a client and prepare the files

```go
package main

import (
    "os"
	
    "github.com/dcaraxes/gotenberg-go-client"
    "github.com/dcaraxes/gotenberg-go-client/document"
)

func main() {
    // Create the Gotenberg client.
    client, err := gotenberg.NewClient("localhost:3000", http.DefaultClient)

    // There are several ways to create documents.
    f1, err := document.FromPath("data.pdf", "/path/to/file")
    f2, err := document.FromString("index.html", "<html>Foo</html>")
    f3, err := document.FromBytes("index.html", []byte("<html>Foo</html>"))

    r, err := os.Open("index.html")
    f4, err := document.FromReader("index.html", r)
}
```

### Converting PDF to HTML

> [!TIP]
> Head to the [documentation](https://gotenberg.dev/) to learn about all request parameters.

```go
package main

import (
    "net/http"
    
    "github.com/dcaraxes/gotenberg-go-client"
    "github.com/dcaraxes/gotenberg-go-client/document"
)

func main() {
    client, err := gotenberg.NewClient("localhost:3000", http.DefaultClient)

    // Creates the Gotenberg documents from a files paths.
    index, err := document.FromPath("index.html", "/path/to/file")
    style, err := document.FromPath("style.css", "/path/to/file")
    img, err := document.FromPath("img.png", "/path/to/file")

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
    resp, err := client.Send(req)
}

```

### Working with metadata
Reading metadata available only for PDF files, but you can write metadata to all Gotenberg supporting files.

#### Writing metadata

> [!TIP]
> You can write metadata to PDF for any request using the Metadata method.

```go
package main

import (
    "net/http"

    "github.com/dcaraxes/gotenberg-go-client"
    "github.com/dcaraxes/gotenberg-go-client/document"
)

func main() {
    client, err := gotenberg.NewClient("localhost:3000", http.DefaultClient)
	
    // Prepare the files required for your conversion.
    doc, err := document.FromPath("filename.ext", "/path/to/file")
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

    md, err := json.Marshal(data)
    req.Metadata(md)

    resp, err := client.Send(req)
}
```

#### Reading metadata

```go
package main

import (
    "encoding/json"
    "net/http"

    "github.com/dcaraxes/gotenberg-go-client"
    "github.com/dcaraxes/gotenberg-go-client/document"
)

func main() {
    client, err := gotenberg.NewClient("localhost:3000", http.DefaultClient)

    // Prepare the files required for your conversion.
    doc, err := document.FromPath("filename.ext", "/path/to/file")
    req := gotenberg.NewReadMetadataRequest(doc)

    resp, err := client.Send(req)

    var data = struct {
        FooPdf struct {
            Author    string `json:"Author"`
            Copyright string `json:"Copyright"`
        } `json:"foo.pdf"`
    }

    // Decode metadata into a struct.
    err = json.NewDecoder(resp.Body).Decode(&data)
}

```

### Creating screenshots

> [!NOTE]
> Screenshot creation is only available for HTML, URL and Markdown requests.

```go
package main

import (
    "net/http"

    "github.com/dcaraxes/gotenberg-go-client"
    "github.com/dcaraxes/gotenberg-go-client/document"
)

func main() {
    c, err := gotenberg.NewClient("localhost:3000", http.DefaultClient)

    index, err := document.FromPath("index.html", "/path/to/file")

    // Create the HTML request and set the image format (optional).
    req := gotenberg.NewHTMLRequest(index)
    req.Format(gotenberg.JPEG)

    resp, err := client.Screenshot(req)
}

```

---

**For more complete usages, head to the [documentation](https://gotenberg.dev/).**
