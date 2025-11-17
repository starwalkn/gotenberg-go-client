package gotenberg

import "github.com/starwalkn/gotenberg-go-client/v8/document"

type EncryptRequest struct {
	pdfs []document.Document

	*baseRequest
}

// NewEncryptRequest creates a new Encrypt request. However, ownerPassword is optional and can be empty.
func NewEncryptRequest(userPassword, ownerPassword string, pdfs ...document.Document) *EncryptRequest {
	br := newBaseRequest()
	br.fields[fieldUserPassword] = userPassword

	if ownerPassword != "" {
		br.fields[fieldOwnerPassword] = ownerPassword
	}

	return &EncryptRequest{
		pdfs:        pdfs,
		baseRequest: br,
	}
}

func (req *EncryptRequest) endpoint() string {
	return "/forms/pdfengines/encrypt"
}

func (req *EncryptRequest) formDocuments() map[string]document.Document {
	files := make(map[string]document.Document)

	for _, pdf := range req.pdfs {
		files[pdf.Filename()] = pdf
	}

	return files
}

// formEmbeds a stub for implementing the Request interface.
// In the future, we need to think about a more flexible client architecture.
func (req *EncryptRequest) formEmbeds() map[string]document.Document {
	return nil
}

// Compile-time checks to ensure type implements desired interfaces.
var (
	_ = MultipartRequest(new(EncryptRequest))
)
