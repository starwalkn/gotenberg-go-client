package gotenberg

import "github.com/starwalkn/gotenberg-go-client/v8/document"

const endpointMetadataWrite = "/forms/pdfengines/metadata/write"

type WriteMetadataRequest struct {
	pdfs   []document.Document
	embeds []document.Document

	*baseRequest
}

func NewWriteMetadataRequest(pdfs ...document.Document) *WriteMetadataRequest {
	return &WriteMetadataRequest{
		pdfs:        pdfs,
		baseRequest: newBaseRequest(),
	}
}

func (wmd *WriteMetadataRequest) endpoint() string {
	return endpointMetadataWrite
}

func (wmd *WriteMetadataRequest) formDocuments() map[string]document.Document {
	files := make(map[string]document.Document)

	for _, pdf := range wmd.pdfs {
		files[pdf.Filename()] = pdf
	}

	return files
}

func (wmd *WriteMetadataRequest) formEmbeds() map[string]document.Document {
	embeds := make(map[string]document.Document)

	for _, embed := range wmd.embeds {
		embeds[embed.Filename()] = embed
	}

	return embeds
}

func (wmd *WriteMetadataRequest) Embeds(docs ...document.Document) {
	wmd.embeds = append(wmd.embeds, docs...)
}

func (wmd *WriteMetadataRequest) Metadata(md []byte) {
	wmd.fields[fieldMetadata] = string(md)
}

func (wmd *WriteMetadataRequest) Encrypt(userPassword, ownerPassword string) {
	wmd.fields[fieldUserPassword] = userPassword
	wmd.fields[fieldOwnerPassword] = ownerPassword
}

// Compile-time checks to ensure type implements desired interfaces.
var (
	_ = MultipartRequester(new(WriteMetadataRequest))
)
