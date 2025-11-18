package gotenberg

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/starwalkn/gotenberg-go-client/v8/document"
)

const (
	endpointURLConvert    = "/forms/chromium/convert/url"
	endpointURLScreenshot = "/forms/chromium/screenshot/url"
)

// URLRequest facilitates remote URL conversion with the Gotenberg API.
type URLRequest struct {
	embeds []document.Document

	*chromiumRequest
	client *Client
}

func (r *URLRequest) endpoint() string {
	return endpointURLConvert
}

func (r *URLRequest) screenshotEndpoint() string {
	return endpointURLScreenshot
}

func (r *URLRequest) formDocuments() map[string]document.Document {
	files := make(map[string]document.Document)

	if r.header != nil {
		files["header.html"] = r.header
	}
	if r.footer != nil {
		files["footer.html"] = r.footer
	}

	return files
}

func (r *URLRequest) formEmbeds() map[string]document.Document {
	embeds := make(map[string]document.Document)

	for _, embed := range r.embeds {
		embeds[embed.Filename()] = embed
	}

	return embeds
}

func (r *URLRequest) Embeds(docs ...document.Document) *URLRequest {
	r.embeds = append(r.embeds, docs...)
	return r
}

// WaitDelay sets the duration to wait when loading an HTML document before converting it to PDF.
func (r *URLRequest) WaitDelay(delay time.Duration) *URLRequest {
	r.fields[fieldChromiumWaitDelay] = delay.String()
	return r
}

// WaitForExpression sets the JavaScript expression to wait before converting an HTML document into PDF until it returns true.
func (r *URLRequest) WaitForExpression(expression string) *URLRequest {
	r.fields[fieldChromiumWaitForExpression] = expression
	return r
}

// EmulatePrintMediaType forces Chromium to emulate the media type "print".
func (r *URLRequest) EmulatePrintMediaType() *URLRequest {
	r.fields[fieldChromiumEmulatedMediaType] = "print"
	return r
}

// EmulateScreenMediaType forces Chromium to emulate the media type "screen".
func (r *URLRequest) EmulateScreenMediaType() *URLRequest {
	r.fields[fieldChromiumEmulatedMediaType] = "screen"
	return r
}

// Cookies to store in the Chromium cookie jar.
func (r *URLRequest) Cookies(cookies []Cookie) *URLRequest {
	r.fields[fieldChromiumCookies] = mustJSON(cookies)
	return r
}

// UserAgent overrides the default User-Agent HTTP header.
func (r *URLRequest) UserAgent(ua string) *URLRequest {
	r.fields[fieldChromiumUserAgent] = ua
	return r
}

// ExtraHTTPHeaders sets extra HTTP headers that Chromium will send when loading the HTML document.
func (r *URLRequest) ExtraHTTPHeaders(headers map[string]string) *URLRequest {
	r.fields[fieldChromiumExtraHTTPHeaders] = mustJSON(headers)
	return r
}

// FailOnHTTPStatusCodes forces Gotenberg to return a 409 Conflict if the main page returns unacceptable HTTP status codes.
func (r *URLRequest) FailOnHTTPStatusCodes(statusCodes []int) *URLRequest {
	r.fields[fieldChromiumFailOnHTTPStatusCodes] = mustJSON(statusCodes)
	return r
}

// FailOnResourceHTTPStatusCodes forces Gotenberg to return a 409 Conflict if any resource returns unacceptable HTTP status codes.
func (r *URLRequest) FailOnResourceHTTPStatusCodes(statusCodes []int) *URLRequest {
	r.fields[fieldChromiumFailOnResourceHTTPStatusCodes] = mustJSON(statusCodes)
	return r
}

// FailOnConsoleExceptions forces Gotenberg to return a 409 Conflict if there are exceptions in the Chromium console.
func (r *URLRequest) FailOnConsoleExceptions() *URLRequest {
	r.fields[fieldChromiumFailOnConsoleExceptions] = strconv.FormatBool(true)
	return r
}

// FailOnResourceLoadingFailed forces Gotenberg to return a 409 Conflict if Chromium fails to load at least one resource.
func (r *URLRequest) FailOnResourceLoadingFailed() *URLRequest {
	r.fields[fieldChromiumFailOnResourceLoadingFailed] = strconv.FormatBool(true)
	return r
}

