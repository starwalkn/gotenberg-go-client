package gotenberg

import (
	"fmt"
	"strconv"
	"time"
)

const (
	formFieldWaitDelay            string = "waitDelay"
	formFieldPaperWidth           string = "paperWidth"
	formFieldPaperHeight          string = "paperHeight"
	formFieldMarginTop            string = "marginTop"
	formFieldMarginBottom         string = "marginBottom"
	formFieldMarginLeft           string = "marginLeft"
	formFieldMarginRight          string = "marginRight"
	formFieldLandscapeChrome      string = "landscape"
	formFieldNativePageRanges     string = "nativePageRanges"
	formFieldScale                string = "scale"
	formFieldSkipNetworkIdleEvent string = "skipNetworkIdleEvent"
	formFieldSinglePage           string = "singlePage"
	formFieldPreferCSSPageSize    string = "preferCssPageSize"
	formFieldPrintBackground      string = "printBackground"
	formFieldOmitBackground       string = "omitBackground"
	formFieldFormat               string = "format"
)

type SizeUnit string

const (
	PT SizeUnit = "pt" // Points.
	PX SizeUnit = "px" // Pixels.
	IN SizeUnit = "in" // Inches.
	MM SizeUnit = "mm" // Millimeters.
	CM SizeUnit = "cm" // Centimeters.
	PC SizeUnit = "pc" // Picas.
)

type PaperDimensions struct {
	Width  float64
	Height float64
	Unit   SizeUnit
}

// nolint: gochecknoglobals
var (
	// A0 paper size.
	A0 = PaperDimensions{
		Height: 46.8,
		Width:  33.1,
		Unit:   IN,
	}
	// A1 paper size.
	A1 = PaperDimensions{
		Height: 33.1,
		Width:  23.4,
		Unit:   IN,
	}
	// A2 paper size.
	A2 = PaperDimensions{
		Height: 23.4,
		Width:  16.5,
		Unit:   IN,
	}
	// A3 paper size.
	A3 = PaperDimensions{
		Height: 16.5,
		Width:  11.7,
		Unit:   IN,
	}
	// A4 paper size.
	A4 = PaperDimensions{
		Height: 11.7,
		Width:  8.27,
		Unit:   IN,
	}
	// A5 paper size.
	A5 = PaperDimensions{
		Height: 8.3,
		Width:  5.8,
		Unit:   IN,
	}
	// A6 paper size.
	A6 = PaperDimensions{
		Height: 5.8,
		Width:  4.1,
		Unit:   IN,
	}
	// Letter paper size.
	Letter = PaperDimensions{
		Height: 11,
		Width:  8.5,
		Unit:   IN,
	}
	// Legal paper size.
	Legal = PaperDimensions{
		Height: 14,
		Width:  8.5,
		Unit:   IN,
	}
	// Tabloid paper size.
	Tabloid = PaperDimensions{
		Height: 17,
		Width:  11,
		Unit:   IN,
	}
	// Ledger paper size.
	Ledger = PaperDimensions{
		Height: 17,
		Width:  11,
		Unit:   IN,
	}
)

type PageMargins struct {
	Top    float64
	Bottom float64
	Left   float64
	Right  float64
	Unit   SizeUnit
}

// nolint: gochecknoglobals
var (
	// NoMargins removes margins.
	NoMargins = PageMargins{
		Top:    0,
		Bottom: 0,
		Left:   0,
		Right:  0,
		Unit:   IN,
	}
	// NormalMargins uses 1 inch margins.
	NormalMargins = PageMargins{
		Top:    1,
		Bottom: 1,
		Left:   1,
		Right:  1,
		Unit:   IN,
	}
	// LargeMargins uses 2 inch margins.
	LargeMargins = PageMargins{
		Top:    2,
		Bottom: 2,
		Left:   2,
		Right:  2,
		Unit:   IN,
	}
)

type ImageFormat string

const (
	PNG  ImageFormat = "png"
	JPEG ImageFormat = "jpeg"
	WebP ImageFormat = "webp"
)

type chromeRequest struct {
	header Document
	footer Document

	*request
}

