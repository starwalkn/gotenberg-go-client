package gotenberg

import (
	"context"
	"net/http"
	"strconv"

	"github.com/starwalkn/gotenberg-go-client/v8/document"
)

const endpointOfficeConvert = "/forms/libreoffice/convert"

// LibreOfficeRequest facilitates LibreOffice documents conversion with the Gotenberg API.
type LibreOfficeRequest struct {
	docs   []document.Document
	embeds []document.Document

	*baseRequest
	client *Client
}

func (r *LibreOfficeRequest) endpoint() string {
	return endpointOfficeConvert
}

func (r *LibreOfficeRequest) formDocuments() map[string]document.Document {
	files := make(map[string]document.Document)
	for _, doc := range r.docs {
		files[doc.Filename()] = doc
	}

	return files
}

func (r *LibreOfficeRequest) formEmbeds() map[string]document.Document {
	embeds := make(map[string]document.Document)

	for _, embed := range r.embeds {
		embeds[embed.Filename()] = embed
	}

	return embeds
}

func (r *LibreOfficeRequest) Embeds(docs ...document.Document) *LibreOfficeRequest {
	r.embeds = docs
	return r
}

// Password sets the password for opening the source file.
func (r *LibreOfficeRequest) Password(password string) *LibreOfficeRequest {
	r.fields[fieldOfficePassword] = password
	return r
}

// Landscape sets the paper orientation to landscape.
func (r *LibreOfficeRequest) Landscape() *LibreOfficeRequest {
	r.fields[fieldOfficeLandscape] = strconv.FormatBool(true)
	return r
}

// NativePageRanges sets the page ranges to print, e.g., "1-4". Empty means all pages.
func (r *LibreOfficeRequest) NativePageRanges(ranges string) *LibreOfficeRequest {
	r.fields[fieldOfficeNativePageRanges] = ranges
	return r
}

// ExportFormFields specifies whether form fields are exported as widgets
// or only their fixed print representation is exported.
func (r *LibreOfficeRequest) ExportFormFields(export bool) *LibreOfficeRequest {
	r.fields[fieldOfficeExportFormFields] = strconv.FormatBool(export)
	return r
}

// AllowDuplicateFieldNames specifies whether multiple form fields exported are allowed to have the same field name.
func (r *LibreOfficeRequest) AllowDuplicateFieldNames() *LibreOfficeRequest {
	r.fields[fieldOfficeAllowDuplicateFieldNames] = strconv.FormatBool(true)
	return r
}

// ExportBookmarks specifies if bookmarks are exported to PDF.
func (r *LibreOfficeRequest) ExportBookmarks(export bool) *LibreOfficeRequest {
	r.fields[fieldOfficeExportBookmarks] = strconv.FormatBool(export)
	return r
}

// ExportBookmarksToPdfDestination specifies that the bookmarks contained
// in the source LibreOffice file should be exported to the PDF file as Named Destination.
func (r *LibreOfficeRequest) ExportBookmarksToPdfDestination() *LibreOfficeRequest {
	r.fields[fieldOfficeExportBookmarksToPdfDestination] = strconv.FormatBool(true)
	return r
}

// ExportPlaceholders exports the placeholders fields visual markings only. The exported placeholder is ineffective.
func (r *LibreOfficeRequest) ExportPlaceholders() *LibreOfficeRequest {
	r.fields[fieldOfficeExportPlaceholders] = strconv.FormatBool(true)
	return r
}

// ExportNotes specifies if notes are exported to PDF.
func (r *LibreOfficeRequest) ExportNotes() *LibreOfficeRequest {
	r.fields[fieldOfficeExportNotes] = strconv.FormatBool(true)
	return r
}

// ExportNotesPages specifies if notes pages are exported to PDF. Notes pages are available in Impress documents only.
func (r *LibreOfficeRequest) ExportNotesPages() *LibreOfficeRequest {
	r.fields[fieldOfficeExportNotesPages] = strconv.FormatBool(true)
	return r
}

