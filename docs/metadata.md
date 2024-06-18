# Read and write EXIF metadata
Reading metadata available only for PDF files, but you can write metadata to all Gotenberg supporting files.

## Write
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

## Read
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