package gotenberg

import "github.com/starwalkn/gotenberg-go-client/v8/document"

const (
	endpointMarkdownConvert    = "/forms/chromium/convert/markdown"
	endpointMarkdownScreenshot = "/forms/chromium/screenshot/markdown"
)

// MarkdownRequest facilitates Markdown conversion with the Gotenberg API.
type MarkdownRequest struct {
	index     document.Document
	markdowns []document.Document
	assets    []document.Document

	*chromiumRequest
}

func NewMarkdownRequest(index document.Document, markdowns ...document.Document) *MarkdownRequest {
	return &MarkdownRequest{index, markdowns, []document.Document{}, newChromiumRequest()}
}

func (req *MarkdownRequest) endpoint() string {
	return endpointMarkdownConvert
}

func (req *MarkdownRequest) screenshotEndpoint() string {
	return endpointMarkdownScreenshot
}

func (req *MarkdownRequest) formDocuments() map[string]document.Document {
	files := make(map[string]document.Document)
	files["index.html"] = req.index
	for _, markdown := range req.markdowns {
		files[markdown.Name()] = markdown
	}
	if req.header != nil {
		files["header.html"] = req.header
	}
	if req.footer != nil {
		files["footer.html"] = req.footer
	}
	for _, asset := range req.assets {
		files[asset.Name()] = asset
	}

	return files
}

// Assets sets assets form files.
func (req *MarkdownRequest) Assets(assets ...document.Document) {
	req.assets = assets
}

// Compile-time checks to ensure type implements desired interfaces.
var (
	_ = multipartRequest(new(MarkdownRequest))
)
