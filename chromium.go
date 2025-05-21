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
func (req *chromiumRequest) WaitDelay(delay time.Duration) {
	req.fields[fieldChromiumWaitDelay] = delay.String()
}

// WaitForExpression sets the JavaScript expression to wait before converting an
// HTML document into PDF until it returns true.
func (req *chromiumRequest) WaitForExpression(expression string) {
	req.fields[fieldChromiumWaitForExpression] = expression
}

// EmulatePrintMediaType forces Chromium to emulate the media type "print".
func (req *chromiumRequest) EmulatePrintMediaType() {
	req.fields[fieldChromiumEmulatedMediaType] = "print"
}

// EmulateScreenMediaType forces Chromium to emulate the media type "screen".
func (req *chromiumRequest) EmulateScreenMediaType() {
	req.fields[fieldChromiumEmulatedMediaType] = "screen"
}

// Cookies to store in the Chromium cookie jar.
func (req *chromiumRequest) Cookies(cookies []Cookie) error {
	for _, cookie := range cookies {
		if err := cookie.validate(); err != nil {
			return fmt.Errorf("validate cookies: %w", err)
		}
	}

	marshaledCookies, err := json.Marshal(cookies)
	if err != nil {
		return fmt.Errorf("marshal cookies to JSON: %w", err)
	}

	req.fields[fieldChromiumCookies] = string(marshaledCookies)

	return nil
}

// UserAgent overrides the default User-Agent HTTP header.
func (req *chromiumRequest) UserAgent(ua string) {
	req.fields[fieldChromiumUserAgent] = ua
}

// ExtraHTTPHeaders sets extra HTTP headers that Chromium will send when loading the HTML document.
func (req *chromiumRequest) ExtraHTTPHeaders(headers map[string]string) error {
	marshaledHeaders, err := json.Marshal(headers)
	if err != nil {
		return fmt.Errorf("marshal headers to JSON: %w", err)
	}

	req.fields[fieldChromiumExtraHTTPHeaders] = string(marshaledHeaders)

	return nil
}

// FailOnHTTPStatusCodes forces Gotenberg to return a 409 Conflict response
// if the HTTP status code from the main page is not acceptable.
func (req *chromiumRequest) FailOnHTTPStatusCodes(statusCodes []int) error {
	marshaledStatusCodes, err := json.Marshal(statusCodes)
	if err != nil {
		return fmt.Errorf("marshal HTTP status codes to JSON: %w", err)
	}

	req.fields[fieldChromiumFailOnHTTPStatusCodes] = string(marshaledStatusCodes)

	return nil
}

// FailOnResourceHTTPStatusCodes forces Gotenberg to return a 409 Conflict response
// if the HTTP status code from at least one resource is not acceptable.
func (req *chromiumRequest) FailOnResourceHTTPStatusCodes(statusCodes []int) error {
	marshaledStatusCodes, err := json.Marshal(statusCodes)
	if err != nil {
		return fmt.Errorf("marshal HTTP status codes to JSON: %w", err)
	}

	req.fields[fieldChromiumFailOnResourceHTTPStatusCodes] = string(marshaledStatusCodes)

	return nil
}

// FailOnConsoleExceptions forces Gotenberg to return a 409 Conflict response
// if there are exceptions in the Chromium console.
func (req *chromiumRequest) FailOnConsoleExceptions() {
	req.fields[fieldChromiumFailOnConsoleExceptions] = strconv.FormatBool(true)
}

// FailOnResourceLoadingFailed forces Gotenberg to return a 409 Conflict if Chromium
// fails to load at least one resource.
func (req *chromiumRequest) FailOnResourceLoadingFailed() {
	req.fields[fieldChromiumFailOnResourceLoadingFailed] = strconv.FormatBool(true)
}

// SkipNetworkIdleEvent specifies whether Chromium have to wait or not for its network to be idle.
// Enabled by default in Gotenberg >= 8.11.0.
func (req *chromiumRequest) SkipNetworkIdleEvent(val bool) {
	req.fields[fieldChromiumSkipNetworkIdleEvent] = strconv.FormatBool(val)
}

// SinglePage defines whether to print the entire content in one single page.
func (req *chromiumRequest) SinglePage() {
	req.fields[fieldChromiumSinglePage] = strconv.FormatBool(true)
}

// PaperSize sets paperWidth and paperHeight form fields with the provided unit.
// If unit is empty, it defaults to inches. Default is Letter (8.5 x 11 inches).
func (req *chromiumRequest) PaperSize(size PaperDimensions) {
	if size.Unit == "" {
		size.Unit = IN
	}

	req.fields[fieldChromiumPaperWidth] = fmt.Sprintf("%f%s", size.Width, size.Unit)
	req.fields[fieldChromiumPaperHeight] = fmt.Sprintf("%f%s", size.Height, size.Unit)
}

