package gotenberg

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"runtime"

	"github.com/starwalkn/gotenberg-go-client/v8/document"
)

var (
	errWebhookNotAllowed = errors.New("webhook is not allowed for request")
	errGenerationFailed  = errors.New("resulting file could not be generated")
	errSendRequestFailed = errors.New("request sending failed")
)

// send is the internal method that prepares and executes the HTTP request to the Gotenberg API.
// It handles request creation, execution via the HTTP client, and basic error wrapping.
//
// This method is not intended for direct external use; `Send` should be used instead.
func (c *Client) send(ctx context.Context, r multipartRequest) (*http.Response, error) {
	req, err := c.createRequest(ctx, r, r.endpoint())
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", errSendRequestFailed, err)
	}

	return resp, nil
}

// save is the internal method that handles sending the request to Gotenberg
// and writing the response body to a local file.
//
// It sends the request via c.send, validates the HTTP response status,
// and then writes the content of the response body to the destination path.
//
// This method is not intended for direct external use; Save should be used instead.
func (c *Client) save(ctx context.Context, req multipartRequest, dest string) error {
	resp, err := c.send(ctx, req)
	if err != nil {
		return err
	}
	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil {
			c.logger.Printf("save: failed to close response body: %s", closeErr.Error())
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%w: %d", errGenerationFailed, resp.StatusCode)
	}

	return copyReaderToFile(dest, resp.Body)
}

func (c *Client) screenshot(ctx context.Context, scr screenshotRequest) (*http.Response, error) {
	req, err := c.createRequest(ctx, scr, scr.screenshotEndpoint())
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", errSendRequestFailed, err)
	}

	return resp, nil
}

func (c *Client) saveScreenshot(ctx context.Context, scr screenshotRequest, dest string) error {
	resp, err := c.screenshot(ctx, scr)
	if err != nil {
		return err
	}
	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil {
			c.logger.Printf("saveScreenshot: failed to close response body: %s", closeErr.Error())
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%w: %d", errGenerationFailed, resp.StatusCode)
	}

	return copyReaderToFile(dest, resp.Body)
}

// stream is the internal method that sends a request to the Gotenberg API
// and returns the raw HTTP response body as an io.ReadCloser.
//
// This method provides direct access to the response stream. The caller is responsible
// for reading from and closing the io.ReadCloser to prevent resource leaks.
// Webhooks are not allowed when using this method.
//
// This method is not intended for direct external use; public method Stream should be used instead.
func (c *Client) stream(ctx context.Context, req multipartRequest) (io.ReadCloser, error) { //nolint:unused
	resp, err := c.send(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request for streaming: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		if closeErr := resp.Body.Close(); closeErr != nil {
			c.logger.Printf("stream: failed to close response body: %s", closeErr.Error())
		}

		return nil, fmt.Errorf("%w: %d", errGenerationFailed, resp.StatusCode)
	}

	return resp.Body, nil
}

// copyReaderToFile creates a new file at 'fpath', ensures its directory exists,
// sets appropriate permissions, and copies all content from 'in' (an io.Reader) into it.
//
// This is an internal helper function used for saving streamed or generated content to disk.
// It handles directory creation, file creation, permissions, and ensures proper closing of the file.
func copyReaderToFile(fpath string, in io.Reader) error {
	if err := os.MkdirAll(filepath.Dir(fpath), 0o755); err != nil {
		return fmt.Errorf("making %s directory: %w", fpath, err)
	}

	out, err := os.Create(fpath)
	if err != nil {
		return fmt.Errorf("creating %s: %w", fpath, err)
	}
	defer func() {
		if closeErr := out.Close(); closeErr != nil && err == nil {
			err = fmt.Errorf("closing %s: %w", fpath, closeErr)
		}
	}()

	if err = out.Chmod(0o644); err != nil && runtime.GOOS != "windows" {
		return fmt.Errorf("setting %s permissions: %w", fpath, err)
	}

	if _, err = io.Copy(out, in); err != nil {
		return fmt.Errorf("writing to %s: %w", fpath, err)
	}

	return nil
}

// createRequest builds an *http.Request object for sending to the Gotenberg API.
// It prepares the multipart form body by adding documents and form fields,
// sets the appropriate content type, and adds any custom headers.
// This is an internal helper method used by send.
func (c *Client) createRequest(ctx context.Context, mr multipartRequest, endpoint string) (r *http.Request, err error) {
	body, contentType, err := c.multipartForm(mr)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s%s", c.hostname, endpoint)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, body)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	req.Header.Set("Content-Type", contentType)
	for key, value := range mr.customHeaders() {
		req.Header.Set(string(key), value)
	}

	return req, nil
}

func (c *Client) multipartForm(mr multipartRequest) (body *bytes.Buffer, contentType string, err error) {
	body = &bytes.Buffer{}

	writer := multipart.NewWriter(body)
	defer func() {
		if closeErr := writer.Close(); closeErr != nil {
			err = fmt.Errorf("error closing writer: %w", closeErr)
		}
	}()

	if err = c.addDocuments(writer, mr.formDocuments()); err != nil {
		return nil, "", err
	}

	if err = c.addFormFields(writer, mr.formFields()); err != nil {
		return nil, "", err
	}

	return body, writer.FormDataContentType(), nil
}

func (c *Client) addFormFields(writer *multipart.Writer, formFields map[formField]string) error {
	for name, value := range formFields {
		if err := writer.WriteField(string(name), value); err != nil {
			return fmt.Errorf("writing %s form field: %w", name, err)
		}
	}

	return nil
}

func (c *Client) addDocuments(writer *multipart.Writer, documents map[string]document.Document) error {
	for fname, doc := range documents {
		in, err := doc.Reader()
		if err != nil {
			return fmt.Errorf("getting %s reader: %w", fname, err)
		}

		part, err := writer.CreateFormFile("files", fname)
		if err != nil {
			if closeErr := in.Close(); closeErr != nil {
				c.logger.Printf("addDocuments: failed to close reader body: %s", closeErr.Error())
			}

			return fmt.Errorf("creating %s form file: %w", fname, err)
		}

		if _, err = io.Copy(part, in); err != nil {
			if closeErr := in.Close(); closeErr != nil {
				c.logger.Printf("addDocuments: failed to close reader body: %s", closeErr.Error())
			}

			return fmt.Errorf("copying %s data: %w", fname, err)
		}

		if closeErr := in.Close(); closeErr != nil {
			c.logger.Printf("addDocuments: failed to close reader body: %s", closeErr.Error())
		}
	}

	return nil
}
