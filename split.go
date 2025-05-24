package gotenberg

import (
	"fmt"
	"strconv"

	"github.com/starwalkn/gotenberg-go-client/v8/document"
)

type SplitRequest struct {
	mode string
	docs []document.Document

	*baseRequest
}

// NewSplitRequest creates a new SplitRequest for the Gotenberg API.
//
// The "mode" parameter defines how the PDF documents will be split.
// It must be either "pages" or "intervals".
// For convenience, you can use the package constants gotenberg.SplitModePages
// or gotenberg.SplitModeIntervals.
// An error will be returned if an invalid mode string is provided.
//
// The "docs" parameter is a slice of PDF documents to be split.
func NewSplitRequest(mode string, docs ...document.Document) (*SplitRequest, error) {
	if mode != SplitModePages && mode != SplitModeIntervals {
		return nil, fmt.Errorf("split mode can only be 'pages' or 'intervals', not '%s'", mode)
	}

	br := newBaseRequest()
	br.fields[fieldSplitMode] = mode

	return &SplitRequest{
		mode:        mode,
		docs:        docs,
		baseRequest: br,
	}, nil
}

func (r *SplitRequest) endpoint() string {
	return "/forms/pdfengines/split"
}

func (r *SplitRequest) formDocuments() map[string]document.Document {
	files := make(map[string]document.Document)

	for _, doc := range r.docs {
		files[doc.Name()] = doc
	}

	return files
}

// SpanPages sets the span for the "pages" split mode.
// This method is only applicable when the SplitRequest was initialized with "pages" mode.
//
// The "span" parameter expects a string representing page ranges,
// for example, "1-4", "2-8", "1-6".
//
// If this method is called with "intervals" mode,
// the form field will be ignored, as the "span" value for pages mode is not relevant for "intervals" mode.
func (r *SplitRequest) SpanPages(span string) {
	if r.mode != SplitModePages {
		return
	}

	r.fields[fieldSplitSpan] = span
}

// SpanIntervals sets the span for the "intervals" split mode.
// This method is only applicable when the SplitRequest was initialized with "intervals" mode.
//
// The "span" parameter expects an integer representing the number of pages per resulting PDF file.
// For example, "3" would split the original PDF into chunks of 3 pages each.
//
// If this method is called with "pages" mode,
// the form field will be ignored, as the integer "span" value for intervals mode is not relevant for "pages" mode.
func (r *SplitRequest) SpanIntervals(span int) {
	if r.mode != SplitModeIntervals {
		return
	}

	r.fields[fieldSplitSpan] = strconv.Itoa(span)
}

func (r *SplitRequest) SplitUnify(val bool) {
	r.fields[fieldSplitUnify] = strconv.FormatBool(val)
}

// Flatten defines whether the resulting PDF should be flattened.
func (r *SplitRequest) Flatten(val bool) {
	r.fields[fieldSplitFlatten] = strconv.FormatBool(val)
}
