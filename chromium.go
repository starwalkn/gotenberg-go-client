package gotenberg

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/starwalkn/gotenberg-go-client/v8/document"
)

type chromiumRequest struct {
	header document.Document
	footer document.Document

	*baseRequest
}

func newChromiumRequest() *chromiumRequest {
	return &chromiumRequest{nil, nil, newBaseRequest()}
}

// WaitDelay sets the duration (i.e., "1s", "2ms", etc.) to wait when loading an
// HTML document before converting it to PDF.
func (r *chromiumRequest) WaitDelay(delay time.Duration) *chromiumRequest {
	r.fields[fieldChromiumWaitDelay] = delay.String()
	return r
}

// WaitForExpression sets the JavaScript expression to wait before converting an
// HTML document into PDF until it returns true.
func (r *chromiumRequest) WaitForExpression(expression string) *chromiumRequest {
	r.fields[fieldChromiumWaitForExpression] = expression
	return r
}

// EmulatePrintMediaType forces Chromium to emulate the media type "print".
func (r *chromiumRequest) EmulatePrintMediaType() *chromiumRequest {
	r.fields[fieldChromiumEmulatedMediaType] = "print"
	return r
}

// EmulateScreenMediaType forces Chromium to emulate the media type "screen".
func (r *chromiumRequest) EmulateScreenMediaType() *chromiumRequest {
	r.fields[fieldChromiumEmulatedMediaType] = "screen"
	return r
}

// Cookies to store in the Chromium cookie jar.
func (r *chromiumRequest) Cookies(cookies []Cookie) *chromiumRequest {
	r.fields[fieldChromiumCookies] = mustJSON(cookies)
	return r
}

// UserAgent overrides the default User-Agent HTTP header.
func (r *chromiumRequest) UserAgent(ua string) *chromiumRequest {
	r.fields[fieldChromiumUserAgent] = ua
	return r
}

// ExtraHTTPHeaders sets extra HTTP headers that Chromium will send when loading the HTML document.
func (r *chromiumRequest) ExtraHTTPHeaders(headers map[string]string) *chromiumRequest {
	r.fields[fieldChromiumExtraHTTPHeaders] = mustJSON(headers)
	return r
}

// FailOnHTTPStatusCodes forces Gotenberg to return a 409 Conflict response
// if the HTTP status code from the main page is not acceptable.
func (r *chromiumRequest) FailOnHTTPStatusCodes(statusCodes []int) *chromiumRequest {
	r.fields[fieldChromiumFailOnHTTPStatusCodes] = mustJSON(statusCodes)
	return r
}

// FailOnResourceHTTPStatusCodes forces Gotenberg to return a 409 Conflict response
// if the HTTP status code from at least one resource is not acceptable.
func (r *chromiumRequest) FailOnResourceHTTPStatusCodes(statusCodes []int) *chromiumRequest {
	r.fields[fieldChromiumFailOnResourceHTTPStatusCodes] = mustJSON(statusCodes)
	return r
}

// FailOnConsoleExceptions forces Gotenberg to return a 409 Conflict response
// if there are exceptions in the Chromium console.
func (r *chromiumRequest) FailOnConsoleExceptions() *chromiumRequest {
	r.fields[fieldChromiumFailOnConsoleExceptions] = strconv.FormatBool(true)
	return r
}

// FailOnResourceLoadingFailed forces Gotenberg to return a 409 Conflict if Chromium
// fails to load at least one resource.
func (r *chromiumRequest) FailOnResourceLoadingFailed() *chromiumRequest {
	r.fields[fieldChromiumFailOnResourceLoadingFailed] = strconv.FormatBool(true)
	return r
}

// SkipNetworkIdleEvent specifies whether Chromium have to wait or not for its network to be idle.
// Enabled by default in Gotenberg >= 8.11.0.
func (r *chromiumRequest) SkipNetworkIdleEvent(val bool) *chromiumRequest {
	r.fields[fieldChromiumSkipNetworkIdleEvent] = strconv.FormatBool(val)
	return r
}

// SinglePage defines whether to print the entire content in one single page.
func (r *chromiumRequest) SinglePage() *chromiumRequest {
	r.fields[fieldChromiumSinglePage] = strconv.FormatBool(true)
	return r
}

// PaperSize sets paperWidth and paperHeight form fields with the provided unit.
// If unit is empty, it defaults to inches. Default is Letter (8.5 x 11 inches).
func (r *chromiumRequest) PaperSize(size PaperDimensions) *chromiumRequest {
	if size.Unit == "" {
		size.Unit = IN
	}

	r.fields[fieldChromiumPaperWidth] = fmt.Sprintf("%f%s", size.Width, size.Unit)
	r.fields[fieldChromiumPaperHeight] = fmt.Sprintf("%f%s", size.Height, size.Unit)

	return r
}

