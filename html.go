package gotenberg

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/starwalkn/gotenberg-go-client/v9/document"
)

const (
	endpointHTMLConvert    = "/forms/chromium/convert/html"
	endpointHTMLScreenshot = "/forms/chromium/screenshot/html"
)

// HTMLRequest facilitates HTML conversion with the Gotenberg API.
type HTMLRequest struct {
	index  document.Document
	assets []document.Document
	embeds []document.Document

	*chromiumRequest
	client *Client
}

func (r *HTMLRequest) endpoint() string {
	return endpointHTMLConvert
}

func (r *HTMLRequest) screenshotEndpoint() string {
	return endpointHTMLScreenshot
}

func (r *HTMLRequest) formDocuments() map[string]document.Document {
	files := make(map[string]document.Document)
	files["index.html"] = r.index

	if r.header != nil {
		files["header.html"] = r.header
	}
	if r.footer != nil {
		files["footer.html"] = r.footer
	}

	for _, asset := range r.assets {
		files[asset.Filename()] = asset
	}

	return files
}

func (r *HTMLRequest) formEmbeds() map[string]document.Document {
	embeds := make(map[string]document.Document)

	for _, embed := range r.embeds {
		embeds[embed.Filename()] = embed
	}

	return embeds
}

func (r *HTMLRequest) Embeds(docs ...document.Document) *HTMLRequest {
	r.embeds = docs
	return r
}

// Assets sets assets form files.
func (r *HTMLRequest) Assets(assets ...document.Document) *HTMLRequest {
	r.assets = assets
	return r
}

// WaitDelay sets the duration to wait when loading an HTML document before converting it to PDF.
func (r *HTMLRequest) WaitDelay(delay time.Duration) *HTMLRequest {
	r.fields[fieldChromiumWaitDelay] = delay.String()
	return r
}

// WaitForExpression sets the JavaScript expression to wait before converting an HTML document into PDF until it returns true.
func (r *HTMLRequest) WaitForExpression(expression string) *HTMLRequest {
	r.fields[fieldChromiumWaitForExpression] = expression
	return r
}

// EmulatePrintMediaType forces Chromium to emulate the media type "print".
func (r *HTMLRequest) EmulatePrintMediaType() *HTMLRequest {
	r.fields[fieldChromiumEmulatedMediaType] = "print"
	return r
}

// EmulateScreenMediaType forces Chromium to emulate the media type "screen".
func (r *HTMLRequest) EmulateScreenMediaType() *HTMLRequest {
	r.fields[fieldChromiumEmulatedMediaType] = "screen"
	return r
}

// Cookies to store in the Chromium cookie jar.
func (r *HTMLRequest) Cookies(cookies []Cookie) *HTMLRequest {
	r.fields[fieldChromiumCookies] = mustJSON(cookies)
	return r
}

// UserAgent overrides the default User-Agent HTTP header.
func (r *HTMLRequest) UserAgent(ua string) *HTMLRequest {
	r.fields[fieldChromiumUserAgent] = ua
	return r
}

// ExtraHTTPHeaders sets extra HTTP headers that Chromium will send when loading the HTML document.
func (r *HTMLRequest) ExtraHTTPHeaders(headers map[string]string) *HTMLRequest {
	r.fields[fieldChromiumExtraHTTPHeaders] = mustJSON(headers)
	return r
}

// FailOnHTTPStatusCodes forces Gotenberg to return a 409 Conflict if the main page returns unacceptable HTTP status codes.
func (r *HTMLRequest) FailOnHTTPStatusCodes(statusCodes []int) *HTMLRequest {
	r.fields[fieldChromiumFailOnHTTPStatusCodes] = mustJSON(statusCodes)
	return r
}

// FailOnResourceHTTPStatusCodes forces Gotenberg to return a 409 Conflict if any resource returns unacceptable HTTP status codes.
func (r *HTMLRequest) FailOnResourceHTTPStatusCodes(statusCodes []int) *HTMLRequest {
	r.fields[fieldChromiumFailOnResourceHTTPStatusCodes] = mustJSON(statusCodes)
	return r
}

// FailOnConsoleExceptions forces Gotenberg to return a 409 Conflict if there are exceptions in the Chromium console.
func (r *HTMLRequest) FailOnConsoleExceptions() *HTMLRequest {
	r.fields[fieldChromiumFailOnConsoleExceptions] = strconv.FormatBool(true)
	return r
}

