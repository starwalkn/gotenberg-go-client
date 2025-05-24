package gotenberg

import "github.com/starwalkn/gotenberg-go-client/v8/document"

const endpointMetadataWrite = "/forms/pdfengines/metadata/write"

type WriteMetadataRequest struct {
	docs []document.Document

	*baseRequest
}

func NewWriteMetadataRequest(docs ...document.Document) *WriteMetadataRequest {
	return &WriteMetadataRequest{
		docs:        docs,
		baseRequest: newBaseRequest(),
	}
}

func (r *WriteMetadataRequest) endpoint() string {
	return endpointMetadataWrite
}

func (r *WriteMetadataRequest) formDocuments() map[string]document.Document {
	docs := make(map[string]document.Document)

	for _, doc := range r.docs {
		docs[doc.Name()] = doc
	}

	return docs
}

func (r *WriteMetadataRequest) Metadata(md []byte) {
	r.fields[fieldMetadata] = string(md)
}

// Compile-time checks to ensure type implements desired interfaces.
var (
	_ = multipartRequest(new(WriteMetadataRequest))
)
