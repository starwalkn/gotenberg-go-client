package gotenberg

import (
	"strconv"

	"github.com/starwalkn/gotenberg-go-client/v8/document"
)

type SplitPagesRequest struct {
	pdfs []document.Document

	*baseRequest
}

func NewSplitPagesRequest(pdfs ...document.Document) *SplitPagesRequest {
	br := newBaseRequest()
	br.fields[fieldSplitMode] = splitModePages

	return &SplitPagesRequest{
		pdfs:        pdfs,
		baseRequest: br,
	}
}

func (req *SplitPagesRequest) endpoint() string {
	return "/forms/pdfengines/split"
}

func (req *SplitPagesRequest) formDocuments() map[string]document.Document {
	files := make(map[string]document.Document)

	for _, pdf := range req.pdfs {
		files[pdf.Filename()] = pdf
	}

	return files
}

// SplitSpan sets the interval for split.
func (req *SplitPagesRequest) SplitSpan(span string) {
	req.fields[fieldSplitSpan] = span
}

func (req *SplitPagesRequest) SplitUnify(val bool) {
	req.fields[fieldSplitUnify] = strconv.FormatBool(val)
}

// Flatten defines whether the resulting PDF should be flattened.
func (req *SplitPagesRequest) Flatten(val bool) {
	req.fields[fieldSplitFlatten] = strconv.FormatBool(val)
}