// SkipNetworkIdleEvent specifies whether Chromium should wait for network idle. Enabled by default in Gotenberg >= 8.11.0.
func (r *URLRequest) SkipNetworkIdleEvent(val bool) *URLRequest {
	r.fields[fieldChromiumSkipNetworkIdleEvent] = strconv.FormatBool(val)
	return r
}

// SinglePage defines whether to print the entire content in one single page.
func (r *URLRequest) SinglePage() *URLRequest {
	r.fields[fieldChromiumSinglePage] = strconv.FormatBool(true)
	return r
}

// PaperSize sets paperWidth and paperHeight form fields.
func (r *URLRequest) PaperSize(size PaperDimensions) *URLRequest {
	if size.Unit == "" {
		size.Unit = IN
	}

	r.fields[fieldChromiumPaperWidth] = fmt.Sprintf("%f%s", size.Width, size.Unit)
	r.fields[fieldChromiumPaperHeight] = fmt.Sprintf("%f%s", size.Height, size.Unit)

	return r
}

// Margins sets page margins.
func (r *URLRequest) Margins(margins PageMargins) *URLRequest {
	if margins.Unit == "" {
		margins.Unit = IN
	}

	r.fields[fieldChromiumMarginTop] = fmt.Sprintf("%f%s", margins.Top, margins.Unit)
	r.fields[fieldChromiumMarginBottom] = fmt.Sprintf("%f%s", margins.Bottom, margins.Unit)
	r.fields[fieldChromiumMarginLeft] = fmt.Sprintf("%f%s", margins.Left, margins.Unit)
	r.fields[fieldChromiumMarginRight] = fmt.Sprintf("%f%s", margins.Right, margins.Unit)

	return r
}

// PreferCSSPageSize forces page size as defined by CSS.
func (r *URLRequest) PreferCSSPageSize() *URLRequest {
	r.fields[fieldChromiumPreferCSSPageSize] = strconv.FormatBool(true)
	return r
}

// PrintBackground prints the background graphics.
func (r *URLRequest) PrintBackground() *URLRequest {
	r.fields[fieldChromiumPrintBackground] = strconv.FormatBool(true)
	return r
}

// OmitBackground hides default white background and allows transparency in PDFs.
func (r *URLRequest) OmitBackground() *URLRequest {
	r.fields[fieldChromiumOmitBackground] = strconv.FormatBool(true)
	return r
}

// Landscape sets paper orientation to landscape.
func (r *URLRequest) Landscape() *URLRequest {
	r.fields[fieldChromiumLandscapeChrome] = strconv.FormatBool(true)
	return r
}

// Scale overrides the default page scale.
func (r *URLRequest) Scale(factor float64) *URLRequest {
	r.fields[fieldChromiumScale] = fmt.Sprintf("%f", factor)
	return r
}

// NativePageRanges sets page ranges to print, e.g., "1-5, 8, 11-13".
func (r *URLRequest) NativePageRanges(ranges string) *URLRequest {
	r.fields[fieldChromiumNativePageRanges] = ranges
	return r
}

// GenerateDocumentOutline embeds the document outline into the PDF.
func (r *URLRequest) GenerateDocumentOutline() *URLRequest {
	r.fields[fieldChromiumGenerateDocumentOutline] = strconv.FormatBool(true)
	return r
}

// Header sets a header document.
func (r *URLRequest) Header(header document.Document) *URLRequest {
	r.header = header
	return r
}

// Footer sets a footer document.
func (r *URLRequest) Footer(footer document.Document) *URLRequest {
	r.footer = footer
	return r
}

// PdfA sets PDF/A format.
func (r *URLRequest) PdfA(pdfa PdfAFormat) *URLRequest {
	r.fields[fieldOfficePdfA] = string(pdfa)
	return r
}

// PdfUA enables PDF/UA accessibility.
func (r *URLRequest) PdfUA() *URLRequest {
	r.fields[fieldOfficePdfUa] = strconv.FormatBool(true)
	return r
}

// Metadata sets the metadata to write.
func (r *URLRequest) Metadata(jsonData []byte) *URLRequest {
	r.fields[fieldMetadata] = string(jsonData)
	return r
}

// SplitIntervals splits PDF by interval.
func (r *URLRequest) SplitIntervals(span int) *URLRequest {
	r.fields[fieldSplitMode] = splitModeIntervals
	r.fields[fieldSplitSpan] = strconv.Itoa(span)
	return r
}

