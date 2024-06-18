# Working with files
```golang
// From a path.
pdf, _ := gotenberg.NewDocumentFromPath("data.pdf", "/path/to/file")

// From a string.
index, _ := gotenberg.NewDocumentFromString("index.html", "<html>Foo</html>")

// From a bytes.
index, _ := gotenberg.NewDocumentFromBytes("index.html", []byte("<html>Foo</html>"))
```