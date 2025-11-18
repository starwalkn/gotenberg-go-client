package gotenberg

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
)

var (
	errWebhookNotAllowed = errors.New("webhook is not allowed for request")
	errGenerationFailed  = errors.New("resulting file could not be generated")
	errSendRequestFailed = errors.New("request sending failed")
)

// MultipartRequest is a type for sending form fields and form files (documents) to the Gotenberg API.
type MultipartRequest interface {
	endpoint() string

	Request
}

// Client facilitates interacting with the Gotenberg API.
type Client struct {
	hostname   string
	httpClient *http.Client
}

// NewClient creates a new gotenberg.Client. If http.Client is passed as nil, then http.DefaultClient is used.
func NewClient(hostname string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	return &Client{
		hostname:   hostname,
		httpClient: httpClient,
	}
}

func (c *Client) Chromium() *ChromiumService {
	return &ChromiumService{c}
}

func (c *Client) LibreOffice() *LibreOfficeService {
	return &LibreOfficeService{c}
}

func (c *Client) PDFEngines() *PDFEnginesService {
	return &PDFEnginesService{c}
}

func (c *Client) send(ctx context.Context, req MultipartRequest) (*http.Response, error) {
	r, err := c.createRequest(ctx, req, req.endpoint())
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(r)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", errSendRequestFailed, err)
	}

	return resp, nil
}

func (c *Client) store(ctx context.Context, req MultipartRequest, dest string) error {
	if hasWebhook(req) {
		return errWebhookNotAllowed
	}

	resp, err := c.send(ctx, req)
	if err != nil {
		return err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%w: %d", errGenerationFailed, resp.StatusCode)
	}

	return writeNewFile(dest, resp.Body)
}

func writeNewFile(fpath string, in io.Reader) error {
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

func (c *Client) createRequest(ctx context.Context, mr MultipartRequest, endpoint string) (*http.Request, error) {
	body, contentType, err := multipartForm(mr)
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
		req.Header.Set(key, value)
	}

	return req, nil
}