// ExportOnlyNotesPages specifies, if the form field exportNotesPages
// is set to true, if only notes pages are exported to PDF.
func (r *LibreOfficeRequest) ExportOnlyNotesPages() *LibreOfficeRequest {
	r.fields[fieldOfficeExportOnlyNotesPages] = strconv.FormatBool(true)
	return r
}

// ExportNotesInMargin specifies if notes in margin are exported to PDF.
func (r *LibreOfficeRequest) ExportNotesInMargin() *LibreOfficeRequest {
	r.fields[fieldOfficeExportNotesInMargin] = strconv.FormatBool(true)
	return r
}

// ConvertOooTargetToPdfTarget specifies that the target documents with .od[tpgs] extension,
// will have that extension changed to .pdf when the link is exported to PDF. The source document remains untouched.
func (r *LibreOfficeRequest) ConvertOooTargetToPdfTarget() *LibreOfficeRequest {
	r.fields[fieldOfficeConvertOooTargetToPdfTarget] = strconv.FormatBool(true)
	return r
}

// ExportLinksRelativeFsys specifies that the file system related hyperlinks (file:// protocol) present
// in the document will be exported as relative to the source document location.
func (r *LibreOfficeRequest) ExportLinksRelativeFsys() *LibreOfficeRequest {
	r.fields[fieldOfficeExportLinksRelativeFsys] = strconv.FormatBool(true)
	return r
}

// ExportHiddenSlides exports, for LibreOffice Impress, slides that are not included in slide shows.
func (r *LibreOfficeRequest) ExportHiddenSlides() *LibreOfficeRequest {
	r.fields[fieldOfficeExportHiddenSlides] = strconv.FormatBool(true)
	return r
}

// SkipEmptyPages Specifies that automatically inserted empty pages are suppressed.
// This option is active only if storing Writer documents.
func (r *LibreOfficeRequest) SkipEmptyPages() *LibreOfficeRequest {
	r.fields[fieldOfficeSkipEmptyPages] = strconv.FormatBool(true)
	return r
}

// AddOriginalDocumentAsStream specifies that a stream is inserted to the PDF file
// which contains the original document for archiving purposes.
func (r *LibreOfficeRequest) AddOriginalDocumentAsStream() *LibreOfficeRequest {
	r.fields[fieldOfficeAddOriginalDocumentAsStream] = strconv.FormatBool(true)
	return r
}

// SinglePageSheets ignores each sheet’s paper size, print ranges and shown/hidden
// status and puts every sheet (even hidden sheets) on exactly one page.
func (r *LibreOfficeRequest) SinglePageSheets() *LibreOfficeRequest {
	r.fields[fieldOfficeSinglePageSheets] = strconv.FormatBool(true)
	return r
}

// LosslessImageCompression specifies if images are exported to PDF using a lossless compression
// format like PNG or compressed using the JPEG format.
func (r *LibreOfficeRequest) LosslessImageCompression() *LibreOfficeRequest {
	r.fields[fieldOfficeLosslessImageCompression] = strconv.FormatBool(true)
	return r
}

// Quality specifies the quality of the JPG export. A higher value produces a higher-quality image and a larger file.
// Between 1 and 100.
func (r *LibreOfficeRequest) Quality(quality int) *LibreOfficeRequest {
	r.fields[fieldOfficeQuality] = strconv.Itoa(quality)
	return r
}

// ReduceImageResolution Specifies if the resolution of each image is reduced to the resolution specified by the
// form field maxImageResolution.
func (r *LibreOfficeRequest) ReduceImageResolution() *LibreOfficeRequest {
	r.fields[fieldOfficeReduceImageResolution] = strconv.FormatBool(true)
	return r
}

// MaxImageResolution If the form field reduceImageResolution is set to true, tells if all images will be reduced
// to the given value in DPI. Possible values are: 75, 150, 300, 600 and 1200.
func (r *LibreOfficeRequest) MaxImageResolution(res int) *LibreOfficeRequest {
	r.fields[fieldOfficeMaxImageResolution] = strconv.Itoa(res)
	return r
}

