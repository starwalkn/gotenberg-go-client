package gotenberg

const (
	endpointHTMLConvert    = "/forms/chromium/convert/html"
	endpointHTMLScreenshot = "/forms/chromium/screenshot/html"
)

// HTMLRequest facilitates HTML conversion with the Gotenberg API.
type HTMLRequest struct {
	index  Document
	assets []Document

	*chromiumRequest
}

func NewHTMLRequest(index Document) *HTMLRequest {
	return &HTMLRequest{index, []Document{}, newChromiumRequest()}
}

func (req *HTMLRequest) endpoint() string {
	return endpointHTMLConvert
}

func (req *HTMLRequest) screenshotEndpoint() string {
	return endpointHTMLScreenshot
}

func (req *HTMLRequest) formDocuments() map[string]Document {
	files := make(map[string]Document)
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
func (req *HTMLRequest) Assets(assets ...Document) {
	req.assets = assets
}

// Compile-time checks to ensure type implements desired interfaces.
var (
	_ = MainRequester(new(HTMLRequest))
)
