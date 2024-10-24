package gotenberg

import (
	"encoding/base64"
	"net/http"
)

type baseRequester interface {
	customHeaders() map[httpHeader]string
	formFields() map[formField]string
	formDocuments() map[string]Document
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
func (br *baseRequest) SetWebhookExtraHeaders(headers []byte) {
	br.headers[headerWebhookExtraHeaders] = string(headers)
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
