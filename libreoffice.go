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

func (req *LibreOfficeRequest) endpoint() string {
	return endpointOfficeConvert
}

func (req *LibreOfficeRequest) formDocuments() map[string]document.Document {
	files := make(map[string]document.Document)
	for _, doc := range req.docs {
		files[doc.Filename()] = doc
	}

	return files
}

// Password sets the password for opening the source file.
func (req *LibreOfficeRequest) Password(password string) {
	req.fields[fieldOfficePassword] = password
}

// Landscape sets the paper orientation to landscape.
func (req *LibreOfficeRequest) Landscape() {
	req.fields[fieldOfficeLandscape] = strconv.FormatBool(true)
}

// NativePageRanges sets the page ranges to print, e.g., "1-4". Empty means all pages.
func (req *LibreOfficeRequest) NativePageRanges(ranges string) {
	req.fields[fieldOfficeNativePageRanges] = ranges
}

// ExportFormFields specifies whether form fields are exported as widgets
// or only their fixed print representation is exported.
func (req *LibreOfficeRequest) ExportFormFields(export bool) {
	req.fields[fieldOfficeExportFormFields] = strconv.FormatBool(export)
}

// AllowDuplicateFieldNames specifies whether multiple form fields exported are allowed to have the same field name.
func (req *LibreOfficeRequest) AllowDuplicateFieldNames() {
	req.fields[fieldOfficeAllowDuplicateFieldNames] = strconv.FormatBool(true)
}

// ExportBookmarks specifies if bookmarks are exported to PDF.
func (req *LibreOfficeRequest) ExportBookmarks(export bool) {
	req.fields[fieldOfficeExportBookmarks] = strconv.FormatBool(export)
}

// ExportBookmarksToPdfDestination specifies that the bookmarks contained
// in the source LibreOffice file should be exported to the PDF file as Named Destination.
func (req *LibreOfficeRequest) ExportBookmarksToPdfDestination() {
	req.fields[fieldOfficeExportBookmarksToPdfDestination] = strconv.FormatBool(true)
}

// ExportPlaceholders exports the placeholders fields visual markings only. The exported placeholder is ineffective.
func (req *LibreOfficeRequest) ExportPlaceholders() {
	req.fields[fieldOfficeExportPlaceholders] = strconv.FormatBool(true)
}

// ExportNotes specifies if notes are exported to PDF.
func (req *LibreOfficeRequest) ExportNotes() {
	req.fields[fieldOfficeExportNotes] = strconv.FormatBool(true)
}

// ExportNotesPages specifies if notes pages are exported to PDF. Notes pages are available in Impress documents only.
func (req *LibreOfficeRequest) ExportNotesPages() {
	req.fields[fieldOfficeExportNotesPages] = strconv.FormatBool(true)
}

// ExportOnlyNotesPages specifies, if the form field exportNotesPages
// is set to true, if only notes pages are exported to PDF.
func (req *LibreOfficeRequest) ExportOnlyNotesPages() {
	req.fields[fieldOfficeExportOnlyNotesPages] = strconv.FormatBool(true)
}

// ExportNotesInMargin specifies if notes in margin are exported to PDF.
func (req *LibreOfficeRequest) ExportNotesInMargin() {
	req.fields[fieldOfficeExportNotesInMargin] = strconv.FormatBool(true)
}

// ConvertOooTargetToPdfTarget specifies that the target documents with .od[tpgs] extension,
// will have that extension changed to .pdf when the link is exported to PDF. The source document remains untouched.
func (req *LibreOfficeRequest) ConvertOooTargetToPdfTarget() {
	req.fields[fieldOfficeConvertOooTargetToPdfTarget] = strconv.FormatBool(true)
}

// ExportLinksRelativeFsys specifies that the file system related hyperlinks (file:// protocol) present
// in the document will be exported as relative to the source document location.
func (req *LibreOfficeRequest) ExportLinksRelativeFsys() {
	req.fields[fieldOfficeExportLinksRelativeFsys] = strconv.FormatBool(true)
}

// ExportHiddenSlides exports, for LibreOffice Impress, slides that are not included in slide shows.
func (req *LibreOfficeRequest) ExportHiddenSlides() {
	req.fields[fieldOfficeExportHiddenSlides] = strconv.FormatBool(true)
}

// SkipEmptyPages Specifies that automatically inserted empty pages are suppressed.
// This option is active only if storing Writer documents.
func (req *LibreOfficeRequest) SkipEmptyPages() {
	req.fields[fieldOfficeSkipEmptyPages] = strconv.FormatBool(true)
}

// AddOriginalDocumentAsStream specifies that a stream is inserted to the PDF file
// which contains the original document for archiving purposes.
func (req *LibreOfficeRequest) AddOriginalDocumentAsStream() {
	req.fields[fieldOfficeAddOriginalDocumentAsStream] = strconv.FormatBool(true)
}

// SinglePageSheets ignores each sheet’s paper size, print ranges and shown/hidden
// status and puts every sheet (even hidden sheets) on exactly one page.
func (req *LibreOfficeRequest) SinglePageSheets() {
	req.fields[fieldOfficeSinglePageSheets] = strconv.FormatBool(true)
}

// LosslessImageCompression specifies if images are exported to PDF using a lossless compression
// format like PNG or compressed using the JPEG format.
func (req *LibreOfficeRequest) LosslessImageCompression() {
	req.fields[fieldOfficeLosslessImageCompression] = strconv.FormatBool(true)
}

// Quality specifies the quality of the JPG export. A higher value produces a higher-quality image and a larger file.
// Between 1 and 100.
func (req *LibreOfficeRequest) Quality(quality int) {
	req.fields[fieldOfficeQuality] = strconv.Itoa(quality)
}

// ReduceImageResolution Specifies if the resolution of each image is reduced to the resolution specified by the
// form field maxImageResolution.
func (req *LibreOfficeRequest) ReduceImageResolution() {
	req.fields[fieldOfficeReduceImageResolution] = strconv.FormatBool(true)
}

// MaxImageResolution If the form field reduceImageResolution is set to true, tells if all images will be reduced
// to the given value in DPI. Possible values are: 75, 150, 300, 600 and 1200.
func (req *LibreOfficeRequest) MaxImageResolution(res int) {
	req.fields[fieldOfficeMaxImageResolution] = strconv.Itoa(res)
}

// PdfA sets the PDF/A format of the resulting PDF.
func (req *LibreOfficeRequest) PdfA(pdfa PdfAFormat) {
	req.fields[fieldOfficePdfA] = string(pdfa)
}

// PdfUA enables PDF for Universal Access for optimal accessibility.
func (req *LibreOfficeRequest) PdfUA() {
	req.fields[fieldOfficePdfUa] = strconv.FormatBool(true)
}

// Metadata sets the metadata to write.
func (req *LibreOfficeRequest) Metadata(md []byte) {
	req.fields[fieldMetadata] = string(md)
}

// Merge merges the resulting PDFs.
func (req *LibreOfficeRequest) Merge() {
	req.fields[fieldOfficeMerge] = strconv.FormatBool(true)
}

// Compile-time checks to ensure type implements desired interfaces.
var (
	_ = multipartRequester(new(LibreOfficeRequest))
)
