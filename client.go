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
	errEmptyHostname     = errors.New("empty hostname")
	errWebhookNotAllowed = errors.New("webhook is not allowed for request")
	errGenerationFailed  = errors.New("resulting file could not be generated")
	errSendRequestFailed = errors.New("request sending failed")
)

// multipartRequester is a type for sending form fields and form files (documents) to the Gotenberg API.
type multipartRequester interface {
	endpoint() string

	baseRequester
}

// Client facilitates interacting with the Gotenberg API.
type Client struct {
	hostname   string
	httpClient *http.Client
}

// NewClient creates a new gotenberg.Client. If http.Client is passed as nil, then http.DefaultClient is used.
func NewClient(hostname string, httpClient *http.Client) (*Client, error) {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	if hostname == "" {
		return nil, errEmptyHostname
	}

	return &Client{
		hostname:   hostname,
		httpClient: httpClient,
	}, nil
}

// Send sends a request to the Gotenberg API and returns the response.
func (c *Client) Send(ctx context.Context, req multipartRequester) (*http.Response, error) {
	return c.send(ctx, req)
}

func (c *Client) send(ctx context.Context, r multipartRequester) (*http.Response, error) {
	req, err := c.createRequest(ctx, r, r.endpoint())
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", errSendRequestFailed, err)
	}

	return resp, nil
}

// Store creates the resulting file to given destination.
func (c *Client) Store(ctx context.Context, req multipartRequester, dest string) error {
	return c.store(ctx, req, dest)
}

func (c *Client) store(ctx context.Context, req multipartRequester, dest string) error {
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

func (c *Client) createRequest(ctx context.Context, mr multipartRequester, endpoint string) (*http.Request, error) {
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
		req.Header.Set(string(key), value)
	}

	return req, nil
}
