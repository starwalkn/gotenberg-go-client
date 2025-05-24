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
	return &chromiumRequest{
		header:      nil,
		footer:      nil,
		baseRequest: newBaseRequest(),
	}
}

// WaitDelay sets the duration (i.e., "1s", "2ms", etc.) to wait when loading an
// HTML document before converting it to PDF.
func (r *chromiumRequest) WaitDelay(delay time.Duration) {
	r.fields[fieldChromiumWaitDelay] = delay.String()
}

// WaitForExpression sets the JavaScript expression to wait before converting an
// HTML document into PDF until it returns true.
func (r *chromiumRequest) WaitForExpression(expression string) {
	r.fields[fieldChromiumWaitForExpression] = expression
}

// EmulatePrintMediaType forces Chromium to emulate the media type "print".
func (r *chromiumRequest) EmulatePrintMediaType() {
	r.fields[fieldChromiumEmulatedMediaType] = "print"
}

// EmulateScreenMediaType forces Chromium to emulate the media type "screen".
func (r *chromiumRequest) EmulateScreenMediaType() {
	r.fields[fieldChromiumEmulatedMediaType] = "screen"
}

// Cookies to store in the Chromium cookie jar.
func (r *chromiumRequest) Cookies(cookies []Cookie) error {
	for _, cookie := range cookies {
		if err := cookie.validate(); err != nil {
			return fmt.Errorf("validate cookies: %w", err)
		}
	}

	marshaledCookies, err := json.Marshal(cookies)
	if err != nil {
		return fmt.Errorf("marshal cookies to JSON: %w", err)
	}

	r.fields[fieldChromiumCookies] = string(marshaledCookies)

	return nil
}

// UserAgent overrides the default User-Agent HTTP header.
func (r *chromiumRequest) UserAgent(ua string) {
	r.fields[fieldChromiumUserAgent] = ua
}

// ExtraHTTPHeaders sets extra HTTP headers that Chromium will send when loading the HTML document.
func (r *chromiumRequest) ExtraHTTPHeaders(headers map[string]string) error {
	marshaledHeaders, err := json.Marshal(headers)
	if err != nil {
		return fmt.Errorf("marshal headers to JSON: %w", err)
	}

	r.fields[fieldChromiumExtraHTTPHeaders] = string(marshaledHeaders)

	return nil
}

// FailOnHTTPStatusCodes forces Gotenberg to return a 409 Conflict response
// if the HTTP status code from the main page is not acceptable.
func (r *chromiumRequest) FailOnHTTPStatusCodes(statusCodes []int) error {
	marshaledStatusCodes, err := json.Marshal(statusCodes)
	if err != nil {
		return fmt.Errorf("marshal HTTP status codes to JSON: %w", err)
	}

	r.fields[fieldChromiumFailOnHTTPStatusCodes] = string(marshaledStatusCodes)

	return nil
}

// FailOnResourceHTTPStatusCodes forces Gotenberg to return a 409 Conflict response
// if the HTTP status code from at least one resource is not acceptable.
func (r *chromiumRequest) FailOnResourceHTTPStatusCodes(statusCodes []int) error {
	marshaledStatusCodes, err := json.Marshal(statusCodes)
	if err != nil {
		return fmt.Errorf("marshal HTTP status codes to JSON: %w", err)
	}

	r.fields[fieldChromiumFailOnResourceHTTPStatusCodes] = string(marshaledStatusCodes)

	return nil
}

// FailOnConsoleExceptions forces Gotenberg to return a 409 Conflict response
// if there are exceptions in the Chromium console.
func (r *chromiumRequest) FailOnConsoleExceptions() {
	r.fields[fieldChromiumFailOnConsoleExceptions] = strconv.FormatBool(true)
}

// FailOnResourceLoadingFailed forces Gotenberg to return a 409 Conflict if Chromium
// fails to load at least one resource.
func (r *chromiumRequest) FailOnResourceLoadingFailed() {
	r.fields[fieldChromiumFailOnResourceLoadingFailed] = strconv.FormatBool(true)
}

// SkipNetworkIdleEvent specifies whether Chromium have to wait or not for its network to be idle.
// Enabled by default in Gotenberg >= 8.11.0.
func (r *chromiumRequest) SkipNetworkIdleEvent(val bool) {
	r.fields[fieldChromiumSkipNetworkIdleEvent] = strconv.FormatBool(val)
}

// SinglePage defines whether to print the entire content in one single page.
func (r *chromiumRequest) SinglePage() {
	r.fields[fieldChromiumSinglePage] = strconv.FormatBool(true)
}

// PaperSize sets paperWidth and paperHeight form fields with the provided unit.
// If unit is empty, it defaults to inches. Default is Letter (8.5 x 11 inches).
func (r *chromiumRequest) PaperSize(size PaperDimensions) {
	if size.Unit == "" {
		size.Unit = IN
	}

	r.fields[fieldChromiumPaperWidth] = fmt.Sprintf("%f%s", size.Width, size.Unit)
	r.fields[fieldChromiumPaperHeight] = fmt.Sprintf("%f%s", size.Height, size.Unit)
}