// FailOnResourceLoadingFailed forces Gotenberg to return a 409 Conflict if Chromium fails to load at least one resource.
func (r *HTMLRequest) FailOnResourceLoadingFailed() *HTMLRequest {
	r.fields[fieldChromiumFailOnResourceLoadingFailed] = strconv.FormatBool(true)
	return r
}

// SkipNetworkIdleEvent specifies whether Chromium should wait for network idle. Enabled by default in Gotenberg >= 8.11.0.
func (r *HTMLRequest) SkipNetworkIdleEvent(val bool) *HTMLRequest {
	r.fields[fieldChromiumSkipNetworkIdleEvent] = strconv.FormatBool(val)
	return r
}

// SinglePage defines whether to print the entire content in one single page.
func (r *HTMLRequest) SinglePage() *HTMLRequest {
	r.fields[fieldChromiumSinglePage] = strconv.FormatBool(true)
	return r
}

// PaperSize sets paperWidth and paperHeight form fields.
func (r *HTMLRequest) PaperSize(size PaperDimensions) *HTMLRequest {
	if size.Unit == "" {
		size.Unit = IN
	}

	r.fields[fieldChromiumPaperWidth] = fmt.Sprintf("%f%s", size.Width, size.Unit)
	r.fields[fieldChromiumPaperHeight] = fmt.Sprintf("%f%s", size.Height, size.Unit)

	return r
}

