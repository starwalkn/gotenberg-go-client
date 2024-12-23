<p align="center">
    <h1 align="center">Gotenberg Go Client</h1>
    <p align="center">The Go client for interacting with a Gotenberg API. This project is a further development of the <a href="https://github.com/thecodingmachine/gotenberg-go-client">client from TheCodingMachine</a>, which does not support the new functionality since version 7 of Gotenberg.
</p>

---

|Gotenberg version |                                               Client version                                               | 
|:----------------:|:----------------------------------------------------------------------------------------------------------:|
|`8.x` **(actual)**|                         `8.8.0` **(actual)**                                 <br/>                         |                            
|`7.x`             |                                                 `<= 8.5.0`                                                 |
|`6.x`             | <a href="https://github.com/thecodingmachine/gotenberg-go-client">thecodingmachine/gotenberg-go-client</a> |

---

## Installation

To get the latest version of the client:

```zsh
$ go get github.com/starwalkn/gotenberg-go-client/v8@latest
```

## Preparing a documents

```go
package main

import (
    "net/http"
    "os"
	
    "github.com/starwalkn/gotenberg-go-client/v8"
    "github.com/starwalkn/gotenberg-go-client/v8/document"
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

## Converting HTML to PDF

> [!TIP]
> Head to the [documentation](https://gotenberg.dev/) to learn about all request parameters. For the PaperSize 
> method, you can use predefined parameters such as gotenberg.A4, gotenberg.A3 and so on. The full list of 
> predefined parameters can be found in [types file](https://github.com/dcaraxes/gotenberg-go-client/v8/blob/master/types.go).

> [!IMPORTANT]
> To use basic authorization, you must run Gotenberg with the --api-enable-basic-auth flag and have GOTENBERG_API_BASIC_AUTH_USERNAME and GOTENBERG_API_BASIC_AUTH_PASSWORD environment variables. 

```go
package main

import (
    "context"
    "net/http"
    
    "github.com/starwalkn/gotenberg-go-client/v8"
    "github.com/starwalkn/gotenberg-go-client/v8/document"
)

func main() {
    client, err := gotenberg.NewClient("localhost:3000", http.DefaultClient)

    // Creates the Gotenberg documents from a files paths.
    index, err := document.FromPath("index.html", "/path/to/file")

    // Create the HTML request.
    req := gotenberg.NewHTMLRequest(index)

    // Loading style and image from the specified urls. 
    downloads := make(map[string]map[string]string)
    downloads["http://my.style.css"] = nil
    downloads["http://my.img.gif"] = map[string]string{"X-Header": "Foo"}

    req.DownloadFrom(downloads)

    // Setting up basic auth (if needed).
    req.UseBasicAuth("username", "password")

    // Set the document parameters to request (optional).
    req.Margins(gotenberg.NoMargins)
    req.Scale(0.75)
    req.PaperSize(gotenberg.A4)

    // Skips the IDLE events for faster PDF conversion.
    req.SkipNetworkIdleEvent()

    // Store method allows you to store the resulting PDF in a particular destination.
    err := client.Store(context.Background(), req, "path/to/store.pdf")

    // If you wish to redirect the response directly to the browser, you may also use:
    resp, err := client.Send(context.Background(), req)
}

```

## Working with metadata
Reading metadata available only for PDF files, but you can write metadata to all Gotenberg supporting files.

### Writing metadata:

> [!TIP]
> You can write metadata to PDF for any request using the Metadata method.

```go
package main

import (
    "context"
    "encoding/json"
    "net/http"

    "github.com/starwalkn/gotenberg-go-client/v8"
    "github.com/starwalkn/gotenberg-go-client/v8/document"
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

    resp, err := client.Send(context.Background(), req)
}
```

### Reading metadata:

```go
package main

import (
    "context"
    "encoding/json"
    "net/http"

    "github.com/starwalkn/gotenberg-go-client/v8"
    "github.com/starwalkn/gotenberg-go-client/v8/document"
)

func main() {
    client, err := gotenberg.NewClient("localhost:3000", http.DefaultClient)

    // Prepare the files required for your conversion.
    doc, err := document.FromPath("filename.ext", "/path/to/file")
    req := gotenberg.NewReadMetadataRequest(doc)

    resp, err := client.Send(context.Background(), req)

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

## Creating screenshots

> [!NOTE]
> Screenshot creation is only available for HTML, URL and Markdown requests.

```go
package main

import (
    "context"
    "net/http"

    "github.com/starwalkn/gotenberg-go-client/v8"
    "github.com/starwalkn/gotenberg-go-client/v8/document"
)

func main() {
    client, err := gotenberg.NewClient("localhost:3000", http.DefaultClient)

    index, err := document.FromPath("index.html", "/path/to/file")

    // Create the HTML request and set the image format (optional).
    req := gotenberg.NewHTMLRequest(index)
    req.Format(gotenberg.JPEG)

    resp, err := client.Screenshot(context.Background(), req)
}

```

## PDF splitting
These queries allow you to split a PDF file page by page or at a specified interval.

### Split by pages

> [!IMPORTANT]
> When splitting a PDF file, it is important to note that specifying `req.Unify(true)` will return/save the PDF file, while `req.Unify(false)` will cause Gotenberg to return a ZIP archive with the files.

```go
package main

import (
    "context"
    "net/http"

    "github.com/starwalkn/gotenberg-go-client/v8"
    "github.com/starwalkn/gotenberg-go-client/v8/document"
)

func main() {
    client, err := gotenberg.NewClient("localhost:3000", http.DefaultClient)

    doc, err := document.FromPath("gotenberg.pdf", "/path/to/file")

    req := gotenberg.NewSplitPagesRequest(doc)
    req.Span("1-3")
    req.Unify(false)

    resp, err := client.Store(context.Background(), req)
}
```

### Split by intervals

```go
package main

import (
    "context"
    "net/http"

    "github.com/starwalkn/gotenberg-go-client/v8"
    "github.com/starwalkn/gotenberg-go-client/v8/document"
)

func main() {
    client, err := gotenberg.NewClient("localhost:3000", http.DefaultClient)

    doc, err := document.FromPath("gotenberg.pdf", "/path/to/file")

    req := gotenberg.NewSplitIntervalsRequest(doc)
    req.Span(2)

    resp, err := client.Store(context.Background(), req)
}
```

---

**For more complete usages, head to the [documentation](https://gotenberg.dev/).**
