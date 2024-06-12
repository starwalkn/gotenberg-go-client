package gotenberg

// MergeRequest facilitates work with PDF files
// with the Gotenberg API.
type MergeRequest struct {
	pdfs []Document

	*request
}

type ReadMetadataRequest struct {
	pdfs []Document

	*request
}

type WriteMetadataRequest struct {
	pdfs []Document

	*request
}

// NewMergeRequest create MergeRequest.
func NewMergeRequest(pdfs ...Document) *MergeRequest {
	return &MergeRequest{pdfs, newRequest()}
}

func NewReadMetadataRequest(pdfs ...Document) *ReadMetadataRequest {
	return &ReadMetadataRequest{pdfs, newRequest()}
}

func NewWriteMetadataRequest(pdfs ...Document) *WriteMetadataRequest {
	return &WriteMetadataRequest{pdfs, newRequest()}
}

func (req *MergeRequest) Metadata(jsonData []byte) {
	req.values[formFieldMetadata] = string(jsonData)
}

func (req *WriteMetadataRequest) Metadata(jsonData []byte) {
	req.values[formFieldMetadata] = string(jsonData)
}

func (req *MergeRequest) postURL() string {
	return "/forms/pdfengines/merge"
}

func (req *ReadMetadataRequest) postURL() string {
	return "/forms/pdfengines/metadata/read"
}

func (req *WriteMetadataRequest) postURL() string {
	return "/forms/pdfengines/metadata/write"
}

func (req *MergeRequest) formFiles() map[string]Document {
	files := make(map[string]Document)
	for _, pdf := range req.pdfs {
		files[pdf.Filename()] = pdf
	}
	return files
}

func (req *ReadMetadataRequest) formFiles() map[string]Document {
	files := make(map[string]Document)
	for _, pdf := range req.pdfs {
		files[pdf.Filename()] = pdf
	}
	return files
}

func (req *WriteMetadataRequest) formFiles() map[string]Document {
	files := make(map[string]Document)
	for _, pdf := range req.pdfs {
		files[pdf.Filename()] = pdf
	}
	return files
}

// Compile-time checks to ensure type implements desired interfaces.
var (
	_ = Request(new(MergeRequest))
)