// PdfA sets the PDF/A format of the resulting PDF.
func (r *LibreOfficeRequest) PdfA(pdfa PdfAFormat) *LibreOfficeRequest {
	r.fields[fieldOfficePdfA] = string(pdfa)
	return r
}

// PdfUA enables PDF for Universal Access for optimal accessibility.
func (r *LibreOfficeRequest) PdfUA() *LibreOfficeRequest {
	r.fields[fieldOfficePdfUa] = strconv.FormatBool(true)
	return r
}

// Metadata sets the metadata to write.
func (r *LibreOfficeRequest) Metadata(md []byte) *LibreOfficeRequest {
	r.fields[fieldMetadata] = string(md)
	return r
}

// SplitIntervals splits the resulting PDF by interval.
func (r *LibreOfficeRequest) SplitIntervals(span int) *LibreOfficeRequest {
	r.fields[fieldSplitMode] = splitModeIntervals
	r.fields[fieldSplitSpan] = strconv.Itoa(span)

	return r
}

// SplitPages splits the resulting PDF by pages. Span should have the format “1-5”.
func (r *LibreOfficeRequest) SplitPages(span string, unify bool) *LibreOfficeRequest {
	r.fields[fieldSplitMode] = splitModePages
	r.fields[fieldSplitSpan] = span
	r.fields[fieldSplitUnify] = strconv.FormatBool(unify)

	return r
}

// Flatten defines whether the resulting PDF should be flattened.
func (r *LibreOfficeRequest) Flatten(val bool) *LibreOfficeRequest {
	r.fields[fieldOfficeFlatten] = strconv.FormatBool(val)
	return r
}

// Merge merges the resulting PDFs.
func (r *LibreOfficeRequest) Merge() *LibreOfficeRequest {
	r.fields[fieldOfficeMerge] = strconv.FormatBool(true)
	return r
}

// UpdateIndexes specifies whether to update the indexes before conversion, keeping in
// mind that doing so might result in missing links in the final PDF.
func (r *LibreOfficeRequest) UpdateIndexes(value bool) *LibreOfficeRequest {
	r.fields[fieldOfficeUpdateIndexes] = strconv.FormatBool(value)
	return r
}

func (r *LibreOfficeRequest) Encrypt(userPassword, ownerPassword string) *LibreOfficeRequest {
	r.fields[fieldUserPassword] = userPassword
	r.fields[fieldOwnerPassword] = ownerPassword

	return r
}

func (r *LibreOfficeRequest) OutputFilename(filename string) *LibreOfficeRequest {
	r.baseRequest.OutputFilename(filename)
	return r
}

func (r *LibreOfficeRequest) Trace(trace string) *LibreOfficeRequest {
	r.baseRequest.Trace(trace)
	return r
}

func (r *LibreOfficeRequest) BasicAuth(username, password string) *LibreOfficeRequest {
	r.baseRequest.BasicAuth(username, password)
	return r
}

func (r *LibreOfficeRequest) UseWebhook(hookURL, errorURL string) *LibreOfficeRequest {
	r.baseRequest.UseWebhook(hookURL, errorURL)
	return r
}

func (r *LibreOfficeRequest) SetWebhookMethod(method string) *LibreOfficeRequest {
	r.baseRequest.SetWebhookMethod(method)
	return r
}

func (r *LibreOfficeRequest) SetWebhookErrorMethod(method string) *LibreOfficeRequest {
	r.baseRequest.SetWebhookErrorMethod(method)
	return r
}

func (r *LibreOfficeRequest) SetWebhookExtraHeaders(headers map[string]string) *LibreOfficeRequest {
	r.baseRequest.SetWebhookExtraHeaders(headers)
	return r
}

func (r *LibreOfficeRequest) Store(ctx context.Context, path string) error {
	return r.client.store(ctx, r, path)
}

func (r *LibreOfficeRequest) Send(ctx context.Context) (*http.Response, error) {
	return r.client.send(ctx, r)
}

// Compile-time checks to ensure type implements desired interfaces.
var (
	_ = MultipartRequest(new(LibreOfficeRequest))
)
