package gotenberg

import (
	"strconv"

	"github.com/starwalkn/gotenberg-go-client/v8/document"
)

const endpointMerge = "/forms/pdfengines/merge"

// MergeRequest facilitates work with PDF files with the Gotenberg API.
type MergeRequest struct {
	pdfs []document.Document

	*baseRequest
}

func NewMergeRequest(pdfs ...document.Document) *MergeRequest {
	return &MergeRequest{pdfs, newBaseRequest()}
}

func (req *MergeRequest) endpoint() string {
	return endpointMerge
}

func (req *MergeRequest) formDocuments() map[string]document.Document {
	files := make(map[string]document.Document)

	for _, pdf := range req.pdfs {
		files[pdf.Filename()] = pdf
	}

	return files
}

// PdfA sets the PDF/A format of the resulting PDF.
func (req *MergeRequest) PdfA(pdfa PdfAFormat) {
	req.fields[fieldMergePdfA] = string(pdfa)
}

// PdfUA enables PDF for Universal Access for optimal accessibility.
func (req *MergeRequest) PdfUA() {
	req.fields[fieldMergePdfUA] = strconv.FormatBool(true)
}

// Metadata sets the metadata to write.
func (req *MergeRequest) Metadata(md []byte) {
	req.fields[fieldMetadata] = string(md)
}

// Compile-time checks to ensure type implements desired interfaces.
var (
	_ = multipartRequester(new(MergeRequest))
)
