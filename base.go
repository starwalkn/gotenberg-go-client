package gotenberg

import (
	"encoding/base64"

	"github.com/starwalkn/gotenberg-go-client/v8/document"
)

const (
	headerAuthorization       = "Authorization"
	headerOutputFilename      = "Gotenberg-Output-Filename"
	headerTrace               = "Gotenberg-Trace"
	headerWebhookURL          = "Gotenberg-Webhook-Url"
	headerWebhookErrorURL     = "Gotenberg-Webhook-Error-Url"
	headerWebhookMethod       = "Gotenberg-Webhook-Method"
	headerWebhookErrorMethod  = "Gotenberg-Webhook-Error-Method"
	headerWebhookExtraHeaders = "Gotenberg-Webhook-Extra-Http-Headers"
)

type Request interface {
	customHeaders() map[string]string
	formFields() map[formField]string
	formDocuments() map[string]document.Document
	formEmbeds() map[string]document.Document
}

type baseRequest struct {
	headers map[string]string
	fields  map[formField]string
}

func newBaseRequest() *baseRequest {
	return &baseRequest{
		headers: make(map[string]string),
		fields:  make(map[formField]string),
	}
}

func (br *baseRequest) customHeaders() map[string]string {
	return br.headers
}

func (br *baseRequest) formFields() map[formField]string {
	return br.fields
}

// OutputFilename overrides the default UUID output filename.
//
// NOTE: Gotenberg adds the file extension automatically; you don't have to set it.
func (br *baseRequest) OutputFilename(filename string) *baseRequest {
	br.headers[headerOutputFilename] = filename
	return br
}

// Trace overrides the default UUID trace, or request ID, that identifies a request in Gotenberg's logs.
func (br *baseRequest) Trace(trace string) *baseRequest {
	br.headers[headerTrace] = trace
	return br
}

// BasicAuth sets the basic authentication credentials.
func (br *baseRequest) BasicAuth(username, password string) *baseRequest {
	auth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
	br.headers[headerAuthorization] = "Basic " + auth

	return br
}

// UseWebhook sets the callback and error callback that Gotenberg will use to send
// respectively the output file and the error response.
func (br *baseRequest) UseWebhook(hookURL string, errorURL string) *baseRequest {
	br.headers[headerWebhookURL] = hookURL
	br.headers[headerWebhookErrorURL] = errorURL

	return br
}

// SetWebhookMethod Overrides the default HTTP method that Gotenberg will use to call the webhook.
func (br *baseRequest) SetWebhookMethod(method string) *baseRequest {
	br.headers[headerWebhookMethod] = method
	return br
}

// SetWebhookErrorMethod overrides the default HTTP method that Gotenberg will use to call the error webhook.
func (br *baseRequest) SetWebhookErrorMethod(method string) *baseRequest {
	br.headers[headerWebhookErrorMethod] = method
	return br
}

// SetWebhookExtraHeaders sets the extra HTTP headers that Gotenberg will send alongside the
// request to the webhook and error webhook.
func (br *baseRequest) SetWebhookExtraHeaders(headers map[string]string) *baseRequest {
	br.headers[headerWebhookExtraHeaders] = mustJSON(headers)
	return br
}

func hasWebhook(req Request) bool {
	url, ok := req.customHeaders()[headerWebhookURL]
	if !ok {
		return false
	}

	return url != ""
}

// DownloadFrom sets the URLs to download files from.
// This method accepts a JSON string e.g., [{"url":"http://localhost:80/","extraHttpHeaders":{"X-Foo":"Bar"}}]. For Go,
// this is equivalent to map[string]map[string]string, which this method accepts, but headers map can be nil.
//
// URLs MUST return a Content-Disposition header with a filename parameter.
func (br *baseRequest) DownloadFrom(downloads map[string]map[string]string, embedded bool) *baseRequest {
	dfs := make([]downloadFrom, 0, len(downloads))

	for url, headers := range downloads {
		dfs = append(dfs, downloadFrom{
			URL:              url,
			ExtraHTTPHeaders: headers,
			Embedded:         embedded,
		})
	}

	br.fields[fieldDownloadFrom] = mustJSON(dfs)

	return br
}

type downloadFrom struct {
	URL              string            `json:"url"`
	ExtraHTTPHeaders map[string]string `json:"extraHttpHeaders"`
	Embedded         bool              `json:"embedded"`
}
