<p align="center">
    <h1 align="center">Gotenberg Go Client</h1>
    <p align="center">The Go client for interacting with a Gotenberg API. This project is a further development of the <a href="https://github.com/thecodingmachine/gotenberg-go-client">client from TheCodingMachine</a>, which does not support the new functionality since version 7 of Gotenberg.
</p>

<div align="center">

<b>Client Compatibility Table</b>

| Gotenberg version  |                                               Client version                                               | 
|:------------------:|:----------------------------------------------------------------------------------------------------------:|
| `8.x` **(actual)** |                         `9.0.0` **(actual)**                                 <br/>                         |                            
|       `7.x`        |                                                 `<= 8.5.0`                                                 |
|       `6.x`        | <a href="https://github.com/thecodingmachine/gotenberg-go-client">thecodingmachine/gotenberg-go-client</a> |

---
</div>

## Installation

To get the latest version of the client:

```zsh
$ go get github.com/starwalkn/gotenberg-go-client/v8@latest
```

## Supported Routes

| Category        | Routes                                                     |
|-----------------|------------------------------------------------------------|
| **Chromium**    | HTML, Markdown, URL → PDF, Screenshots                     |
| **LibreOffice** | Office documents → PDF                                     |
| **PDFEngines**  | Merge, Split, Flatten, Embed, Encrypt, Read/Write Metadata |

## Preparing a documents

```go
package main

import (
	"os"

	"github.com/starwalkn/gotenberg-go-client/v8"
	"github.com/starwalkn/gotenberg-go-client/v8/document"
)

func main() {
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
> predefined parameters can be found
> in [types file](https://github.com/dcaraxes/gotenberg-go-client/v8/blob/master/types.go).

> [!IMPORTANT]
> To use basic authorization, you must run Gotenberg with the --api-enable-basic-auth flag and have
> GOTENBERG_API_BASIC_AUTH_USERNAME and GOTENBERG_API_BASIC_AUTH_PASSWORD environment variables.

```go
package main

import (
	"context"
	"log"
	"net/http"

	"github.com/starwalkn/gotenberg-go-client/v8"
	"github.com/starwalkn/gotenberg-go-client/v8/document"
)

func main() {
	c := gotenberg.NewClient("localhost:3000", http.DefaultClient)

	// Creates the Gotenberg documents from a files paths.
	index, err := document.FromPath("index.html", "/path/to/file")
	if err != nil {
		// Handle error.
	}

	err = c.Chromium().HTML(index).
		// Setting up basic auth (if enabled).
		BasicAuth("username", "password").
		// Set the document parameters to request (optional).
		Margins(gotenberg.NoMargins).
		Scale(0.75).
		PaperSize(gotenberg.A4).
		// Store method allows you to store the resulting PDF in a particular destination.
		Store(context.Background(), "path/to/store.pdf")

	// If you wish to redirect the response directly to the browser, you may also use:
	resp, err := c.Chromium().HTML(index).
		// ...
		PaperSize(gotenberg.A4).
		Send(context.Background())

	// Handle response...
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
	c := gotenberg.NewClient("localhost:3000", http.DefaultClient)

	// Prepare the files required for your operation.
	doc, err := document.FromPath("filename.ext", "/path/to/file")
	if err != nil {
		// Handle error.
	}

	// Sample metadata for writing to PDF.
	data := struct {
		Author    string `json:"Author"`
		Copyright string `json:"Copyright"`
	}{
		Author:    "Author name",
		Copyright: "Copyright",
	}

	md, err := json.Marshal(data)
	if err != nil {
		// Handle error.
	}

	resp, err := c.PDFEngines().WriteMetadata(doc).
		OutputFilename("document.pdf").
		Metadata(md).
		Send(context.Background())

	// Handle response...
}
```

### Reading metadata:

```go
package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/starwalkn/gotenberg-go-client/v8/document"
)

func main() {
	c := gotenberg.NewClient("localhost:3000", http.DefaultClient)

	// Prepare the files required for your operation.
	doc, err := document.FromPath("filename.ext", "/path/to/file")
	if err != nil {
		// Handle error.
	}

	resp, err := c.PDFEngines().ReadMetadata(doc).Send(context.Background())

	var data = struct {
		DocPdf struct {
			Author    string `json:"Author"`
			Copyright string `json:"Copyright"`
		} `json:"doc.pdf"`
	}

	// Decode metadata into a structure...
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
	c := gotenberg.NewClient("localhost:3000", http.DefaultClient)

	index, err := document.FromPath("index.html", "/path/to/file")
	if err != nil {
		// Handle error.
	}

	// Handle response for stores screenshot into a specifies path with
	// StoreScreenshot(context.Background(), "path/to/store/.image.jpeg") instead of Screenshot(context.Background())
	resp, err := c.Chromium().HTML(index).
		Format(gotenberg.JPEG).
		Screenshot(context.Background())
}

```

## PDF splitting

These queries allow you to split a PDF file page by page or at a specified interval.

### Split by pages

> [!IMPORTANT]
> When splitting a PDF file, it is important to note that specifying `req.Unify(true)` will return/save the PDF file,
> while `req.Unify(false)` will cause Gotenberg to return a ZIP archive with the files.

```go
package main

import (
	"context"
	"net/http"

	"github.com/starwalkn/gotenberg-go-client/v8"
	"github.com/starwalkn/gotenberg-go-client/v8/document"
)

func main() {
	c := gotenberg.NewClient("localhost:3000", http.DefaultClient)

	doc, err := document.FromPath("gotenberg.pdf", "/path/to/file")
	if err != nil {
		// Handle error
	}

	err = c.PDFEngines().SplitPages(doc).
		Span("1-3").
		Unify(false).
		Store(context.Background(), "path/to/store.zip")
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
	c := gotenberg.NewClient("localhost:3000", http.DefaultClient)

	doc, err := document.FromPath("gotenberg.pdf", "/path/to/file")
	if err != nil {
		// Handle error.
	}

	err = c.PDFEngines().SplitIntervals(doc).
		Span(2).
		Store(context.Background(), "path/to/document.pdf")
}
```

---

**For more complete usages, head to the [documentation](https://gotenberg.dev/).**
