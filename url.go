package gotenberg

import "github.com/starwalkn/gotenberg-go-client/v8/document"

const (
	endpointURLConvert    = "/forms/chromium/convert/url"
	endpointURLScreenshot = "/forms/chromium/screenshot/url"
)

// URLRequest facilitates remote URL conversion with the Gotenberg API.
type URLRequest struct {
	embeds []document.Document

	*chromiumRequest
}

func NewURLRequest(url string) *URLRequest {
	req := &URLRequest{
		embeds:          []document.Document{},
		chromiumRequest: newChromiumRequest(),
	}
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

func (req *URLRequest) formEmbeds() map[string]document.Document {
	embeds := make(map[string]document.Document)

	for _, embed := range req.embeds {
		embeds[embed.Filename()] = embed
	}

	return embeds
}

func (req *URLRequest) Embeds(docs ...document.Document) {
	req.embeds = append(req.embeds, docs...)
}

// Compile-time checks to ensure type implements desired interfaces.
var (
	_ = MultipartRequest(new(URLRequest))
)
