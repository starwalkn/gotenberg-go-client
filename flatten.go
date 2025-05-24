package gotenberg

import "github.com/starwalkn/gotenberg-go-client/v8/document"

type FlattenRequest struct {
	docs []document.Document

	*baseRequest
}

func NewFlattenRequest(docs ...document.Document) *FlattenRequest {
	return &FlattenRequest{
		docs:        docs,
		baseRequest: newBaseRequest(),
	}
}

func (r *FlattenRequest) endpoint() string {
	return "/forms/pdfengines/flatten"
}

func (r *FlattenRequest) formDocuments() map[string]document.Document {
	files := make(map[string]document.Document)

	for _, pdf := range r.docs {
		files[pdf.Name()] = pdf
	}

	return files
}