// Margins sets marginTop, marginBottom, marginLeft and marginRight form fields. Default unit is inches.
func (r *chromiumRequest) Margins(margins PageMargins) *chromiumRequest {
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
func (r *chromiumRequest) PreferCSSPageSize() *chromiumRequest {
	r.fields[fieldChromiumPreferCSSPageSize] = strconv.FormatBool(true)
	return r
}

// PrintBackground prints the background graphics.
func (r *chromiumRequest) PrintBackground() *chromiumRequest {
	r.fields[fieldChromiumPrintBackground] = strconv.FormatBool(true)
	return r
}

// OmitBackground hides default white background and allows generating PDFs with transparency.
func (r *chromiumRequest) OmitBackground() *chromiumRequest {
	r.fields[fieldChromiumOmitBackground] = strconv.FormatBool(true)
	return r
}

// Landscape sets the paper orientation to landscape.
func (r *chromiumRequest) Landscape() *chromiumRequest {
	r.fields[fieldChromiumLandscapeChrome] = strconv.FormatBool(true)
	return r
}

// Scale overrides the default scale of the page rendering (i.e., 1.0).
func (r *chromiumRequest) Scale(factor float64) *chromiumRequest {
	r.fields[fieldChromiumScale] = fmt.Sprintf("%f", factor)
	return r
}

// NativePageRanges sets the page ranges to print, e.g., "1-5, 8, 11-13". Empty means all pages.
func (r *chromiumRequest) NativePageRanges(ranges string) *chromiumRequest {
	r.fields[fieldChromiumNativePageRanges] = ranges
	return r
}

// GenerateDocumentOutline embeds the document outline into the PDF.
func (r *chromiumRequest) GenerateDocumentOutline() *chromiumRequest {
	r.fields[fieldChromiumGenerateDocumentOutline] = strconv.FormatBool(true)
	return r
}

// Header adds a header to each page.
func (r *chromiumRequest) Header(header document.Document) *chromiumRequest {
	r.header = header
	return r
}

// Footer adds a footer to each page.
func (r *chromiumRequest) Footer(footer document.Document) *chromiumRequest {
	r.footer = footer
	return r
}

// PdfA sets the PDF/A format of the resulting PDF.
func (r *chromiumRequest) PdfA(pdfa PdfAFormat) *chromiumRequest {
	r.fields[fieldOfficePdfA] = string(pdfa)
	return r
}

// PdfUA enables PDF for Universal Access for optimal accessibility.
func (r *chromiumRequest) PdfUA() *chromiumRequest {
	r.fields[fieldOfficePdfUa] = strconv.FormatBool(true)
	return r
}

// Metadata sets the metadata to write.
func (r *chromiumRequest) Metadata(jsonData []byte) *chromiumRequest {
	r.fields[fieldMetadata] = string(jsonData)
	return r
}

// SplitIntervals splits the resulting PDF by interval.
func (r *chromiumRequest) SplitIntervals(span int) *chromiumRequest {
	r.fields[fieldSplitMode] = splitModeIntervals
	r.fields[fieldSplitSpan] = strconv.Itoa(span)

	return r
}

// SplitPages splits the resulting PDF by pages. Span should have the format “1-5”.
func (r *chromiumRequest) SplitPages(span string, unify bool) *chromiumRequest {
	r.fields[fieldSplitMode] = splitModePages
	r.fields[fieldSplitSpan] = span
	r.fields[fieldSplitUnify] = strconv.FormatBool(unify)

	return r
}

// GenerateTaggedPDF generates tagged (accessible) PDF.
func (r *chromiumRequest) GenerateTaggedPDF(val bool) *chromiumRequest {
	r.fields[fieldChromiumGenerateTaggedPDF] = strconv.FormatBool(val)
	return r
}

func (r *chromiumRequest) Encrypt(userPassword, ownerPassword string) *chromiumRequest {
	r.fields[fieldUserPassword] = userPassword
	r.fields[fieldOwnerPassword] = ownerPassword

	return r
}

// ScreenshotWidth Width sets the device screen width in pixels.
func (r *chromiumRequest) ScreenshotWidth(width int) *chromiumRequest {
	r.fields[fieldScreenshotWidth] = strconv.Itoa(width)
	return r
}

// ScreenshotHeight sets the device screen height in pixels.
func (r *chromiumRequest) ScreenshotHeight(height int) *chromiumRequest {
	r.fields[fieldScreenshotHeight] = strconv.Itoa(height)
	return r
}

// ScreenshotClip defines whether to clip the screenshot according to the device dimensions.
func (r *chromiumRequest) ScreenshotClip() *chromiumRequest {
	r.fields[fieldScreenshotClip] = strconv.FormatBool(true)
	return r
}

// ScreenshotQuality sets the compression quality from range 0 to 100 (jpeg only).
func (r *chromiumRequest) ScreenshotQuality(quality int) *chromiumRequest {
	r.fields[fieldScreenshotQuality] = strconv.Itoa(quality)
	return r
}

// ScreenshotOptimizeForSpeed defines whether to optimize image encoding for speed, not for resulting size.
func (r *chromiumRequest) ScreenshotOptimizeForSpeed() *chromiumRequest {
	r.fields[fieldScreenshotOptimizeForSpeed] = strconv.FormatBool(true)
	return r
}

// Format sets the image compression format, either PNG, JPEG or WEBP. Default is PNG.
func (r *chromiumRequest) Format(format ImageFormat) *chromiumRequest {
	r.fields[fieldScreenshotFormat] = string(format)
	return r
}

func mustJSON(data interface{}) string {
	b, err := json.Marshal(data)
	if err != nil {
		panic(fmt.Errorf("json marshaling failed: %w", err))
	}

	return string(b)
}
