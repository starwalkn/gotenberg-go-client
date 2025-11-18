package gotenberg

import (
	"context"
	"net/http"

	"github.com/starwalkn/gotenberg-go-client/v8/document"
)

const endpointMetadataRead = "/forms/pdfengines/metadata/read"

type ReadMetadataRequest struct {
	pdfs   []document.Document
	embeds []document.Document

	*baseRequest
	client *Client
}

func (r *ReadMetadataRequest) endpoint() string {
	return endpointMetadataRead
}

func (r *ReadMetadataRequest) formDocuments() map[string]document.Document {
	files := make(map[string]document.Document)

	for _, pdf := range r.pdfs {
		files[pdf.Filename()] = pdf
	}

	return files
}

func (r *ReadMetadataRequest) formEmbeds() map[string]document.Document {
	embeds := make(map[string]document.Document)

	for _, embed := range r.embeds {
		embeds[embed.Filename()] = embed
	}

	return embeds
}

func (r *ReadMetadataRequest) Embeds(docs ...document.Document) *ReadMetadataRequest {
	r.embeds = append(r.embeds, docs...)
	return r
}

func (r *ReadMetadataRequest) Encrypt(userPassword, ownerPassword string) *ReadMetadataRequest {
	r.fields[fieldUserPassword] = userPassword
	r.fields[fieldOwnerPassword] = ownerPassword

	return r
}

func (r *ReadMetadataRequest) OutputFilename(filename string) *ReadMetadataRequest {
	r.baseRequest.OutputFilename(filename)
	return r
}

func (r *ReadMetadataRequest) Trace(trace string) *ReadMetadataRequest {
	r.baseRequest.Trace(trace)
	return r
}

func (r *ReadMetadataRequest) BasicAuth(username, password string) *ReadMetadataRequest {
	r.baseRequest.BasicAuth(username, password)
	return r
}

func (r *ReadMetadataRequest) UseWebhook(hookURL, errorURL string) *ReadMetadataRequest {
	r.baseRequest.UseWebhook(hookURL, errorURL)
	return r
}

func (r *ReadMetadataRequest) SetWebhookMethod(method string) *ReadMetadataRequest {
	r.baseRequest.SetWebhookMethod(method)
	return r
}

func (r *ReadMetadataRequest) SetWebhookErrorMethod(method string) *ReadMetadataRequest {
	r.baseRequest.SetWebhookErrorMethod(method)
	return r
}

func (r *ReadMetadataRequest) SetWebhookExtraHeaders(headers map[string]string) *ReadMetadataRequest {
	r.baseRequest.SetWebhookExtraHeaders(headers)
	return r
}

func (r *ReadMetadataRequest) DownloadFrom(downloads map[string]map[string]string, embedded bool) *ReadMetadataRequest {
	r.baseRequest.DownloadFrom(downloads, embedded)
	return r
}

func (r *ReadMetadataRequest) Store(ctx context.Context, path string) error {
	return r.client.store(ctx, r, path)
}

func (r *ReadMetadataRequest) Send(ctx context.Context) (*http.Response, error) {
	return r.client.send(ctx, r)
}

// Compile-time checks to ensure type implements desired interfaces.
var (
	_ = MultipartRequest(new(ReadMetadataRequest))
)
