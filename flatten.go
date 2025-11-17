package gotenberg

import "github.com/starwalkn/gotenberg-go-client/v8/document"

type FlattenRequest struct {
	pdfs   []document.Document
	embeds []document.Document

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

func (req *FlattenRequest) formEmbeds() map[string]document.Document {
	embeds := make(map[string]document.Document)

	for _, embed := range req.embeds {
		embeds[embed.Filename()] = embed
	}

	return embeds
}

func (req *FlattenRequest) Embeds(docs ...document.Document) {
	req.embeds = append(req.embeds, docs...)
}

// Compile-time checks to ensure type implements desired interfaces.
var (
	_ = MultipartRequester(new(FlattenRequest))
)
