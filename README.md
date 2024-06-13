<!-- TOC start (generated with https://github.com/derlin/bitdowntoc) -->

- [Gotenberg Go Client](#gotenberg-go-client)
   * [Install](#install)
   * [Usage](#usage)
      + [Generating PDF from HTML](#generating-pdf-from-html)
      + [Read and write EXIF metadata](#read-and-write-exif-metadata)
         - [Write](#write)
         - [Read](#read)
      + [Making screenshots](#making-screenshots)
   * [Badges](#badges)

<!-- TOC end -->

**🔥 Working with Gotenberg version 8 and higher! 🔥**

# Gotenberg Go Client

A simple Go client for interacting with a Gotenberg API (forked github.com/thecodingmachine/gotenberg-go-client/v7).

## Install

```bash
$ go get -u github.com/dcaraxes/gotenberg-go-client/v8
```

## Usage

### Generating PDF from HTML
```golang
import (
    "time"
    "net/http"

    "github.com/dcaraxes/gotenberg-go-client/v8"
)

// create the client.
httpClient := &http.Client{
    Timeout: time.Duration(5) * time.Second,
}
client := &gotenberg.Client{Hostname: "http://localhost:3000", HTTPClient: httpClient}

// prepare the files required for your conversion.

// from a path.
index, _ := gotenberg.NewDocumentFromPath("index.html", "/path/to/file")
// ... or from a string.
index, _ := gotenberg.NewDocumentFromString("index.html", "<html>Foo</html>")
// ... or from bytes.
index, _ := gotenberg.NewDocumentFromBytes("index.html", []byte("<html>Foo</html>"))

header, _ := gotenberg.NewDocumentFromPath("header.html", "/path/to/file")
footer, _ := gotenberg.NewDocumentFromPath("footer.html", "/path/to/file")
style, _ := gotenberg.NewDocumentFromPath("style.css", "/path/to/file")
img, _ := gotenberg.NewDocumentFromPath("img.png", "/path/to/file")

req := gotenberg.NewHTMLRequest(index)
// set up basic auth (if needed)
req.SetBasicAuth("your_username", "your_password")

req.Header(header)
req.Footer(footer)
req.Assets(style, img)
req.Margins(gotenberg.NoMargins)
req.Scale(0.75)
req.PaperSize(gotenberg.A4)
// Optional, you can change paper and margins size unit. For example:
paperSize := gotenberg.PaperDimensions{
	Height: 17, 
	Width: 11,
	// IN - inches. Other available units are PT (Points), PX (Pixels), 
	// MM (Millimeters), CM (Centimeters), PC (Picas).
	Unit: gotenberg.IN,
}
req.PaperSize(paperSize)

req.SkipNetworkIdleEvent() // for faster PDF generation

// store method allows you to... store the resulting PDF in a particular destination.
client.Store(req, "path/you/want/the/pdf/to/be/stored.pdf")

// if you wish to redirect the response directly to the browser, you may also use:
resp, _ := client.Post(req)
```

### Read and write EXIF metadata
Reading metadata available only for PDF files, but you can write metadata to all Gotenberg supporting files.
#### Write
```golang
import (
	"time"
	"net/http"

	"github.com/dcaraxes/gotenberg-go-client/v8"
)

httpClient := &http.Client{
Timeout: time.Duration(5) * time.Second,
}
client := &gotenberg.Client{Hostname: "http://localhost:3000", HTTPClient: httpClient}

// prepare the files required for your conversion.
pdfFile, err := NewDocumentFromPath("gotenberg1.pdf", test.PDFTestFilePath(t, "gotenberg.pdf"))
req := gotenberg.NewWriteMetadataRequest(pdfFile)
req.SetBasicAuth("your_username", "your_password") // if needed
req.ResultFilename("foo.pdf")
req.SkipNetworkIdleEvent() // for faster PDF generation

writeDataStruct := struct {
    Author    string `json:"Author"`
    Copyright string `json:"Copyright"`
}{
    Author:    "Author name",
    Copyright: "Copyright",
}
jsonMetadata, _ := json.Marshal(writeDataStruct)
req.Metadata(jsonMetadata)
_ = client.Store(req, "path/you/want/the/pdf/to/be/stored.pdf")

// if you wish to redirect the response directly to the browser, you may also use:
resp, _ := client.Post(req)
```

#### Read
```golang
import (
	"time"
	"net/http"

	"github.com/dcaraxes/gotenberg-go-client/v8"
)

httpClient := &http.Client{
Timeout: time.Duration(5) * time.Second,
}
client := &gotenberg.Client{Hostname: "http://localhost:3000", HTTPClient: httpClient}

// prepare the files required for your conversion.
pdfFile, err := NewDocumentFromPath("gotenberg1.pdf", test.PDFTestFilePath(t, "gotenberg.pdf"))
req := gotenberg.NewReadMetadataRequest(pdfFile)
req.SetBasicAuth("your_username", "your_password") // if needed
req.ResultFilename("foo.pdf")
req.SkipNetworkIdleEvent() // for faster PDF generation

respRead, _ := client.Post(req)

var readData = struct {
    FooPdf struct {
        Author    string `json:"Author"`
        Copyright string `json:"Copyright"`
    } `json:"foo.pdf"`
}
_ = json.NewDecoder(respRead.Body).Decode(&readData)
```

### Making screenshots
Making screenshots only available for HTML, URL and Markdown requests.
```go
import (
    "time"
    "net/http"

    "github.com/dcaraxes/gotenberg-go-client/v8"
)

httpClient := &http.Client{
    Timeout: time.Duration(5) * time.Second,
}
client := &gotenberg.Client{Hostname: "http://localhost:3000", HTTPClient: httpClient}

// prepare the files required for your conversion.

// from a path.
index, _ := gotenberg.NewDocumentFromPath("index.html", "/path/to/file")

req := gotenberg.NewHTMLRequest(index)
// set up basic auth (if needed)
req.SetBasicAuth("your_username", "your_password")
req.Format(gotenberg.JPEG) // PNG, JPEG and WebP available now
// store method allows you to... store the resulting PDF in a particular destination.
client.StoreScreenshot(req, "path/you/want/the/pdf/to/be/stored.jpeg")

// if you wish to redirect the response directly to the browser, you may also use:
resp, _ := client.Screenshot(req)
```


For more complete usages, head to the [documentation](https://gotenberg.dev/).

## Badges

[![GoDoc](https://godoc.org/github.com/dcaraxes/gotenberg-go-client/v8?status.svg)](https://godoc.org/github.com/dcaraxes/gotenberg-go-client/v8)
[![Go Report Card](https://goreportcard.com/badge/github.com/dcaraxes/gotenberg-go-client/v8)](https://goreportcard.com/report/github.com/dcaraxes/gotenberg-go-client/v8)
