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

func (r *MarkdownRequest) endpoint() string {
	return endpointMarkdownConvert
}

func (r *MarkdownRequest) screenshotEndpoint() string {
	return endpointMarkdownScreenshot
}

func (r *MarkdownRequest) formDocuments() map[string]document.Document {
	docs := make(map[string]document.Document)
	docs["index.html"] = r.index
	for _, markdown := range r.markdowns {
		docs[markdown.Name()] = markdown
	}
	if r.header != nil {
		docs["header.html"] = r.header
	}
	if r.footer != nil {
		docs["footer.html"] = r.footer
	}
	for _, asset := range r.assets {
		docs[asset.Name()] = asset
	}

	return docs
}

// Assets sets assets form files.
func (r *MarkdownRequest) Assets(assets ...document.Document) {
	r.assets = assets
}

// Compile-time checks to ensure type implements desired interfaces.
var (
	_ = multipartRequest(new(MarkdownRequest))
)