// SplitPages splits PDF by pages.
func (r *URLRequest) SplitPages(span string, unify bool) *URLRequest {
	r.fields[fieldSplitMode] = splitModePages
	r.fields[fieldSplitSpan] = span
	r.fields[fieldSplitUnify] = strconv.FormatBool(unify)
	return r
}

// GenerateTaggedPDF enables tagged PDF generation.
func (r *URLRequest) GenerateTaggedPDF(val bool) *URLRequest {
	r.fields[fieldChromiumGenerateTaggedPDF] = strconv.FormatBool(val)
	return r
}

// Encrypt sets user and owner password.
func (r *URLRequest) Encrypt(userPassword, ownerPassword string) *URLRequest {
	r.fields[fieldUserPassword] = userPassword
	r.fields[fieldOwnerPassword] = ownerPassword
	return r
}

// ScreenshotWidth sets device screen width.
func (r *URLRequest) ScreenshotWidth(width int) *URLRequest {
	r.fields[fieldScreenshotWidth] = strconv.Itoa(width)
	return r
}

// ScreenshotHeight sets device screen height.
func (r *URLRequest) ScreenshotHeight(height int) *URLRequest {
	r.fields[fieldScreenshotHeight] = strconv.Itoa(height)
	return r
}

// ScreenshotClip enables clipping to device dimensions.
func (r *URLRequest) ScreenshotClip() *URLRequest {
	r.fields[fieldScreenshotClip] = strconv.FormatBool(true)
	return r
}

// ScreenshotQuality sets JPEG quality.
func (r *URLRequest) ScreenshotQuality(quality int) *URLRequest {
	r.fields[fieldScreenshotQuality] = strconv.Itoa(quality)
	return r
}

// ScreenshotOptimizeForSpeed optimizes image encoding for speed.
func (r *URLRequest) ScreenshotOptimizeForSpeed() *URLRequest {
	r.fields[fieldScreenshotOptimizeForSpeed] = strconv.FormatBool(true)
	return r
}

// Format sets screenshot format (PNG, JPEG, WEBP).
func (r *URLRequest) Format(format ImageFormat) *URLRequest {
	r.fields[fieldScreenshotFormat] = string(format)
	return r
}

func (r *URLRequest) OutputFilename(filename string) *URLRequest {
	r.chromiumRequest.OutputFilename(filename)
	return r
}

func (r *URLRequest) Trace(trace string) *URLRequest {
	r.chromiumRequest.Trace(trace)
	return r
}

func (r *URLRequest) BasicAuth(username, password string) *URLRequest {
	r.chromiumRequest.BasicAuth(username, password)
	return r
}

func (r *URLRequest) UseWebhook(hookURL, errorURL string) *URLRequest {
	r.chromiumRequest.UseWebhook(hookURL, errorURL)
	return r
}

func (r *URLRequest) SetWebhookMethod(method string) *URLRequest {
	r.chromiumRequest.SetWebhookMethod(method)
	return r
}

func (r *URLRequest) SetWebhookErrorMethod(method string) *URLRequest {
	r.chromiumRequest.SetWebhookErrorMethod(method)
	return r
}

func (r *URLRequest) SetWebhookExtraHeaders(headers map[string]string) *URLRequest {
	r.chromiumRequest.SetWebhookExtraHeaders(headers)
	return r
}

func (r *URLRequest) DownloadFrom(downloads map[string]map[string]string, embedded bool) *URLRequest {
	r.baseRequest.DownloadFrom(downloads, embedded)
	return r
}

func (r *URLRequest) Store(ctx context.Context, path string) error {
	return r.client.store(ctx, r, path)
}

func (r *URLRequest) Send(ctx context.Context) (*http.Response, error) {
	return r.client.send(ctx, r)
}

func (r *URLRequest) Screenshot(ctx context.Context) (*http.Response, error) {
	return r.client.screenshot(ctx, r)
}

func (r *URLRequest) StoreScreenshot(ctx context.Context, path string) error {
	return r.client.storeScreenshot(ctx, r, path)
}

// Compile-time checks to ensure type implements desired interfaces.
var (
	_ = MultipartRequest(new(URLRequest))
	_ = ScreenshotRequest(new(URLRequest))
)
