package gotenberg

import "github.com/starwalkn/gotenberg-go-client/v8/document"

const endpointMetadataRead = "/forms/pdfengines/metadata/read"

type ReadMetadataRequest struct {
	pdfs []document.Document

	*baseRequest
}

func NewReadMetadataRequest(pdfs ...document.Document) *ReadMetadataRequest {
	return &ReadMetadataRequest{
		pdfs:        pdfs,
		baseRequest: newBaseRequest(),
	}
}

func (rmd *ReadMetadataRequest) endpoint() string {
	return endpointMetadataRead
}

func (rmd *ReadMetadataRequest) formDocuments() map[string]document.Document {
	files := make(map[string]document.Document)

	for _, pdf := range rmd.pdfs {
		files[pdf.Filename()] = pdf
	}

	return files
}

// Compile-time checks to ensure type implements desired interfaces.
var (
	_ = multipartRequester(new(ReadMetadataRequest))
)
