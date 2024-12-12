package gotenberg

import (
	"github.com/runatal/gotenberg-go-client/v8/document"
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
	return &HTMLRequest{index, []document.Document{}, newChromiumRequest()}
}

func (req *HTMLRequest) endpoint() string {
	return endpointHTMLConvert
}

func (req *HTMLRequest) screenshotEndpoint() string {
	return endpointHTMLScreenshot
}

func (req *HTMLRequest) formDocuments() map[string]document.Document {
	files := make(map[string]document.Document)
	files["index.html"] = req.index

	if req.header != nil {
		files["header.html"] = req.header
	}
	if req.footer != nil {
		files["footer.html"] = req.footer
	}

	for _, asset := range req.assets {
		files[asset.Filename()] = asset
	}

	return files
}

// Assets sets assets form files.
func (req *HTMLRequest) Assets(assets ...document.Document) {
	req.assets = assets
}

// Compile-time checks to ensure type implements desired interfaces.
var (
	_ = MultipartRequester(new(HTMLRequest))
)