// Margins sets marginTop, marginBottom, marginLeft and marginRight form fields. Default unit is inches.
func (r *chromiumRequest) Margins(margins PageMargins) {
	if margins.Unit == "" {
		margins.Unit = IN
	}

	r.fields[fieldChromiumMarginTop] = fmt.Sprintf("%f%s", margins.Top, margins.Unit)
	r.fields[fieldChromiumMarginBottom] = fmt.Sprintf("%f%s", margins.Bottom, margins.Unit)
	r.fields[fieldChromiumMarginLeft] = fmt.Sprintf("%f%s", margins.Left, margins.Unit)
	r.fields[fieldChromiumMarginRight] = fmt.Sprintf("%f%s", margins.Right, margins.Unit)
}

// PreferCSSPageSize forces page size as defined by CSS.
func (r *chromiumRequest) PreferCSSPageSize() {
	r.fields[fieldChromiumPreferCSSPageSize] = strconv.FormatBool(true)
}

// PrintBackground prints the background graphics.
func (r *chromiumRequest) PrintBackground() {
	r.fields[fieldChromiumPrintBackground] = strconv.FormatBool(true)
}

// OmitBackground hides default white background and allows generating PDFs with transparency.
func (r *chromiumRequest) OmitBackground() {
	r.fields[fieldChromiumOmitBackground] = strconv.FormatBool(true)
}

// Landscape sets the paper orientation to landscape.
func (r *chromiumRequest) Landscape() {
	r.fields[fieldChromiumLandscapeChrome] = strconv.FormatBool(true)
}

// Scale overrides the default scale of the page rendering (i.e., 1.0).
func (r *chromiumRequest) Scale(factor float64) {
	r.fields[fieldChromiumScale] = fmt.Sprintf("%f", factor)
}

// NativePageRanges sets the page ranges to print, e.g., "1-5, 8, 11-13". Empty means all pages.
func (r *chromiumRequest) NativePageRanges(ranges string) {
	r.fields[fieldChromiumNativePageRanges] = ranges
}

// GenerateDocumentOutline embeds the document outline into the PDF.
func (r *chromiumRequest) GenerateDocumentOutline() {
	r.fields[fieldChromiumGenerateDocumentOutline] = strconv.FormatBool(true)
}

// Header adds a header to each page.
func (r *chromiumRequest) Header(header document.Document) {
	r.header = header
}

// Footer adds a footer to each page.
func (r *chromiumRequest) Footer(footer document.Document) {
	r.footer = footer
}

// PdfA sets the PDF/A format of the resulting PDF.
func (r *chromiumRequest) PdfA(pdfa PdfAFormat) {
	r.fields[fieldOfficePdfA] = string(pdfa)
}

// PdfUA enables PDF for Universal Access for optimal accessibility.
func (r *chromiumRequest) PdfUA() {
	r.fields[fieldOfficePdfUa] = strconv.FormatBool(true)
}

// Metadata sets the metadata to write.
func (r *chromiumRequest) Metadata(jsonData []byte) {
	r.fields[fieldMetadata] = string(jsonData)
}

// SplitIntervals splits the resulting PDF by interval.
func (r *chromiumRequest) SplitIntervals(span int) {
	r.fields[fieldSplitMode] = SplitModeIntervals
	r.fields[fieldSplitSpan] = strconv.Itoa(span)
}

// SplitPages splits the resulting PDF by pages. Span should have the format “1-5”.
func (r *chromiumRequest) SplitPages(span string, unify bool) {
	r.fields[fieldSplitMode] = SplitModePages
	r.fields[fieldSplitSpan] = span
	r.fields[fieldSplitUnify] = strconv.FormatBool(unify)
}

// GenerateTaggedPDF generates tagged (accessible) PDF.
func (r *chromiumRequest) GenerateTaggedPDF(val bool) {
	r.fields[fieldChromiumGenerateTaggedPDF] = strconv.FormatBool(val)
}

// ScreenshotWidth Width sets the device screen width in pixels.
func (r *chromiumRequest) ScreenshotWidth(width int) {
	r.fields[fieldScreenshotWidth] = strconv.Itoa(width)
}

// ScreenshotHeight sets the device screen height in pixels.
func (r *chromiumRequest) ScreenshotHeight(height int) {
	r.fields[fieldScreenshotHeight] = strconv.Itoa(height)
}

// ScreenshotClip defines whether to clip the screenshot according to the device dimensions.
func (r *chromiumRequest) ScreenshotClip() {
	r.fields[fieldScreenshotClip] = strconv.FormatBool(true)
}

// ScreenshotQuality sets the compression quality from range 0 to 100 (jpeg only).
func (r *chromiumRequest) ScreenshotQuality(quality int) {
	r.fields[fieldScreenshotQuality] = strconv.Itoa(quality)
}

// ScreenshotOptimizeForSpeed defines whether to optimize image encoding for speed, not for resulting size.
func (r *chromiumRequest) ScreenshotOptimizeForSpeed() {
	r.fields[fieldScreenshotOptimizeForSpeed] = strconv.FormatBool(true)
}

// ScreenshotFormat sets the image compression format, either PNG, JPEG or WEBP. Default is PNG.
func (r *chromiumRequest) ScreenshotFormat(format ImageFormat) {
	r.fields[fieldScreenshotFormat] = string(format)
}
