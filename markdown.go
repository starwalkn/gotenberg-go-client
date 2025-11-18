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
	endpointMarkdownConvert    = "/forms/chromium/convert/markdown"
	endpointMarkdownScreenshot = "/forms/chromium/screenshot/markdown"
)

// MarkdownRequest facilitates Markdown conversion with the Gotenberg API.
type MarkdownRequest struct {
	index     document.Document
	markdowns []document.Document
	assets    []document.Document
	embeds    []document.Document

	*chromiumRequest
	client *Client
}

func (r *MarkdownRequest) endpoint() string {
	return endpointMarkdownConvert
}

func (r *MarkdownRequest) screenshotEndpoint() string {
	return endpointMarkdownScreenshot
}

func (r *MarkdownRequest) formDocuments() map[string]document.Document {
	files := make(map[string]document.Document)
	files["index.html"] = r.index
	for _, markdown := range r.markdowns {
		files[markdown.Filename()] = markdown
	}
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

func (r *MarkdownRequest) formEmbeds() map[string]document.Document {
	embeds := make(map[string]document.Document)

	for _, embed := range r.embeds {
		embeds[embed.Filename()] = embed
	}

	return embeds
}

func (r *MarkdownRequest) Embeds(docs ...document.Document) *MarkdownRequest {
	r.embeds = append(r.embeds, docs...)
	return r
}

// Assets sets assets form files.
func (r *MarkdownRequest) Assets(assets ...document.Document) *MarkdownRequest {
	r.assets = assets
	return r
}

// WaitDelay sets the duration to wait when loading an HTML document before converting it to PDF.
func (r *MarkdownRequest) WaitDelay(delay time.Duration) *MarkdownRequest {
	r.fields[fieldChromiumWaitDelay] = delay.String()
	return r
}

// WaitForExpression sets the JavaScript expression to wait before converting an HTML document into PDF until it returns true.
func (r *MarkdownRequest) WaitForExpression(expression string) *MarkdownRequest {
	r.fields[fieldChromiumWaitForExpression] = expression
	return r
}

// EmulatePrintMediaType forces Chromium to emulate the media type "print".
func (r *MarkdownRequest) EmulatePrintMediaType() *MarkdownRequest {
	r.fields[fieldChromiumEmulatedMediaType] = "print"
	return r
}

// EmulateScreenMediaType forces Chromium to emulate the media type "screen".
func (r *MarkdownRequest) EmulateScreenMediaType() *MarkdownRequest {
	r.fields[fieldChromiumEmulatedMediaType] = "screen"
	return r
}

// Cookies to store in the Chromium cookie jar.
func (r *MarkdownRequest) Cookies(cookies []Cookie) *MarkdownRequest {
	r.fields[fieldChromiumCookies] = mustJSON(cookies)
	return r
}

// UserAgent overrides the default User-Agent HTTP header.
func (r *MarkdownRequest) UserAgent(ua string) *MarkdownRequest {
	r.fields[fieldChromiumUserAgent] = ua
	return r
}

// ExtraHTTPHeaders sets extra HTTP headers that Chromium will send when loading the HTML document.
func (r *MarkdownRequest) ExtraHTTPHeaders(headers map[string]string) *MarkdownRequest {
	r.fields[fieldChromiumExtraHTTPHeaders] = mustJSON(headers)
	return r
}

// FailOnHTTPStatusCodes forces Gotenberg to return a 409 Conflict if the main page returns unacceptable HTTP status codes.
func (r *MarkdownRequest) FailOnHTTPStatusCodes(statusCodes []int) *MarkdownRequest {
	r.fields[fieldChromiumFailOnHTTPStatusCodes] = mustJSON(statusCodes)
	return r
}

// FailOnResourceHTTPStatusCodes forces Gotenberg to return a 409 Conflict if any resource returns unacceptable HTTP status codes.
func (r *MarkdownRequest) FailOnResourceHTTPStatusCodes(statusCodes []int) *MarkdownRequest {
	r.fields[fieldChromiumFailOnResourceHTTPStatusCodes] = mustJSON(statusCodes)
	return r
}

// FailOnConsoleExceptions forces Gotenberg to return a 409 Conflict if there are exceptions in the Chromium console.
func (r *MarkdownRequest) FailOnConsoleExceptions() *MarkdownRequest {
	r.fields[fieldChromiumFailOnConsoleExceptions] = strconv.FormatBool(true)
	return r
}

// FailOnResourceLoadingFailed forces Gotenberg to return a 409 Conflict if Chromium fails to load at least one resource.
func (r *MarkdownRequest) FailOnResourceLoadingFailed() *MarkdownRequest {
	r.fields[fieldChromiumFailOnResourceLoadingFailed] = strconv.FormatBool(true)
	return r
}

// SkipNetworkIdleEvent specifies whether Chromium should wait for network idle. Enabled by default in Gotenberg >= 8.11.0.
func (r *MarkdownRequest) SkipNetworkIdleEvent(val bool) *MarkdownRequest {
	r.fields[fieldChromiumSkipNetworkIdleEvent] = strconv.FormatBool(val)
	return r
}

// SinglePage defines whether to print the entire content in one single page.
func (r *MarkdownRequest) SinglePage() *MarkdownRequest {
	r.fields[fieldChromiumSinglePage] = strconv.FormatBool(true)
	return r
}

// PaperSize sets paperWidth and paperHeight form fields.
func (r *MarkdownRequest) PaperSize(size PaperDimensions) *MarkdownRequest {
	if size.Unit == "" {
		size.Unit = IN
	}

	r.fields[fieldChromiumPaperWidth] = fmt.Sprintf("%f%s", size.Width, size.Unit)
	r.fields[fieldChromiumPaperHeight] = fmt.Sprintf("%f%s", size.Height, size.Unit)

	return r
}

// Margins sets page margins.
func (r *MarkdownRequest) Margins(margins PageMargins) *MarkdownRequest {
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
func (r *MarkdownRequest) PreferCSSPageSize() *MarkdownRequest {
	r.fields[fieldChromiumPreferCSSPageSize] = strconv.FormatBool(true)
	return r
}

// PrintBackground prints the background graphics.
func (r *MarkdownRequest) PrintBackground() *MarkdownRequest {
	r.fields[fieldChromiumPrintBackground] = strconv.FormatBool(true)
	return r
}

// OmitBackground hides default white background and allows transparency in PDFs.
func (r *MarkdownRequest) OmitBackground() *MarkdownRequest {
	r.fields[fieldChromiumOmitBackground] = strconv.FormatBool(true)
	return r
}

// Landscape sets paper orientation to landscape.
func (r *MarkdownRequest) Landscape() *MarkdownRequest {
	r.fields[fieldChromiumLandscapeChrome] = strconv.FormatBool(true)
	return r
}

// Scale overrides the default page scale.
func (r *MarkdownRequest) Scale(factor float64) *MarkdownRequest {
	r.fields[fieldChromiumScale] = fmt.Sprintf("%f", factor)
	return r
}

// NativePageRanges sets page ranges to print, e.g., "1-5, 8, 11-13".
func (r *MarkdownRequest) NativePageRanges(ranges string) *MarkdownRequest {
	r.fields[fieldChromiumNativePageRanges] = ranges
	return r
}

// GenerateDocumentOutline embeds the document outline into the PDF.
func (r *MarkdownRequest) GenerateDocumentOutline() *MarkdownRequest {
	r.fields[fieldChromiumGenerateDocumentOutline] = strconv.FormatBool(true)
	return r
}

// Header sets a header document.
func (r *MarkdownRequest) Header(header document.Document) *MarkdownRequest {
	r.header = header
	return r
}

// Footer sets a footer document.
func (r *MarkdownRequest) Footer(footer document.Document) *MarkdownRequest {
	r.footer = footer
	return r
}

// PdfA sets PDF/A format.
func (r *MarkdownRequest) PdfA(pdfa PdfAFormat) *MarkdownRequest {
	r.fields[fieldOfficePdfA] = string(pdfa)
	return r
}

// PdfUA enables PDF/UA accessibility.
func (r *MarkdownRequest) PdfUA() *MarkdownRequest {
	r.fields[fieldOfficePdfUa] = strconv.FormatBool(true)
	return r
}

// Metadata sets the metadata to write.
func (r *MarkdownRequest) Metadata(jsonData []byte) *MarkdownRequest {
	r.fields[fieldMetadata] = string(jsonData)
	return r
}

// SplitIntervals splits PDF by interval.
func (r *MarkdownRequest) SplitIntervals(span int) *MarkdownRequest {
	r.fields[fieldSplitMode] = splitModeIntervals
	r.fields[fieldSplitSpan] = strconv.Itoa(span)
	return r
}

// SplitPages splits PDF by pages.
func (r *MarkdownRequest) SplitPages(span string, unify bool) *MarkdownRequest {
	r.fields[fieldSplitMode] = splitModePages
	r.fields[fieldSplitSpan] = span
	r.fields[fieldSplitUnify] = strconv.FormatBool(unify)
	return r
}

// GenerateTaggedPDF enables tagged PDF generation.
func (r *MarkdownRequest) GenerateTaggedPDF(val bool) *MarkdownRequest {
	r.fields[fieldChromiumGenerateTaggedPDF] = strconv.FormatBool(val)
	return r
}

// Encrypt sets user and owner password.
func (r *MarkdownRequest) Encrypt(userPassword, ownerPassword string) *MarkdownRequest {
	r.fields[fieldUserPassword] = userPassword
	r.fields[fieldOwnerPassword] = ownerPassword
	return r
}

// ScreenshotWidth sets device screen width.
func (r *MarkdownRequest) ScreenshotWidth(width int) *MarkdownRequest {
	r.fields[fieldScreenshotWidth] = strconv.Itoa(width)
	return r
}

// ScreenshotHeight sets device screen height.
func (r *MarkdownRequest) ScreenshotHeight(height int) *MarkdownRequest {
	r.fields[fieldScreenshotHeight] = strconv.Itoa(height)
	return r
}

// ScreenshotClip enables clipping to device dimensions.
func (r *MarkdownRequest) ScreenshotClip() *MarkdownRequest {
	r.fields[fieldScreenshotClip] = strconv.FormatBool(true)
	return r
}

// ScreenshotQuality sets JPEG quality.
func (r *MarkdownRequest) ScreenshotQuality(quality int) *MarkdownRequest {
	r.fields[fieldScreenshotQuality] = strconv.Itoa(quality)
	return r
}

// ScreenshotOptimizeForSpeed optimizes image encoding for speed.
func (r *MarkdownRequest) ScreenshotOptimizeForSpeed() *MarkdownRequest {
	r.fields[fieldScreenshotOptimizeForSpeed] = strconv.FormatBool(true)
	return r
}

// Format sets screenshot format (PNG, JPEG, WEBP).
func (r *MarkdownRequest) Format(format ImageFormat) *MarkdownRequest {
	r.fields[fieldScreenshotFormat] = string(format)
	return r
}

func (r *MarkdownRequest) OutputFilename(filename string) *MarkdownRequest {
	r.baseRequest.OutputFilename(filename)
	return r
}

func (r *MarkdownRequest) Trace(trace string) *MarkdownRequest {
	r.baseRequest.Trace(trace)
	return r
}

func (r *MarkdownRequest) BasicAuth(username, password string) *MarkdownRequest {
	r.baseRequest.BasicAuth(username, password)
	return r
}

func (r *MarkdownRequest) UseWebhook(hookURL, errorURL string) *MarkdownRequest {
	r.baseRequest.UseWebhook(hookURL, errorURL)
	return r
}

func (r *MarkdownRequest) SetWebhookMethod(method string) *MarkdownRequest {
	r.baseRequest.SetWebhookMethod(method)
	return r
}

func (r *MarkdownRequest) SetWebhookErrorMethod(method string) *MarkdownRequest {
	r.baseRequest.SetWebhookErrorMethod(method)
	return r
}

func (r *MarkdownRequest) SetWebhookExtraHeaders(headers map[string]string) *MarkdownRequest {
	r.baseRequest.SetWebhookExtraHeaders(headers)
	return r
}

func (r *MarkdownRequest) DownloadFrom(downloads map[string]map[string]string, embedded bool) *MarkdownRequest {
	r.baseRequest.DownloadFrom(downloads, embedded)
	return r
}

func (r *MarkdownRequest) Store(ctx context.Context, path string) error {
	return r.client.store(ctx, r, path)
}

func (r *MarkdownRequest) Send(ctx context.Context) (*http.Response, error) {
	return r.client.send(ctx, r)
}

func (r *MarkdownRequest) Screenshot(ctx context.Context) (*http.Response, error) {
	return r.client.screenshot(ctx, r)
}

func (r *MarkdownRequest) StoreScreenshot(ctx context.Context, path string) error {
	return r.client.storeScreenshot(ctx, r, path)
}

// Compile-time checks to ensure type implements desired interfaces.
var (
	_ = MultipartRequest(new(MarkdownRequest))
	_ = ScreenshotRequest(new(MarkdownRequest))
)
