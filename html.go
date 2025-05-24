package gotenberg

import (
	"github.com/starwalkn/gotenberg-go-client/v8/document"
)

const (
	endpointHTMLConvert    = "/forms/chromium/convert/html"
	endpointHTMLScreenshot = "/forms/chromium/screenshot/html"
)

// HTMLRequest facilitates HTML conversion with the Gotenberg API.
type HTMLRequest struct {
	index  document.Document
	assets []document.Document

	*chromiumRequest
}

func NewHTMLRequest(index document.Document) *HTMLRequest {
	return &HTMLRequest{
		index:           index,
		assets:          []document.Document{},
		chromiumRequest: newChromiumRequest()}
}

func (r *HTMLRequest) endpoint() string {
	return endpointHTMLConvert
}

func (r *HTMLRequest) screenshotEndpoint() string {
	return endpointHTMLScreenshot
}

func (r *HTMLRequest) formDocuments() map[string]document.Document {
	files := make(map[string]document.Document)
	files["index.html"] = r.index

	if r.header != nil {
		files["header.html"] = r.header
	}
	if r.footer != nil {
		files["footer.html"] = r.footer
	}

	for _, asset := range r.assets {
		files[asset.Name()] = asset
	}

	return files
}

// Assets set assets form files.
func (r *HTMLRequest) Assets(assets ...document.Document) {
	r.assets = assets
}

// Compile-time checks to ensure type implements desired interfaces.
var (
	_ = multipartRequest(new(HTMLRequest))
)
