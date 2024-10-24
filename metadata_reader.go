package gotenberg

type ReadMetadataRequest struct {
	pdfs []Document

	*baseRequest
}

func NewReadMetadataRequest(pdfs ...Document) *ReadMetadataRequest {
	return &ReadMetadataRequest{
		pdfs:        pdfs,
		baseRequest: newBaseRequest(),
	}
}

func (rmd *ReadMetadataRequest) endpoint() string {
	return "/forms/pdfengines/metadata/read"
}

func (rmd *ReadMetadataRequest) formDocuments() map[string]Document {
	files := make(map[string]Document)

	for _, pdf := range rmd.pdfs {
		files[pdf.Filename()] = pdf
	}

	return files
}

// Compile-time checks to ensure type implements desired interfaces.
var (
	_ = MainRequester(new(ReadMetadataRequest))
)
