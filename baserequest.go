package gotenberg

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/starwalkn/gotenberg-go-client/v8/document"
)

type httpHeader string

const (
	headerOutputFilename httpHeader = "Gotenberg-Output-Filename"
	headerTrace          httpHeader = "Gotenberg-Trace"
	headerAuthorization  httpHeader = "Authorization"

	headerWebhookURL          httpHeader = "Gotenberg-Webhook-Url"
	headerWebhookErrorURL     httpHeader = "Gotenberg-Webhook-Error-Url"
	headerWebhookMethod       httpHeader = "Gotenberg-Webhook-Method"
	headerWebhookErrorMethod  httpHeader = "Gotenberg-Webhook-Error-Method"
	headerWebhookExtraHeaders httpHeader = "Gotenberg-Webhook-Extra-Http-Headers"
)

type baseRequester interface {
	hasWebhook() bool
	customHeaders() map[httpHeader]string
	formFields() map[formField]string
	formDocuments() map[string]document.Document
}

type baseRequest struct {
	headers map[httpHeader]string
	fields  map[formField]string
}

func newBaseRequest() *baseRequest {
	return &baseRequest{
		headers: make(map[httpHeader]string),
		fields:  make(map[formField]string),
	}
}

func (r *baseRequest) customHeaders() map[httpHeader]string {
	return r.headers
}

func (r *baseRequest) formFields() map[formField]string {
	return r.fields
}

// OutputFilename overrides the default UUID output filename.
//
// NOTE: Gotenberg adds the file extension automatically; you don't have to set it.
func (r *baseRequest) OutputFilename(filename string) {
	r.headers[headerOutputFilename] = filename
}

// Trace overrides the default UUID trace, or request ID, that identifies a request in Gotenberg's logs.
func (r *baseRequest) Trace(trace string) {
	r.headers[headerTrace] = trace
}

// UseBasicAuth sets the basic authentication credentials.
func (r *baseRequest) UseBasicAuth(username, password string) {
	auth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
	r.headers[headerAuthorization] = "Basic " + auth
}

// UseWebhook sets the callback and error callback that Gotenberg will use to send
// respectively the output file and the error response.
func (r *baseRequest) UseWebhook(hookURL string, errorURL string) {
	r.headers[headerWebhookURL] = hookURL
	r.headers[headerWebhookErrorURL] = errorURL
}

// SetWebhookMethod Overrides the default HTTP method that Gotenberg will use to call the webhook.
func (r *baseRequest) SetWebhookMethod(method string) {
	r.headers[headerWebhookMethod] = ensureWebhookMethod(method)
}

// SetWebhookErrorMethod overrides the default HTTP method that Gotenberg will use to call the error webhook.
func (r *baseRequest) SetWebhookErrorMethod(method string) {
	r.headers[headerWebhookErrorMethod] = ensureWebhookMethod(method)
}

// SetWebhookExtraHeaders sets the extra HTTP headers that Gotenberg will send alongside the
// request to the webhook and error webhook.
func (r *baseRequest) SetWebhookExtraHeaders(headers map[string]string) error {
	marshaledHeaders, err := json.Marshal(headers)
	if err != nil {
		return fmt.Errorf("marshal headers to JSON: %w", err)
	}

	r.headers[headerWebhookExtraHeaders] = string(marshaledHeaders)

	return nil
}

func (r *baseRequest) hasWebhook() bool {
	url, ok := r.headers[headerWebhookURL]
	if !ok {
		return false
	}

	return url != ""
}

func ensureWebhookMethod(method string) string {
	if method == http.MethodPut || method == http.MethodPost || method == http.MethodPatch {
		return method
	} else {
		return http.MethodGet
	}
}

// DownloadFrom sets the URLs to download files from.
// This method accepts a JSON string e.g., [{"url":"http://localhost:80/","extraHttpHeaders":{"X-Foo":"Bar"}}]. For Go,
// this is equivalent to map[string]map[string]string, which this method accepts, but headers map can be nil.
//
// URLs MUST return a Content-Disposition header with a filename parameter.
func (r *baseRequest) DownloadFrom(downloads map[string]map[string]string) {
	dfs := make([]downloadFrom, 0, len(downloads))

	for url, headers := range downloads {
		dfs = append(dfs, downloadFrom{
			URL:              url,
			ExtraHTTPHeaders: headers,
		})
	}

	marshaled, err := json.Marshal(dfs)
	if err != nil {
		return
	}

	r.fields[fieldDownloadFrom] = string(marshaled)
}

type downloadFrom struct {
	URL              string            `json:"url"`
	ExtraHTTPHeaders map[string]string `json:"extraHttpHeaders"`
}
