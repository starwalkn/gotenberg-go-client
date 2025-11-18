package gotenberg

import (
	"context"
	"net/http"

	"github.com/starwalkn/gotenberg-go-client/v9/document"
)

type FlattenRequest struct {
	pdfs   []document.Document
	embeds []document.Document

	*baseRequest
	client *Client
}

func (r *FlattenRequest) endpoint() string {
	return "/forms/pdfengines/flatten"
}

func (r *FlattenRequest) formDocuments() map[string]document.Document {
	files := make(map[string]document.Document)

	for _, pdf := range r.pdfs {
		files[pdf.Filename()] = pdf
	}

	return files
}

func (r *FlattenRequest) formEmbeds() map[string]document.Document {
	embeds := make(map[string]document.Document)

	for _, embed := range r.embeds {
		embeds[embed.Filename()] = embed
	}

	return embeds
}

func (r *FlattenRequest) Embeds(docs ...document.Document) *FlattenRequest {
	r.embeds = append(r.embeds, docs...)
	return r
}

func (r *FlattenRequest) OutputFilename(filename string) *FlattenRequest {
	r.baseRequest.OutputFilename(filename)
	return r
}

func (r *FlattenRequest) Trace(trace string) *FlattenRequest {
	r.baseRequest.Trace(trace)
	return r
}

func (r *FlattenRequest) BasicAuth(username, password string) *FlattenRequest {
	r.baseRequest.BasicAuth(username, password)
	return r
}

func (r *FlattenRequest) UseWebhook(hookURL, errorURL string) *FlattenRequest {
	r.baseRequest.UseWebhook(hookURL, errorURL)
	return r
}

func (r *FlattenRequest) SetWebhookMethod(method string) *FlattenRequest {
	r.baseRequest.SetWebhookMethod(method)
	return r
}

func (r *FlattenRequest) SetWebhookErrorMethod(method string) *FlattenRequest {
	r.baseRequest.SetWebhookErrorMethod(method)
	return r
}

func (r *FlattenRequest) SetWebhookExtraHeaders(headers map[string]string) *FlattenRequest {
	r.baseRequest.SetWebhookExtraHeaders(headers)
	return r
}

func (r *FlattenRequest) DownloadFrom(downloads map[string]map[string]string, embedded bool) *FlattenRequest {
	r.baseRequest.DownloadFrom(downloads, embedded)
	return r
}

func (r *FlattenRequest) Store(ctx context.Context, path string) error {
	return r.client.store(ctx, r, path)
}

func (r *FlattenRequest) Send(ctx context.Context) (*http.Response, error) {
	return r.client.send(ctx, r)
}

// Compile-time checks to ensure type implements desired interfaces.
var (
	_ = MultipartRequest(new(FlattenRequest))
)
