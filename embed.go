package gotenberg

import "github.com/starwalkn/gotenberg-go-client/v8/document"

type EmbedRequest struct {
	pdfs   []document.Document
	embeds []document.Document

	*baseRequest
}

func NewEmbedRequest(pdfs, embeds []document.Document) *EmbedRequest {
	return &EmbedRequest{
		pdfs:        pdfs,
		embeds:      embeds,
		baseRequest: newBaseRequest(),
	}
}

func (req *EmbedRequest) endpoint() string {
	return "/forms/pdfengines/embed"
}

func (req *EmbedRequest) formDocuments() map[string]document.Document {
	files := make(map[string]document.Document)

	for _, pdf := range req.pdfs {
		files[pdf.Filename()] = pdf
	}

	return files
}

func (req *EmbedRequest) formEmbeds() map[string]document.Document {
	embeds := make(map[string]document.Document)

	for _, embed := range req.embeds {
		embeds[embed.Filename()] = embed
	}

	return embeds
}

// Compile-time checks to ensure type implements desired interfaces.
var (
	_ = MultipartRequest(new(EmbedRequest))
)
