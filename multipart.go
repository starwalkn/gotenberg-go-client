package gotenberg

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"

	"github.com/starwalkn/gotenberg-go-client/v8/document"
)

func multipartForm(mr MultipartRequester) (body *bytes.Buffer, contentType string, err error) {
	body = &bytes.Buffer{}

	writer := multipart.NewWriter(body)
	defer func() {
		if closeErr := writer.Close(); closeErr != nil {
			err = fmt.Errorf("error closing writer: %w", closeErr)
		}
	}()

	if err = addDocuments(writer, mr.formDocuments(), "files"); err != nil {
		return nil, "", err
	}

	if err = addDocuments(writer, mr.formEmbeds(), "embeds"); err != nil {
		return nil, "", err
	}

	if err = addFormFields(writer, mr.formFields()); err != nil {
		return nil, "", err
	}

	return body, writer.FormDataContentType(), nil
}

func addFormFields(writer *multipart.Writer, formFields map[formField]string) error {
	for name, value := range formFields {
		if err := writer.WriteField(string(name), value); err != nil {
			return fmt.Errorf("writing %s form field: %w", name, err)
		}
	}

	return nil
}

func addDocuments(writer *multipart.Writer, documents map[string]document.Document, fieldname string) error {
	for fname, doc := range documents {
		in, err := doc.Reader()
		if err != nil {
			return fmt.Errorf("getting %s reader: %w", fname, err)
		}

		part, err := writer.CreateFormFile(fieldname, fname)
		if err != nil {
			_ = in.Close()

			return fmt.Errorf("creating %s form file: %w", fname, err)
		}

		if _, err = io.Copy(part, in); err != nil {
			_ = in.Close()

			return fmt.Errorf("copying %s data: %w", fname, err)
		}
	}

	return nil
}
