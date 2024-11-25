package gotenberg

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/runatal/gotenberg-go-client/v8/document"
)

type baseRequester interface {
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

func (br *baseRequest) customHeaders() map[httpHeader]string {
	return br.headers
}

func (br *baseRequest) formFields() map[formField]string {
	return br.fields
}

// OutputFilename overrides the default UUID output filename.
//
// NOTE: Gotenberg adds the file extension automatically; you don't have to set it.
func (br *baseRequest) OutputFilename(filename string) {
	br.headers[headerOutputFilename] = filename
}

// Trace overrides the default UUID trace, or request ID, that identifies a request in Gotenberg's logs.
func (br *baseRequest) Trace(trace string) {
	br.headers[headerTrace] = trace
}

// UseBasicAuth sets the basic authentication credentials.
func (br *baseRequest) UseBasicAuth(username, password string) {
	auth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
	br.headers[headerAuthorization] = "Basic " + auth
}

// UseWebhook sets the callback and error callback that Gotenberg will use to send
// respectively the output file and the error response.
func (br *baseRequest) UseWebhook(hookURL string, errorURL string) {
	br.headers[headerWebhookURL] = hookURL
	br.headers[headerWebhookErrorURL] = errorURL
}

// SetWebhookMethod Overrides the default HTTP method that Gotenberg will use to call the webhook.
func (br *baseRequest) SetWebhookMethod(method string) {
	br.headers[headerWebhookMethod] = ensureWebhookMethod(method)
}

// SetWebhookErrorMethod overrides the default HTTP method that Gotenberg will use to call the error webhook.
func (br *baseRequest) SetWebhookErrorMethod(method string) {
	br.headers[headerWebhookErrorMethod] = ensureWebhookMethod(method)
}

// SetWebhookExtraHeaders sets the extra HTTP headers that Gotenberg will send alongside the
// request to the webhook and error webhook.
func (br *baseRequest) SetWebhookExtraHeaders(headers http.Header) error {
	marshaledHeaders, err := json.Marshal(headers)
	if err != nil {
		return fmt.Errorf("marshal headers to JSON: %w", err)
	}

	br.headers[headerWebhookExtraHeaders] = string(marshaledHeaders)

	return nil
}

func hasWebhook(req baseRequester) bool {
	url, ok := req.customHeaders()[headerWebhookURL]
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
func (br *baseRequest) DownloadFrom(downloads map[string]map[string]string) {
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

	br.fields[fieldDownloadFrom] = string(marshaled)
}

type downloadFrom struct {
	URL              string            `json:"url"`
	ExtraHTTPHeaders map[string]string `json:"extraHttpHeaders"`
}