// Margins sets page margins.
func (r *HTMLRequest) Margins(margins PageMargins) *HTMLRequest {
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
func (r *HTMLRequest) PreferCSSPageSize() *HTMLRequest {
	r.fields[fieldChromiumPreferCSSPageSize] = strconv.FormatBool(true)
	return r
}

// PrintBackground prints the background graphics.
func (r *HTMLRequest) PrintBackground() *HTMLRequest {
	r.fields[fieldChromiumPrintBackground] = strconv.FormatBool(true)
	return r
}

// OmitBackground hides default white background and allows transparency in PDFs.
func (r *HTMLRequest) OmitBackground() *HTMLRequest {
	r.fields[fieldChromiumOmitBackground] = strconv.FormatBool(true)
	return r
}

// Landscape sets paper orientation to landscape.
func (r *HTMLRequest) Landscape() *HTMLRequest {
	r.fields[fieldChromiumLandscapeChrome] = strconv.FormatBool(true)
	return r
}

// Scale overrides the default page scale.
func (r *HTMLRequest) Scale(factor float64) *HTMLRequest {
	r.fields[fieldChromiumScale] = fmt.Sprintf("%f", factor)
	return r
}

// NativePageRanges sets page ranges to print, e.g., "1-5, 8, 11-13".
func (r *HTMLRequest) NativePageRanges(ranges string) *HTMLRequest {
	r.fields[fieldChromiumNativePageRanges] = ranges
	return r
}

// GenerateDocumentOutline embeds the document outline into the PDF.
func (r *HTMLRequest) GenerateDocumentOutline() *HTMLRequest {
	r.fields[fieldChromiumGenerateDocumentOutline] = strconv.FormatBool(true)
	return r
}

// Header sets a header document.
func (r *HTMLRequest) Header(header document.Document) *HTMLRequest {
	r.header = header
	return r
}

// Footer sets a footer document.
func (r *HTMLRequest) Footer(footer document.Document) *HTMLRequest {
	r.footer = footer
	return r
}

// PdfA sets PDF/A format.
func (r *HTMLRequest) PdfA(pdfa PdfAFormat) *HTMLRequest {
	r.fields[fieldOfficePdfA] = string(pdfa)
	return r
}

// PdfUA enables PDF/UA accessibility.
func (r *HTMLRequest) PdfUA() *HTMLRequest {
	r.fields[fieldOfficePdfUa] = strconv.FormatBool(true)
	return r
}

// Metadata sets the metadata to write.
func (r *HTMLRequest) Metadata(jsonData []byte) *HTMLRequest {
	r.fields[fieldMetadata] = string(jsonData)
	return r
}

// SplitIntervals splits PDF by interval.
func (r *HTMLRequest) SplitIntervals(span int) *HTMLRequest {
	r.fields[fieldSplitMode] = splitModeIntervals
	r.fields[fieldSplitSpan] = strconv.Itoa(span)
	return r
}

// SplitPages splits PDF by pages.
func (r *HTMLRequest) SplitPages(span string, unify bool) *HTMLRequest {
	r.fields[fieldSplitMode] = splitModePages
	r.fields[fieldSplitSpan] = span
	r.fields[fieldSplitUnify] = strconv.FormatBool(unify)
	return r
}

// GenerateTaggedPDF enables tagged PDF generation.
func (r *HTMLRequest) GenerateTaggedPDF(val bool) *HTMLRequest {
	r.fields[fieldChromiumGenerateTaggedPDF] = strconv.FormatBool(val)
	return r
}

// Encrypt sets user and owner password.
func (r *HTMLRequest) Encrypt(userPassword, ownerPassword string) *HTMLRequest {
	r.fields[fieldUserPassword] = userPassword
	r.fields[fieldOwnerPassword] = ownerPassword
	return r
}

// ScreenshotWidth sets device screen width.
func (r *HTMLRequest) ScreenshotWidth(width int) *HTMLRequest {
	r.fields[fieldScreenshotWidth] = strconv.Itoa(width)
	return r
}

// ScreenshotHeight sets device screen height.
func (r *HTMLRequest) ScreenshotHeight(height int) *HTMLRequest {
	r.fields[fieldScreenshotHeight] = strconv.Itoa(height)
	return r
}

// ScreenshotClip enables clipping to device dimensions.
func (r *HTMLRequest) ScreenshotClip() *HTMLRequest {
	r.fields[fieldScreenshotClip] = strconv.FormatBool(true)
	return r
}

// ScreenshotQuality sets JPEG quality.
func (r *HTMLRequest) ScreenshotQuality(quality int) *HTMLRequest {
	r.fields[fieldScreenshotQuality] = strconv.Itoa(quality)
	return r
}

// ScreenshotOptimizeForSpeed optimizes image encoding for speed.
func (r *HTMLRequest) ScreenshotOptimizeForSpeed() *HTMLRequest {
	r.fields[fieldScreenshotOptimizeForSpeed] = strconv.FormatBool(true)
	return r
}

// Format sets screenshot format (PNG, JPEG, WEBP).
func (r *HTMLRequest) Format(format ImageFormat) *HTMLRequest {
	r.fields[fieldScreenshotFormat] = string(format)
	return r
}

func (r *HTMLRequest) OutputFilename(filename string) *HTMLRequest {
	r.baseRequest.OutputFilename(filename)
	return r
}

func (r *HTMLRequest) Trace(trace string) *HTMLRequest {
	r.baseRequest.Trace(trace)
	return r
}

func (r *HTMLRequest) BasicAuth(username, password string) *HTMLRequest {
	r.baseRequest.BasicAuth(username, password)
	return r
}

func (r *HTMLRequest) UseWebhook(hookURL, errorURL string) *HTMLRequest {
	r.baseRequest.UseWebhook(hookURL, errorURL)
	return r
}

func (r *HTMLRequest) SetWebhookMethod(method string) *HTMLRequest {
	r.baseRequest.SetWebhookMethod(method)
	return r
}

func (r *HTMLRequest) SetWebhookErrorMethod(method string) *HTMLRequest {
	r.baseRequest.SetWebhookErrorMethod(method)
	return r
}

func (r *HTMLRequest) SetWebhookExtraHeaders(headers map[string]string) *HTMLRequest {
	r.baseRequest.SetWebhookExtraHeaders(headers)
	return r
}

func (r *HTMLRequest) DownloadFrom(downloads map[string]map[string]string, embedded bool) *HTMLRequest {
	r.baseRequest.DownloadFrom(downloads, embedded)
	return r
}

func (r *HTMLRequest) Store(ctx context.Context, path string) error {
	return r.client.store(ctx, r, path)
}

func (r *HTMLRequest) Send(ctx context.Context) (*http.Response, error) {
	return r.client.send(ctx, r)
}

func (r *HTMLRequest) Screenshot(ctx context.Context) (*http.Response, error) {
	return r.client.screenshot(ctx, r)
}

func (r *HTMLRequest) StoreScreenshot(ctx context.Context, path string) error {
	return r.client.storeScreenshot(ctx, r, path)
}

// Compile-time checks to ensure type implements desired interfaces.
var (
	_ = MultipartRequest(new(HTMLRequest))
	_ = ScreenshotRequest(new(HTMLRequest))
)
