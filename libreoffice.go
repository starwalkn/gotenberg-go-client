package gotenberg

import (
	"strconv"

	"github.com/starwalkn/gotenberg-go-client/v8/document"
)

const endpointOfficeConvert = "/forms/libreoffice/convert"

// LibreOfficeRequest facilitates LibreOffice documents conversion with the Gotenberg API.
type LibreOfficeRequest struct {
	docs []document.Document

	*baseRequest
}

func NewLibreOfficeRequest(docs ...document.Document) *LibreOfficeRequest {
	return &LibreOfficeRequest{
		docs:        docs,
		baseRequest: newBaseRequest(),
	}
}

func (r *LibreOfficeRequest) endpoint() string {
	return endpointOfficeConvert
}

func (r *LibreOfficeRequest) formDocuments() map[string]document.Document {
	docs := make(map[string]document.Document)
	for _, doc := range r.docs {
		docs[doc.Name()] = doc
	}

	return docs
}

// Password sets the password for opening the source file.
func (r *LibreOfficeRequest) Password(password string) {
	r.fields[fieldOfficePassword] = password
}

// Landscape sets the paper orientation to landscape.
func (r *LibreOfficeRequest) Landscape() {
	r.fields[fieldOfficeLandscape] = strconv.FormatBool(true)
}

// NativePageRanges sets the page ranges to print, e.g., "1-4". Empty means all pages.
func (r *LibreOfficeRequest) NativePageRanges(ranges string) {
	r.fields[fieldOfficeNativePageRanges] = ranges
}

// ExportFormFields specifies whether form fields are exported as widgets
// or only their fixed print representation is exported.
func (r *LibreOfficeRequest) ExportFormFields(export bool) {
	r.fields[fieldOfficeExportFormFields] = strconv.FormatBool(export)
}

// AllowDuplicateFieldNames specifies whether multiple form fields exported are allowed to have the same field name.
func (r *LibreOfficeRequest) AllowDuplicateFieldNames() {
	r.fields[fieldOfficeAllowDuplicateFieldNames] = strconv.FormatBool(true)
}

// ExportBookmarks specifies if bookmarks are exported to PDF.
func (r *LibreOfficeRequest) ExportBookmarks(export bool) {
	r.fields[fieldOfficeExportBookmarks] = strconv.FormatBool(export)
}

// ExportBookmarksToPdfDestination specifies that the bookmarks contained
// in the source LibreOffice file should be exported to the PDF file as Named Destination.
func (r *LibreOfficeRequest) ExportBookmarksToPdfDestination() {
	r.fields[fieldOfficeExportBookmarksToPdfDestination] = strconv.FormatBool(true)
}

// ExportPlaceholders exports the placeholders fields visual markings only. The exported placeholder is ineffective.
func (r *LibreOfficeRequest) ExportPlaceholders() {
	r.fields[fieldOfficeExportPlaceholders] = strconv.FormatBool(true)
}

// ExportNotes specifies if notes are exported to PDF.
func (r *LibreOfficeRequest) ExportNotes() {
	r.fields[fieldOfficeExportNotes] = strconv.FormatBool(true)
}

// ExportNotesPages specifies if notes pages are exported to PDF. Notes pages are available in Impress documents only.
func (r *LibreOfficeRequest) ExportNotesPages() {
	r.fields[fieldOfficeExportNotesPages] = strconv.FormatBool(true)
}

// ExportOnlyNotesPages specifies, if the form field exportNotesPages
// is set to true, if only notes pages are exported to PDF.
func (r *LibreOfficeRequest) ExportOnlyNotesPages() {
	r.fields[fieldOfficeExportOnlyNotesPages] = strconv.FormatBool(true)
}

// ExportNotesInMargin specifies if notes in margin are exported to PDF.
func (r *LibreOfficeRequest) ExportNotesInMargin() {
	r.fields[fieldOfficeExportNotesInMargin] = strconv.FormatBool(true)
}

// ConvertOooTargetToPdfTarget specifies that the target documents with .od[tpgs] extension,
// will have that extension changed to .pdf when the link is exported to PDF. The source document remains untouched.
func (r *LibreOfficeRequest) ConvertOooTargetToPdfTarget() {
	r.fields[fieldOfficeConvertOooTargetToPdfTarget] = strconv.FormatBool(true)
}

// ExportLinksRelativeFsys specifies that the file system related hyperlinks (file:// protocol) present
// in the document will be exported as relative to the source document location.
func (r *LibreOfficeRequest) ExportLinksRelativeFsys() {
	r.fields[fieldOfficeExportLinksRelativeFsys] = strconv.FormatBool(true)
}

