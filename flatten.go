package gotenberg

import "github.com/starwalkn/gotenberg-go-client/v8/document"

type FlattenRequest struct {
	pdfs []document.Document

	*baseRequest
}

func NewFlattenRequest(pdfs ...document.Document) *FlattenRequest {
	return &FlattenRequest{
		pdfs:        pdfs,
		baseRequest: newBaseRequest(),
	}
}

func (req *FlattenRequest) endpoint() string {
	return "/forms/pdfengines/flatten"
}

func (req *FlattenRequest) formDocuments() map[string]document.Document {
	files := make(map[string]document.Document)

	for _, pdf := range req.pdfs {
		files[pdf.Filename()] = pdf
	}

	return files
}
