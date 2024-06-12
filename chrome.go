package gotenberg

import (
	"fmt"
	"strconv"
	"time"
)

const (
	formFieldWaitDelay                  string = "waitDelay"
	formFieldPaperWidth                 string = "formFieldPaperWidth"
	formFieldPaperHeight                string = "paperHeight"
	formFieldMarginTop                  string = "marginTop"
	formFieldMarginBottom               string = "marginBottom"
	formFieldMarginLeft                 string = "marginLeft"
	formFieldMarginRight                string = "marginRight"
	formFieldLandscapeChrome            string = "landscape"
	formFieldPageRanges                 string = "pageRanges"
	formFieldGoogleChromeRpccBufferSize string = "googleChromeRpccBufferSize"
	formFieldScale                      string = "scale"
	formFieldSkipNetworkIdleEvent       string = "skipNetworkIdleEvent"
	formFieldSinglePage                 string = "singlePage"
	formFieldPreferCssPageSize          string = "preferCssPageSize"
	formFieldPrintBackground            string = "printBackground"
	formFieldOmitBackground             string = "omitBackground"
	formFieldFormat                     string = "format"
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

// nolint: gochecknoglobals
var (
	PNG  = "png"
	JPEG = "jpeg"
	WebP = "webp"
)

type chromeRequest struct {
	header Document
	footer Document

	*request
}

func newChromeRequest() *chromeRequest {
	return &chromeRequest{nil, nil, newRequest()}
}

// WaitDelay sets waitDelay form field.
func (req *chromeRequest) WaitDelay(delay time.Duration) {
	req.values[formFieldWaitDelay] = delay.String()
}

// Header sets header form file.
func (req *chromeRequest) Header(header Document) {
	req.header = header
}

// Footer sets footer form file.
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

// Landscape sets landscape form field.
func (req *chromeRequest) Landscape(isLandscape bool) {
	req.values[formFieldLandscapeChrome] = strconv.FormatBool(isLandscape)
}

// PageRanges sets pageRanges form field.
func (req *chromeRequest) PageRanges(ranges string) {
	req.values[formFieldPageRanges] = ranges
}

// GoogleChromeRpccBufferSize sets googleChromeRpccBufferSize form field.
func (req *chromeRequest) GoogleChromeRpccBufferSize(bufferSize int64) {
	req.values[formFieldGoogleChromeRpccBufferSize] = strconv.FormatInt(bufferSize, 10)
}

// Scale sets scale form field.
func (req *chromeRequest) Scale(scaleFactor float64) {
	req.values[formFieldScale] = fmt.Sprintf("%f", scaleFactor)
}

// SkipNetworkIdleEvent sets skipNetworkIdleEvent form field as true.
func (req *chromeRequest) SkipNetworkIdleEvent() {
	req.values[formFieldSkipNetworkIdleEvent] = "true"
}

// SinglePage sets singlePage form field as true.
func (req *chromeRequest) SinglePage() {
	req.values[formFieldSinglePage] = "true"
}

// Define whether to prefer page size as defined by CSS.
func (req *chromeRequest) PreferCssPageSize() {
	req.values[formFieldPreferCssPageSize] = "true"
}

// PrintBackground print the background graphics.
func (req *chromeRequest) PrintBackground() {
	req.values[formFieldPrintBackground] = "true"
}

// OmitBackground omit the background graphics.
func (req *chromeRequest) OmitBackground() {
	req.values[formFieldOmitBackground] = "true"
}

// Format sets format form field
func (req *chromeRequest) Format(format string) {
	req.values[formFieldFormat] = format
	if format == "" {
		req.values[format] = JPEG
	}
}

func (req *chromeRequest) Metadata(jsonData []byte) {
	req.values[formFieldMetadata] = string(jsonData)
}
