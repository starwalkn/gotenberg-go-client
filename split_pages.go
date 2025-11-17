package gotenberg

import (
	"strconv"

	"github.com/starwalkn/gotenberg-go-client/v8/document"
)

type SplitPagesRequest struct {
	pdfs   []document.Document
	embeds []document.Document

	*baseRequest
}

func NewSplitPagesRequest(pdfs ...document.Document) *SplitPagesRequest {
	br := newBaseRequest()
	br.fields[fieldSplitMode] = splitModePages

	return &SplitPagesRequest{
		pdfs:        pdfs,
		embeds:      []document.Document{},
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

func (req *SplitPagesRequest) formEmbeds() map[string]document.Document {
	embeds := make(map[string]document.Document)

	for _, embed := range req.embeds {
		embeds[embed.Filename()] = embed
	}

	return embeds
}

func (req *SplitPagesRequest) Embeds(docs ...document.Document) {
	req.embeds = append(req.embeds, docs...)
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

func (req *SplitPagesRequest) Encrypt(userPassword, ownerPassword string) {
	req.fields[fieldUserPassword] = userPassword
	req.fields[fieldOwnerPassword] = ownerPassword
}
