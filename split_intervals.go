package gotenberg

import (
	"strconv"

	"github.com/starwalkn/gotenberg-go-client/v8/document"
)

type SplitIntervalsRequest struct {
	pdfs   []document.Document
	embeds []document.Document

	*baseRequest
}

func NewSplitIntervalsRequest(pdfs ...document.Document) *SplitIntervalsRequest {
	br := newBaseRequest()
	br.fields[fieldSplitMode] = splitModeIntervals

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

func (req *SplitIntervalsRequest) formEmbeds() map[string]document.Document {
	embeds := make(map[string]document.Document)

	for _, embed := range req.embeds {
		embeds[embed.Filename()] = embed
	}

	return embeds
}

func (req *SplitIntervalsRequest) Embeds(docs ...document.Document) {
	req.embeds = append(req.embeds, docs...)
}

// SplitSpan sets the interval for split.
func (req *SplitIntervalsRequest) SplitSpan(span int) {
	req.fields[fieldSplitSpan] = strconv.Itoa(span)
}

// Flatten defines whether the resulting PDF should be flattened.
func (req *SplitIntervalsRequest) Flatten(val bool) {
	req.fields[fieldSplitFlatten] = strconv.FormatBool(val)
}

func (req *SplitIntervalsRequest) Encrypt(userPassword, ownerPassword string) {
	req.fields[fieldUserPassword] = userPassword
	req.fields[fieldOwnerPassword] = ownerPassword
}