// Margins sets marginTop, marginBottom, marginLeft and marginRight form fields. Default unit is inches.
func (req *chromiumRequest) Margins(margins PageMargins) {
	if margins.Unit == "" {
		margins.Unit = IN
	}

	req.fields[fieldChromiumMarginTop] = fmt.Sprintf("%f%s", margins.Top, margins.Unit)
	req.fields[fieldChromiumMarginBottom] = fmt.Sprintf("%f%s", margins.Bottom, margins.Unit)
	req.fields[fieldChromiumMarginLeft] = fmt.Sprintf("%f%s", margins.Left, margins.Unit)
	req.fields[fieldChromiumMarginRight] = fmt.Sprintf("%f%s", margins.Right, margins.Unit)
}

// PreferCSSPageSize forces page size as defined by CSS.
func (req *chromiumRequest) PreferCSSPageSize() {
	req.fields[fieldChromiumPreferCSSPageSize] = strconv.FormatBool(true)
}

// PrintBackground prints the background graphics.
func (req *chromiumRequest) PrintBackground() {
	req.fields[fieldChromiumPrintBackground] = strconv.FormatBool(true)
}

// OmitBackground hides default white background and allows generating PDFs with transparency.
func (req *chromiumRequest) OmitBackground() {
	req.fields[fieldChromiumOmitBackground] = strconv.FormatBool(true)
}

// Landscape sets the paper orientation to landscape.
func (req *chromiumRequest) Landscape() {
	req.fields[fieldChromiumLandscapeChrome] = strconv.FormatBool(true)
}

// Scale overrides the default scale of the page rendering (i.e., 1.0).
func (req *chromiumRequest) Scale(factor float64) {
	req.fields[fieldChromiumScale] = fmt.Sprintf("%f", factor)
}

// NativePageRanges sets the page ranges to print, e.g., "1-5, 8, 11-13". Empty means all pages.
func (req *chromiumRequest) NativePageRanges(ranges string) {
	req.fields[fieldChromiumNativePageRanges] = ranges
}

// GenerateDocumentOutline embeds the document outline into the PDF.
func (req *chromiumRequest) GenerateDocumentOutline() {
	req.fields[fieldChromiumGenerateDocumentOutline] = strconv.FormatBool(true)
}

// Header adds a header to each page.
func (req *chromiumRequest) Header(header document.Document) {
	req.header = header
}

// Footer adds a footer to each page.
func (req *chromiumRequest) Footer(footer document.Document) {
	req.footer = footer
}

// PdfA sets the PDF/A format of the resulting PDF.
func (req *chromiumRequest) PdfA(pdfa PdfAFormat) {
	req.fields[fieldOfficePdfA] = string(pdfa)
}

// PdfUA enables PDF for Universal Access for optimal accessibility.
func (req *chromiumRequest) PdfUA() {
	req.fields[fieldOfficePdfUa] = strconv.FormatBool(true)
}

// Metadata sets the metadata to write.
func (req *chromiumRequest) Metadata(jsonData []byte) {
	req.fields[fieldMetadata] = string(jsonData)
}

// SplitIntervals splits the resulting PDF by interval.
func (req *chromiumRequest) SplitIntervals(span int) {
	req.fields[fieldSplitMode] = SplitModeIntervals
	req.fields[fieldSplitSpan] = strconv.Itoa(span)
}

// SplitPages splits the resulting PDF by pages. Span should have the format “1-5”.
func (req *chromiumRequest) SplitPages(span string, unify bool) {
	req.fields[fieldSplitMode] = SplitModePages
	req.fields[fieldSplitSpan] = span
	req.fields[fieldSplitUnify] = strconv.FormatBool(unify)
}

// GenerateTaggedPDF generates tagged (accessible) PDF.
func (req *chromiumRequest) GenerateTaggedPDF(val bool) {
	req.fields[fieldChromiumGenerateTaggedPDF] = strconv.FormatBool(val)
}

// ScreenshotWidth Width sets the device screen width in pixels.
func (req *chromiumRequest) ScreenshotWidth(width int) {
	req.fields[fieldScreenshotWidth] = strconv.Itoa(width)
}

// ScreenshotHeight sets the device screen height in pixels.
func (req *chromiumRequest) ScreenshotHeight(height int) {
	req.fields[fieldScreenshotHeight] = strconv.Itoa(height)
}

// ScreenshotClip defines whether to clip the screenshot according to the device dimensions.
func (req *chromiumRequest) ScreenshotClip() {
	req.fields[fieldScreenshotClip] = strconv.FormatBool(true)
}

// ScreenshotQuality sets the compression quality from range 0 to 100 (jpeg only).
func (req *chromiumRequest) ScreenshotQuality(quality int) {
	req.fields[fieldScreenshotQuality] = strconv.Itoa(quality)
}

// ScreenshotOptimizeForSpeed defines whether to optimize image encoding for speed, not for resulting size.
func (req *chromiumRequest) ScreenshotOptimizeForSpeed() {
	req.fields[fieldScreenshotOptimizeForSpeed] = strconv.FormatBool(true)
}

// Format sets the image compression format, either PNG, JPEG or WEBP. Default is PNG.
func (req *chromiumRequest) Format(format ImageFormat) {
	req.fields[fieldScreenshotFormat] = string(format)
}
