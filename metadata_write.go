package gotenberg

import (
	"context"
	"net/http"

	"github.com/starwalkn/gotenberg-go-client/v8/document"
)

const endpointMetadataWrite = "/forms/pdfengines/metadata/write"

type WriteMetadataRequest struct {
	pdfs   []document.Document
	embeds []document.Document

	*baseRequest
	client *Client
}

func (r *WriteMetadataRequest) endpoint() string {
	return endpointMetadataWrite
}

func (r *WriteMetadataRequest) formDocuments() map[string]document.Document {
	files := make(map[string]document.Document)

	for _, pdf := range r.pdfs {
		files[pdf.Filename()] = pdf
	}

	return files
}

func (r *WriteMetadataRequest) formEmbeds() map[string]document.Document {
	embeds := make(map[string]document.Document)

	for _, embed := range r.embeds {
		embeds[embed.Filename()] = embed
	}

	return embeds
}

func (r *WriteMetadataRequest) Embeds(docs ...document.Document) *WriteMetadataRequest {
	r.embeds = append(r.embeds, docs...)
	return r
}

func (r *WriteMetadataRequest) Metadata(md []byte) *WriteMetadataRequest {
	r.fields[fieldMetadata] = string(md)
	return r
}

func (r *WriteMetadataRequest) Encrypt(userPassword, ownerPassword string) *WriteMetadataRequest {
	r.fields[fieldUserPassword] = userPassword
	r.fields[fieldOwnerPassword] = ownerPassword

	return r
}

func (r *WriteMetadataRequest) OutputFilename(filename string) *WriteMetadataRequest {
	r.baseRequest.OutputFilename(filename)
	return r
}

func (r *WriteMetadataRequest) Trace(trace string) *WriteMetadataRequest {
	r.baseRequest.Trace(trace)
	return r
}

func (r *WriteMetadataRequest) BasicAuth(username, password string) *WriteMetadataRequest {
	r.baseRequest.BasicAuth(username, password)
	return r
}

func (r *WriteMetadataRequest) UseWebhook(hookURL, errorURL string) *WriteMetadataRequest {
	r.baseRequest.UseWebhook(hookURL, errorURL)
	return r
}

func (r *WriteMetadataRequest) SetWebhookMethod(method string) *WriteMetadataRequest {
	r.baseRequest.SetWebhookMethod(method)
	return r
}

func (r *WriteMetadataRequest) SetWebhookErrorMethod(method string) *WriteMetadataRequest {
	r.baseRequest.SetWebhookErrorMethod(method)
	return r
}

func (r *WriteMetadataRequest) SetWebhookExtraHeaders(headers map[string]string) *WriteMetadataRequest {
	r.baseRequest.SetWebhookExtraHeaders(headers)
	return r
}

func (r *WriteMetadataRequest) DownloadFrom(downloads map[string]map[string]string, embedded bool) *WriteMetadataRequest {
	r.baseRequest.DownloadFrom(downloads, embedded)
	return r
}

func (r *WriteMetadataRequest) Store(ctx context.Context, path string) error {
	return r.client.store(ctx, r, path)
}

func (r *WriteMetadataRequest) Send(ctx context.Context) (*http.Response, error) {
	return r.client.send(ctx, r)
}

// Compile-time checks to ensure type implements desired interfaces.
var (
	_ = MultipartRequest(new(WriteMetadataRequest))
)
