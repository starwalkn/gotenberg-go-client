package gotenberg

import (
	"context"
	"net/http"
	"strconv"

	"github.com/starwalkn/gotenberg-go-client/v8/document"
)

type SplitIntervalsRequest struct {
	pdfs   []document.Document
	embeds []document.Document

	*baseRequest
	client *Client
}

func (r *SplitIntervalsRequest) endpoint() string {
	return "/forms/pdfengines/split"
}

func (r *SplitIntervalsRequest) formDocuments() map[string]document.Document {
	files := make(map[string]document.Document)

	for _, pdf := range r.pdfs {
		files[pdf.Filename()] = pdf
	}

	return files
}

func (r *SplitIntervalsRequest) formEmbeds() map[string]document.Document {
	embeds := make(map[string]document.Document)

	for _, embed := range r.embeds {
		embeds[embed.Filename()] = embed
	}

	return embeds
}

func (r *SplitIntervalsRequest) Embeds(docs ...document.Document) *SplitIntervalsRequest {
	r.embeds = append(r.embeds, docs...)
	return r
}

// Span sets the interval for split.
func (r *SplitIntervalsRequest) Span(span int) *SplitIntervalsRequest {
	r.fields[fieldSplitSpan] = strconv.Itoa(span)
	return r
}

// Flatten defines whether the resulting PDF should be flattened.
func (r *SplitIntervalsRequest) Flatten(val bool) *SplitIntervalsRequest {
	r.fields[fieldSplitFlatten] = strconv.FormatBool(val)
	return r
}

func (r *SplitIntervalsRequest) Encrypt(userPassword, ownerPassword string) *SplitIntervalsRequest {
	r.fields[fieldUserPassword] = userPassword
	r.fields[fieldOwnerPassword] = ownerPassword

	return r
}

func (r *SplitIntervalsRequest) OutputFilename(filename string) *SplitIntervalsRequest {
	r.baseRequest.OutputFilename(filename)
	return r
}

func (r *SplitIntervalsRequest) Trace(trace string) *SplitIntervalsRequest {
	r.baseRequest.Trace(trace)
	return r
}

func (r *SplitIntervalsRequest) BasicAuth(username, password string) *SplitIntervalsRequest {
	r.baseRequest.BasicAuth(username, password)
	return r
}

func (r *SplitIntervalsRequest) UseWebhook(hookURL, errorURL string) *SplitIntervalsRequest {
	r.baseRequest.UseWebhook(hookURL, errorURL)
	return r
}

func (r *SplitIntervalsRequest) SetWebhookMethod(method string) *SplitIntervalsRequest {
	r.baseRequest.SetWebhookMethod(method)
	return r
}

func (r *SplitIntervalsRequest) SetWebhookErrorMethod(method string) *SplitIntervalsRequest {
	r.baseRequest.SetWebhookErrorMethod(method)
	return r
}

func (r *SplitIntervalsRequest) SetWebhookExtraHeaders(headers map[string]string) *SplitIntervalsRequest {
	r.baseRequest.SetWebhookExtraHeaders(headers)
	return r
}

func (r *SplitIntervalsRequest) Store(ctx context.Context, path string) error {
	return r.client.store(ctx, r, path)
}

func (r *SplitIntervalsRequest) Send(ctx context.Context) (*http.Response, error) {
	return r.client.send(ctx, r)
}
