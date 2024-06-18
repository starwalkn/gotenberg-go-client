# Generating PDF from HTML
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