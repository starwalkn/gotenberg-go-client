package gotenberg

const (
	endpointURLConvert    = "/forms/chromium/convert/url"
	endpointURLScreenshot = "/forms/chromium/screenshot/url"
)

// URLRequest facilitates remote URL conversion with the Gotenberg API.
type URLRequest struct {
	*chromiumRequest
}

func NewURLRequest(url string) *URLRequest {
	req := &URLRequest{newChromiumRequest()}
	req.fields[fieldURL] = url

	return req
}

func (req *URLRequest) endpoint() string {
	return endpointURLConvert
}

func (req *URLRequest) screenshotEndpoint() string {
	return endpointURLScreenshot
}

func (req *URLRequest) formDocuments() map[string]Document {
	files := make(map[string]Document)

	if req.header != nil {
		files["header.html"] = req.header
	}
	if req.footer != nil {
		files["footer.html"] = req.footer
	}

	return files
}

func (req *URLRequest) Metadata(jsonData []byte) {
	req.fields[fieldMetadata] = string(jsonData)
}

// Compile-time checks to ensure type implements desired interfaces.
var (
	_ = MainRequester(new(URLRequest))
)
