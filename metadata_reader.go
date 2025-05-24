package gotenberg

import "github.com/starwalkn/gotenberg-go-client/v8/document"

const endpointMetadataRead = "/forms/pdfengines/metadata/read"

type ReadMetadataRequest struct {
	docs []document.Document

	*baseRequest
}

func NewReadMetadataRequest(docs ...document.Document) *ReadMetadataRequest {
	return &ReadMetadataRequest{
		docs:        docs,
		baseRequest: newBaseRequest(),
	}
}

func (r *ReadMetadataRequest) endpoint() string {
	return endpointMetadataRead
}

func (r *ReadMetadataRequest) formDocuments() map[string]document.Document {
	docs := make(map[string]document.Document)

	for _, doc := range r.docs {
		docs[doc.Name()] = doc
	}

	return docs
}

// Compile-time checks to ensure type implements desired interfaces.
var (
	_ = multipartRequest(new(ReadMetadataRequest))
)
