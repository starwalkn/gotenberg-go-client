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
	req := &URLRequest{chromiumRequest: newChromiumRequest()}
	req.fields[fieldURL] = url

	return req
}

func (r *URLRequest) endpoint() string {
	return endpointURLConvert
}

func (r *URLRequest) screenshotEndpoint() string {
	return endpointURLScreenshot
}

func (r *URLRequest) formDocuments() map[string]document.Document {
	docs := make(map[string]document.Document)

	if r.header != nil {
		docs["header.html"] = r.header
	}
	if r.footer != nil {
		docs["footer.html"] = r.footer
	}

	return docs
}

// Compile-time checks to ensure type implements desired interfaces.
var (
	_ = multipartRequest(new(URLRequest))
)
