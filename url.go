package gotenberg

import "github.com/starwalkn/gotenberg-go-client/v8/document"

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

func (req *URLRequest) formDocuments() map[string]document.Document {
	files := make(map[string]document.Document)

	if req.header != nil {
		files["header.html"] = req.header
	}
	if req.footer != nil {
		files["footer.html"] = req.footer
	}

	return files
}

// Compile-time checks to ensure type implements desired interfaces.
var (
	_ = multipartRequest(new(URLRequest))
)
