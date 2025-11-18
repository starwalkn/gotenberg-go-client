package gotenberg

import (
	"context"
	"net/http"

	"github.com/starwalkn/gotenberg-go-client/v9/document"
)

type EmbedRequest struct {
	pdfs   []document.Document
	embeds []document.Document

	*baseRequest
	client *Client
}

func (r *EmbedRequest) endpoint() string {
	return "/forms/pdfengines/embed"
}

func (r *EmbedRequest) formDocuments() map[string]document.Document {
	files := make(map[string]document.Document)

	for _, pdf := range r.pdfs {
		files[pdf.Filename()] = pdf
	}

	return files
}

func (r *EmbedRequest) formEmbeds() map[string]document.Document {
	embeds := make(map[string]document.Document)

	for _, embed := range r.embeds {
		embeds[embed.Filename()] = embed
	}

	return embeds
}

func (r *EmbedRequest) OutputFilename(filename string) *EmbedRequest {
	r.baseRequest.OutputFilename(filename)
	return r
}

func (r *EmbedRequest) Trace(trace string) *EmbedRequest {
	r.baseRequest.Trace(trace)
	return r
}

func (r *EmbedRequest) BasicAuth(username, password string) *EmbedRequest {
	r.baseRequest.BasicAuth(username, password)
	return r
}

func (r *EmbedRequest) UseWebhook(hookURL, errorURL string) *EmbedRequest {
	r.baseRequest.UseWebhook(hookURL, errorURL)
	return r
}

func (r *EmbedRequest) SetWebhookMethod(method string) *EmbedRequest {
	r.baseRequest.SetWebhookMethod(method)
	return r
}

func (r *EmbedRequest) SetWebhookErrorMethod(method string) *EmbedRequest {
	r.baseRequest.SetWebhookErrorMethod(method)
	return r
}

func (r *EmbedRequest) SetWebhookExtraHeaders(headers map[string]string) *EmbedRequest {
	r.baseRequest.SetWebhookExtraHeaders(headers)
	return r
}

func (r *EmbedRequest) DownloadFrom(downloads map[string]map[string]string, embedded bool) *EmbedRequest {
	r.baseRequest.DownloadFrom(downloads, embedded)
	return r
}

func (r *EmbedRequest) Store(ctx context.Context, path string) error {
	return r.client.store(ctx, r, path)
}

func (r *EmbedRequest) Send(ctx context.Context) (*http.Response, error) {
	return r.client.send(ctx, r)
}

// Compile-time checks to ensure type implements desired interfaces.
var (
	_ = MultipartRequest(new(EmbedRequest))
)
