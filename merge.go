package gotenberg

import (
	"context"
	"net/http"
	"strconv"

	"github.com/starwalkn/gotenberg-go-client/v8/document"
)

// MergeRequest facilitates work with PDF files with the Gotenberg API.
type MergeRequest struct {
	pdfs   []document.Document
	embeds []document.Document

	*baseRequest
	client *Client
}

func (r *MergeRequest) endpoint() string {
	return "/forms/pdfengines/merge"
}

func (r *MergeRequest) formDocuments() map[string]document.Document {
	files := make(map[string]document.Document)

	for _, pdf := range r.pdfs {
		files[pdf.Filename()] = pdf
	}

	return files
}

func (r *MergeRequest) formEmbeds() map[string]document.Document {
	embeds := make(map[string]document.Document)

	for _, embed := range r.embeds {
		embeds[embed.Filename()] = embed
	}

	return embeds
}

func (r *MergeRequest) Embeds(docs ...document.Document) *MergeRequest {
	r.embeds = append(r.embeds, docs...)
	return r
}

// PdfA sets the PDF/A format of the resulting PDF.
func (r *MergeRequest) PdfA(pdfa PdfAFormat) *MergeRequest {
	r.fields[fieldMergePdfA] = string(pdfa)
	return r
}

// PdfUA enables PDF for Universal Access for optimal accessibility.
func (r *MergeRequest) PdfUA() *MergeRequest {
	r.fields[fieldMergePdfUA] = strconv.FormatBool(true)
	return r
}

// Metadata sets the metadata to write.
func (r *MergeRequest) Metadata(md []byte) *MergeRequest {
	r.fields[fieldMetadata] = string(md)
	return r
}

// Flatten defines whether the resulting PDF should be flattened.
func (r *MergeRequest) Flatten(val bool) *MergeRequest {
	r.fields[fieldMergeFlatten] = strconv.FormatBool(val)
	return r
}

func (r *MergeRequest) Encrypt(userPassword, ownerPassword string) *MergeRequest {
	r.fields[fieldUserPassword] = userPassword
	r.fields[fieldOwnerPassword] = ownerPassword

	return r
}

func (r *MergeRequest) OutputFilename(filename string) *MergeRequest {
	r.baseRequest.OutputFilename(filename)
	return r
}

func (r *MergeRequest) Trace(trace string) *MergeRequest {
	r.baseRequest.Trace(trace)
	return r
}

func (r *MergeRequest) BasicAuth(username, password string) *MergeRequest {
	r.baseRequest.BasicAuth(username, password)
	return r
}

func (r *MergeRequest) UseWebhook(hookURL, errorURL string) *MergeRequest {
	r.baseRequest.UseWebhook(hookURL, errorURL)
	return r
}

func (r *MergeRequest) SetWebhookMethod(method string) *MergeRequest {
	r.baseRequest.SetWebhookMethod(method)
	return r
}

func (r *MergeRequest) SetWebhookErrorMethod(method string) *MergeRequest {
	r.baseRequest.SetWebhookErrorMethod(method)
	return r
}

func (r *MergeRequest) SetWebhookExtraHeaders(headers map[string]string) *MergeRequest {
	r.baseRequest.SetWebhookExtraHeaders(headers)
	return r
}

func (r *MergeRequest) DownloadFrom(downloads map[string]map[string]string, embedded bool) *MergeRequest {
	r.baseRequest.DownloadFrom(downloads, embedded)
	return r
}

func (r *MergeRequest) Store(ctx context.Context, path string) error {
	return r.client.store(ctx, r, path)
}

func (r *MergeRequest) Send(ctx context.Context) (*http.Response, error) {
	return r.client.send(ctx, r)
}

// Compile-time checks to ensure type implements desired interfaces.
var (
	_ = MultipartRequest(new(MergeRequest))
)
