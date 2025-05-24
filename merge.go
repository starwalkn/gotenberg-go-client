package gotenberg

import (
	"strconv"

	"github.com/starwalkn/gotenberg-go-client/v8/document"
)

// MergeRequest facilitates work with PDF files with the Gotenberg API.
type MergeRequest struct {
	docs []document.Document

	*baseRequest
}

func NewMergeRequest(docs ...document.Document) *MergeRequest {
	return &MergeRequest{
		docs:        docs,
		baseRequest: newBaseRequest(),
	}
}

func (r *MergeRequest) endpoint() string {
	return "/forms/pdfengines/merge"
}

func (r *MergeRequest) formDocuments() map[string]document.Document {
	docs := make(map[string]document.Document)

	for _, doc := range r.docs {
		docs[doc.Name()] = doc
	}

	return docs
}

// PdfA sets the PDF/A format of the resulting PDF.
func (r *MergeRequest) PdfA(pdfa PdfAFormat) {
	r.fields[fieldMergePdfA] = string(pdfa)
}

// PdfUA enables PDF for Universal Access for optimal accessibility.
func (r *MergeRequest) PdfUA() {
	r.fields[fieldMergePdfUA] = strconv.FormatBool(true)
}

// Metadata sets the metadata to write.
func (r *MergeRequest) Metadata(md []byte) {
	r.fields[fieldMetadata] = string(md)
}

// Flatten defines whether the resulting PDF should be flattened.
func (r *MergeRequest) Flatten(val bool) {
	r.fields[fieldMergeFlatten] = strconv.FormatBool(val)
}

// Compile-time checks to ensure type implements desired interfaces.
var (
	_ = multipartRequest(new(MergeRequest))
)
