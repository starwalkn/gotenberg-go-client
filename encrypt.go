package gotenberg

import (
	"context"
	"net/http"

	"github.com/starwalkn/gotenberg-go-client/v8/document"
)

type EncryptRequest struct {
	pdfs []document.Document

	*baseRequest
	client *Client
}

func (r *EncryptRequest) endpoint() string {
	return "/forms/pdfengines/encrypt"
}

func (r *EncryptRequest) formDocuments() map[string]document.Document {
	files := make(map[string]document.Document)

	for _, pdf := range r.pdfs {
		files[pdf.Filename()] = pdf
	}

	return files
}

// formEmbeds a stub for implementing the Request interface.
// In the future, we need to think about a more flexible client architecture.
func (r *EncryptRequest) formEmbeds() map[string]document.Document {
	return nil
}

func (r *EncryptRequest) OutputFilename(filename string) *EncryptRequest {
	r.baseRequest.OutputFilename(filename)
	return r
}

func (r *EncryptRequest) Trace(trace string) *EncryptRequest {
	r.baseRequest.Trace(trace)
	return r
}

func (r *EncryptRequest) BasicAuth(username, password string) *EncryptRequest {
	r.baseRequest.BasicAuth(username, password)
	return r
}

func (r *EncryptRequest) UseWebhook(hookURL, errorURL string) *EncryptRequest {
	r.baseRequest.UseWebhook(hookURL, errorURL)
	return r
}

func (r *EncryptRequest) SetWebhookMethod(method string) *EncryptRequest {
	r.baseRequest.SetWebhookMethod(method)
	return r
}

func (r *EncryptRequest) SetWebhookErrorMethod(method string) *EncryptRequest {
	r.baseRequest.SetWebhookErrorMethod(method)
	return r
}

func (r *EncryptRequest) SetWebhookExtraHeaders(headers map[string]string) *EncryptRequest {
	r.baseRequest.SetWebhookExtraHeaders(headers)
	return r
}

func (r *EncryptRequest) Store(ctx context.Context, path string) error {
	return r.client.store(ctx, r, path)
}

func (r *EncryptRequest) Send(ctx context.Context) (*http.Response, error) {
	return r.client.send(ctx, r)
}

// Compile-time checks to ensure type implements desired interfaces.
var (
	_ = MultipartRequest(new(EncryptRequest))
)
