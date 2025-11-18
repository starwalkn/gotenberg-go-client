package gotenberg

import (
	"context"
	"net/http"
	"strconv"

	"github.com/starwalkn/gotenberg-go-client/v9/document"
)

type SplitPagesRequest struct {
	pdfs   []document.Document
	embeds []document.Document

	*baseRequest
	client *Client
}

func (r *SplitPagesRequest) endpoint() string {
	return "/forms/pdfengines/split"
}

func (r *SplitPagesRequest) formDocuments() map[string]document.Document {
	files := make(map[string]document.Document)

	for _, pdf := range r.pdfs {
		files[pdf.Filename()] = pdf
	}

	return files
}

func (r *SplitPagesRequest) formEmbeds() map[string]document.Document {
	embeds := make(map[string]document.Document)

	for _, embed := range r.embeds {
		embeds[embed.Filename()] = embed
	}

	return embeds
}

func (r *SplitPagesRequest) Embeds(docs ...document.Document) *SplitPagesRequest {
	r.embeds = append(r.embeds, docs...)
	return r
}

// Span sets the interval for split.
func (r *SplitPagesRequest) Span(span string) *SplitPagesRequest {
	r.fields[fieldSplitSpan] = span
	return r
}

func (r *SplitPagesRequest) Unify(val bool) *SplitPagesRequest {
	r.fields[fieldSplitUnify] = strconv.FormatBool(val)
	return r
}

// Flatten defines whether the resulting PDF should be flattened.
func (r *SplitPagesRequest) Flatten(val bool) *SplitPagesRequest {
	r.fields[fieldSplitFlatten] = strconv.FormatBool(val)
	return r
}

func (r *SplitPagesRequest) Encrypt(userPassword, ownerPassword string) *SplitPagesRequest {
	r.fields[fieldUserPassword] = userPassword
	r.fields[fieldOwnerPassword] = ownerPassword

	return r
}

func (r *SplitPagesRequest) OutputFilename(filename string) *SplitPagesRequest {
	r.baseRequest.OutputFilename(filename)
	return r
}

func (r *SplitPagesRequest) Trace(trace string) *SplitPagesRequest {
	r.baseRequest.Trace(trace)
	return r
}

func (r *SplitPagesRequest) BasicAuth(username, password string) *SplitPagesRequest {
	r.baseRequest.BasicAuth(username, password)
	return r
}

func (r *SplitPagesRequest) UseWebhook(hookURL, errorURL string) *SplitPagesRequest {
	r.baseRequest.UseWebhook(hookURL, errorURL)
	return r
}

func (r *SplitPagesRequest) SetWebhookMethod(method string) *SplitPagesRequest {
	r.baseRequest.SetWebhookMethod(method)
	return r
}

func (r *SplitPagesRequest) SetWebhookErrorMethod(method string) *SplitPagesRequest {
	r.baseRequest.SetWebhookErrorMethod(method)
	return r
}

func (r *SplitPagesRequest) SetWebhookExtraHeaders(headers map[string]string) *SplitPagesRequest {
	r.baseRequest.SetWebhookExtraHeaders(headers)
	return r
}

func (r *SplitPagesRequest) DownloadFrom(downloads map[string]map[string]string, embedded bool) *SplitPagesRequest {
	r.baseRequest.DownloadFrom(downloads, embedded)
	return r
}

func (r *SplitPagesRequest) Store(ctx context.Context, path string) error {
	return r.client.store(ctx, r, path)
}

func (r *SplitPagesRequest) Send(ctx context.Context) (*http.Response, error) {
	return r.client.send(ctx, r)
}

// Compile-time checks to ensure type implements desired interfaces.
var (
	_ = MultipartRequest(new(SplitPagesRequest))
)
