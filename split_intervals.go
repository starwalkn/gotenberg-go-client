package gotenberg

import (
	"strconv"

	"github.com/starwalkn/gotenberg-go-client/v8/document"
)

type SplitIntervalsRequest struct {
	pdfs []document.Document

	*baseRequest
}

func NewSplitIntervalsRequest(pdfs ...document.Document) *SplitIntervalsRequest {
	br := newBaseRequest()
	br.fields[fieldSplitMode] = "intervals"

	return &SplitIntervalsRequest{
		pdfs:        pdfs,
		baseRequest: br,
	}
}

func (req *SplitIntervalsRequest) endpoint() string {
	return "/forms/pdfengines/split"
}

func (req *SplitIntervalsRequest) formDocuments() map[string]document.Document {
	files := make(map[string]document.Document)

	for _, pdf := range req.pdfs {
		files[pdf.Filename()] = pdf
	}

	return files
}

// SplitSpan sets the interval for split.
func (req *SplitIntervalsRequest) SplitSpan(span int) {
	req.fields[fieldSplitSpan] = strconv.Itoa(span)
}
