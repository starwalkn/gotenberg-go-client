<!-- TOC start (generated with https://github.com/derlin/bitdowntoc) -->

- [Gotenberg Go Client](#gotenberg-go-client)
   * [Install](#install)
   * [Usage](#usage)
      + [First steps](#first-steps)
         - [Create the client](#create-the-client)
         - [Prepare files](#prepare-files)
      + [Generating PDF from HTML](#generating-pdf-from-html)
      + [Read and write EXIF metadata](#read-and-write-exif-metadata)
         - [Write](#write)
         - [Read](#read)
      + [Making screenshots](#making-screenshots)
   * [Badges](#badges)

<!-- TOC end -->

<!-- TOC --><a name="gotenberg-go-client"></a>
# Gotenberg Go Client
**🔥 Working with Gotenberg version 8 and higher! 🔥**

A simple Go client for interacting with a Gotenberg API (forked github.com/thecodingmachine/gotenberg-go-client/v7).

<!-- TOC --><a name="install"></a>
## Install

```bash
$ go get -u github.com/dcaraxes/gotenberg-go-client/v8
```

<!-- TOC --><a name="usage"></a>
## Usage

<!-- TOC --><a name="first-steps"></a>
### First steps

<!-- TOC --><a name="create-the-client"></a>
#### Create the client
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

<!-- TOC --><a name="prepare-files"></a>
#### Prepare files
```golang
// From a path.
pdf, _ := gotenberg.NewDocumentFromPath("data.pdf", "/path/to/file")

// From a string.
index, _ := gotenberg.NewDocumentFromString("index.html", "<html>Foo</html>")

// From a bytes.
index, _ := gotenberg.NewDocumentFromBytes("index.html", []byte("<html>Foo</html>"))
```

<!-- TOC --><a name="generating-pdf-from-html"></a>
### Generating PDF from HTML
```golang
// Creates the Gotenberg documents from a files paths.
index, _ := gotenberg.NewDocumentFromPath("index.html", "/path/to/file")

header, _ := gotenberg.NewDocumentFromPath("header.html", "/path/to/file")
footer, _ := gotenberg.NewDocumentFromPath("footer.html", "/path/to/file")
style, _ := gotenberg.NewDocumentFromPath("style.css", "/path/to/file")
img, _ := gotenberg.NewDocumentFromPath("img.png", "/path/to/file")

// Create the HTML request.
req := gotenberg.NewHTMLRequest(index)
// Setting up basic auth (if needed).
req.SetBasicAuth("your_username", "your_password")

// Set the document parameters.
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

// Skips the IDLE events for faster PDF conversion.
req.SkipNetworkIdleEvent()

// Store method allows you to store the resulting PDF in a particular destination.
client.Store(req, "path/you/want/the/pdf/to/be/stored.pdf")

// If you wish to redirect the response directly to the browser, you may also use:
resp, _ := client.Post(req)
```

<!-- TOC --><a name="read-and-write-exif-metadata"></a>
### Read and write EXIF metadata
Reading metadata available only for PDF files, but you can write metadata to all Gotenberg supporting files.

<!-- TOC --><a name="write"></a>
#### Write
```golang
// Prepare the files required for your conversion.
pdfFile, err := NewDocumentFromPath("gotenberg1.pdf", test.PDFTestFilePath(t, "gotenberg.pdf"))
req := gotenberg.NewWriteMetadataRequest(pdfFile)
req.SetBasicAuth("your_username", "your_password")
// Sets result file name.
req.ResultFilename("foo.pdf")

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

resp, _ := client.Post(req)
```

<!-- TOC --><a name="read"></a>
#### Read
```golang
// Prepare the files required for your conversion.
pdfFile, err := NewDocumentFromPath("gotenberg1.pdf", test.PDFTestFilePath(t, "gotenberg.pdf"))
req := gotenberg.NewReadMetadataRequest(pdfFile)
req.SetBasicAuth("your_username", "your_password")
// Sets result filename
req.ResultFilename("foo.pdf")

// This response body contains JSON-formatted EXIF metadata.
respRead, _ := client.Post(req)

var readData = struct {
    FooPdf struct {
        Author    string `json:"Author"`
        Copyright string `json:"Copyright"`
    } `json:"foo.pdf"`
}
// Marshal metadata into a struct.
_ = json.NewDecoder(respRead.Body).Decode(&readData)
```

<!-- TOC --><a name="making-screenshots"></a>
### Making screenshots
Making screenshots only available for HTML, URL and Markdown requests.
```golang
index, _ := gotenberg.NewDocumentFromPath("index.html", "/path/to/file")

// Create the HTML request.
req := gotenberg.NewHTMLRequest(index)
req.SetBasicAuth("your_username", "your_password")
// Set image format.
req.Format(gotenberg.JPEG) // PNG, JPEG and WebP available now

// Store to path.
client.StoreScreenshot(req, "path/you/want/the/pdf/to/be/stored.jpeg")
// Or get response directly.
resp, _ := client.Screenshot(req)
```


For more complete usages, head to the [documentation](https://gotenberg.dev/).

<!-- TOC --><a name="badges"></a>
## Badges

[![GoDoc](https://godoc.org/github.com/dcaraxes/gotenberg-go-client/v8?status.svg)](https://godoc.org/github.com/dcaraxes/gotenberg-go-client/v8)
[![Go Report Card](https://goreportcard.com/badge/github.com/dcaraxes/gotenberg-go-client/v8)](https://goreportcard.com/report/github.com/dcaraxes/gotenberg-go-client/v8)