// ExportHiddenSlides exports, for LibreOffice Impress, slides that are not included in slide shows.
func (r *LibreOfficeRequest) ExportHiddenSlides() {
	r.fields[fieldOfficeExportHiddenSlides] = strconv.FormatBool(true)
}

// SkipEmptyPages Specifies that automatically inserted empty pages are suppressed.
// This option is active only if storing Writer documents.
func (r *LibreOfficeRequest) SkipEmptyPages() {
	r.fields[fieldOfficeSkipEmptyPages] = strconv.FormatBool(true)
}

// AddOriginalDocumentAsStream specifies that a stream is inserted to the PDF file
// which contains the original document for archiving purposes.
func (r *LibreOfficeRequest) AddOriginalDocumentAsStream() {
	r.fields[fieldOfficeAddOriginalDocumentAsStream] = strconv.FormatBool(true)
}

// SinglePageSheets ignores each sheet’s paper size, print ranges and shown/hidden
// status and puts every sheet (even hidden sheets) on exactly one page.
func (r *LibreOfficeRequest) SinglePageSheets() {
	r.fields[fieldOfficeSinglePageSheets] = strconv.FormatBool(true)
}

// LosslessImageCompression specifies if images are exported to PDF using a lossless compression
// format like PNG or compressed using the JPEG format.
func (r *LibreOfficeRequest) LosslessImageCompression() {
	r.fields[fieldOfficeLosslessImageCompression] = strconv.FormatBool(true)
}

// Quality specifies the quality of the JPG export. A higher value produces a higher-quality image and a larger file.
// Between 1 and 100.
func (r *LibreOfficeRequest) Quality(quality int) {
	r.fields[fieldOfficeQuality] = strconv.Itoa(quality)
}

// ReduceImageResolution Specifies if the resolution of each image is reduced to the resolution specified by the
// form field maxImageResolution.
func (r *LibreOfficeRequest) ReduceImageResolution() {
	r.fields[fieldOfficeReduceImageResolution] = strconv.FormatBool(true)
}

// MaxImageResolution If the form field reduceImageResolution is set to true, tells if all images will be reduced
// to the given value in DPI. Possible values are: 75, 150, 300, 600 and 1200.
func (r *LibreOfficeRequest) MaxImageResolution(res int) {
	r.fields[fieldOfficeMaxImageResolution] = strconv.Itoa(res)
}

// PdfA sets the PDF/A format of the resulting PDF.
func (r *LibreOfficeRequest) PdfA(pdfa PdfAFormat) {
	r.fields[fieldOfficePdfA] = string(pdfa)
}

// PdfUA enables PDF for Universal Access for optimal accessibility.
func (r *LibreOfficeRequest) PdfUA() {
	r.fields[fieldOfficePdfUa] = strconv.FormatBool(true)
}

// Metadata sets the metadata to write.
func (r *LibreOfficeRequest) Metadata(md []byte) {
	r.fields[fieldMetadata] = string(md)
}

// SplitIntervals splits the resulting PDF by interval.
func (r *LibreOfficeRequest) SplitIntervals(span int) {
	r.fields[fieldSplitMode] = SplitModeIntervals
	r.fields[fieldSplitSpan] = strconv.Itoa(span)
}

// SplitPages splits the resulting PDF by pages. Span should have the format “1-5”.
func (r *LibreOfficeRequest) SplitPages(span string, unify bool) {
	r.fields[fieldSplitMode] = SplitModePages
	r.fields[fieldSplitSpan] = span
	r.fields[fieldSplitUnify] = strconv.FormatBool(unify)
}

// Flatten defines whether the resulting PDF should be flattened.
func (r *LibreOfficeRequest) Flatten(val bool) {
	r.fields[fieldOfficeFlatten] = strconv.FormatBool(val)
}

// Merge merges the resulting PDFs.
func (r *LibreOfficeRequest) Merge() {
	r.fields[fieldOfficeMerge] = strconv.FormatBool(true)
}

// UpdateIndexes specifies whether to update the indexes before conversion, keeping in
// mind that doing so might result in missing links in the final PDF.
func (r *LibreOfficeRequest) UpdateIndexes(value bool) {
	r.fields[fieldOfficeUpdateIndexes] = strconv.FormatBool(value)
}

// Compile-time checks to ensure type implements desired interfaces.
var (
	_ = multipartRequest(new(LibreOfficeRequest))
)
