package gotenberg

import "github.com/starwalkn/gotenberg-go-client/v8/document"

const endpointMetadataRead = "/forms/pdfengines/metadata/read"

type ReadMetadataRequest struct {
	pdfs   []document.Document
	embeds []document.Document

	*baseRequest
}

func NewReadMetadataRequest(pdfs ...document.Document) *ReadMetadataRequest {
	return &ReadMetadataRequest{
		pdfs:        pdfs,
		embeds:      []document.Document{},
		baseRequest: newBaseRequest(),
	}
}

func (rmd *ReadMetadataRequest) endpoint() string {
	return endpointMetadataRead
}

func (rmd *ReadMetadataRequest) formDocuments() map[string]document.Document {
	files := make(map[string]document.Document)

	for _, pdf := range rmd.pdfs {
		files[pdf.Filename()] = pdf
	}

	return files
}

func (rmd *ReadMetadataRequest) formEmbeds() map[string]document.Document {
	embeds := make(map[string]document.Document)

	for _, embed := range rmd.embeds {
		embeds[embed.Filename()] = embed
	}

	return embeds
}

func (rmd *ReadMetadataRequest) Embeds(docs ...document.Document) {
	rmd.embeds = append(rmd.embeds, docs...)
}

func (rmd *ReadMetadataRequest) Encrypt(userPassword, ownerPassword string) {
	rmd.fields[fieldUserPassword] = userPassword
	rmd.fields[fieldOwnerPassword] = ownerPassword
}

// Compile-time checks to ensure type implements desired interfaces.
var (
	_ = MultipartRequester(new(ReadMetadataRequest))
)