func newChromeRequest() *chromeRequest {
	return &chromeRequest{nil, nil, newRequest()}
}

// Duration (e.g, '5s') to wait when loading an HTML document before converting it into PDF.
func (req *chromeRequest) WaitDelay(delay time.Duration) {
	req.values[formFieldWaitDelay] = delay.String()
}

// HTML file containing the header.
func (req *chromeRequest) Header(header Document) {
	req.header = header
}

// HTML file containing the footer.
func (req *chromeRequest) Footer(footer Document) {
	req.footer = footer
}

// PaperSize sets paperWidth and paperHeight form fields with the provided unit.
// If unit is empty, it defaults to inches. Default is Letter (8.5 x 11 inches).
func (req *chromeRequest) PaperSize(size PaperDimensions) {
	if size.Unit == "" {
		size.Unit = IN
	}
	req.values[formFieldPaperWidth] = fmt.Sprintf("%f%s", size.Width, size.Unit)
	req.values[formFieldPaperHeight] = fmt.Sprintf("%f%s", size.Height, size.Unit)
}

// Margins sets marginTop, marginBottom,
// marginLeft and marginRight form fields.
// Default unit is inches.
func (req *chromeRequest) Margins(margins PageMargins) {
	if margins.Unit == "" {
		margins.Unit = IN
	}
	req.values[formFieldMarginTop] = fmt.Sprintf("%f%s", margins.Top, margins.Unit)
	req.values[formFieldMarginBottom] = fmt.Sprintf("%f%s", margins.Bottom, margins.Unit)
	req.values[formFieldMarginLeft] = fmt.Sprintf("%f%s", margins.Left, margins.Unit)
	req.values[formFieldMarginRight] = fmt.Sprintf("%f%s", margins.Right, margins.Unit)
}

// Set the paper orientation to landscape.
func (req *chromeRequest) Landscape(isLandscape bool) {
	req.values[formFieldLandscapeChrome] = strconv.FormatBool(isLandscape)
}

// Page ranges to print, e.g., '1-5, 8, 11-13' - empty means all pages.
func (req *chromeRequest) PageRanges(ranges string) {
	req.values[formFieldNativePageRanges] = ranges
}

// The scale of the page rendering.
// Default is 1.0.
func (req *chromeRequest) Scale(scaleFactor float64) {
	req.values[formFieldScale] = fmt.Sprintf("%f", scaleFactor)
}

// Do not wait for Chromium network to be idle.
func (req *chromeRequest) SkipNetworkIdleEvent() {
	req.values[formFieldSkipNetworkIdleEvent] = strconv.FormatBool(true)
}

// Print the entire content in one single page.
func (req *chromeRequest) SinglePage() {
	req.values[formFieldSinglePage] = strconv.FormatBool(true)
}

// Prefer page size as defined by CSS.
func (req *chromeRequest) PreferCssPageSize() {
	req.values[formFieldPreferCSSPageSize] = strconv.FormatBool(true)
}

// Print the background graphics.
func (req *chromeRequest) PrintBackground() {
	req.values[formFieldPrintBackground] = strconv.FormatBool(true)
}

// Hide the default white background and allow generating PDFs/screenshots with transparency.
func (req *chromeRequest) OmitBackground() {
	req.values[formFieldOmitBackground] = strconv.FormatBool(true)
}

// The image compression format, either "png", "jpeg" or "webp".
// Default is "png".
func (req *chromeRequest) Format(format ImageFormat) {
	req.values[formFieldFormat] = string(format)
}

// The metadata to write (JSON format).
func (req *chromeRequest) Metadata(jsonData []byte) {
	req.values[formFieldMetadata] = string(jsonData)
}

// Convert the resulting PDF into the given PDF/A format.
func (req *chromeRequest) PdfA(pdfa PdfAFormat) {
	req.values[formFieldPdfa] = string(pdfa)
}

// Enable PDF for Universal Access for optimal accessibility.
func (req *chromeRequest) PdfUA() {
	req.values[formFieldPdfUa] = strconv.FormatBool(true)
}
