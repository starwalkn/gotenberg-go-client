package gotenberg

const (
	endpointMarkdownConvert    = "/forms/chromium/convert/markdown"
	endpointMarkdownScreenshot = "/forms/chromium/screenshot/markdown"
)

// MarkdownRequest facilitates Markdown conversion with the Gotenberg API.
type MarkdownRequest struct {
	index     Document
	markdowns []Document
	assets    []Document

	*chromiumRequest
}

func NewMarkdownRequest(index Document, markdowns ...Document) *MarkdownRequest {
	return &MarkdownRequest{index, markdowns, []Document{}, newChromiumRequest()}
}

func (req *MarkdownRequest) endpoint() string {
	return endpointMarkdownConvert
}

func (req *MarkdownRequest) screenshotEndpoint() string {
	return endpointMarkdownScreenshot
}

func (req *MarkdownRequest) formDocuments() map[string]Document {
	files := make(map[string]Document)
	files["index.html"] = req.index
	for _, markdown := range req.markdowns {
		files[markdown.Filename()] = markdown
	}
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
func (req *MarkdownRequest) Assets(assets ...Document) {
	req.assets = assets
}

func (req *MarkdownRequest) Metadata(jsonData []byte) {
	req.fields[fieldMetadata] = string(jsonData)
}

// Compile-time checks to ensure type implements desired interfaces.
var (
	_ = MainRequester(new(MarkdownRequest))
)
