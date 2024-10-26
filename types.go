package gotenberg

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
	// NormalMargins uses 1-inch margins.
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

type PdfAFormat string

const (
	// Deprecated: Beginning with version 7.6, LibreOffice has discontinued support for PDF/A-1a.
	PdfA1b PdfAFormat = "PDF/A-1b"
	PdfA2b PdfAFormat = "PDF/A-2b"
	PdfA3b PdfAFormat = "PDF/A-3b"
)
