# Making screenshots
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