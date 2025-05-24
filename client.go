package gotenberg

import (
	"context"
	"errors"
	"io"
	"net/http"
)

type screenshotRequest interface {
	screenshotEndpoint() string

	multipartRequest
}

// multipartRequest is a type for sending form fields and form files (documents) to the Gotenberg API.
type multipartRequest interface {
	endpoint() string

	baseRequester
}

// Client facilitates interacting with the Gotenberg API.
// It manages the hostname and the underlying HTTP client for making requests.
type Client struct {
	hostname string
	*http.Client
	logger Logger
}

// NewClient creates a new gotenberg.Client.
//
// The 'hostname' parameter is the base URL of your Gotenberg API instance (e.g., "http://localhost:3000").
// If 'httpClient' is passed as 'nil', then 'http.DefaultClient' will be used.
func NewClient(hostname string, httpClient *http.Client, logger Logger) (*Client, error) {
	if hostname == "" {
		return nil, errors.New("hostname is empty")
	}

	client := &Client{hostname: hostname, Client: httpClient}

	if httpClient == nil {
		client.Client = http.DefaultClient
	}

	if logger == nil {
		client.logger = NopLogger{}
	} else {
		client.logger = logger
	}

	return client, nil
}

// Send sends a request to the Gotenberg API and returns the raw HTTP response.
//
// This is the public entry point for making requests to the Gotenberg service.
// It wraps the internal send method, providing a clear and simple API for clients
// that need to handle the *http.Response directly.
func (c *Client) Send(ctx context.Context, req multipartRequest) (*http.Response, error) {
	return c.send(ctx, req)
}

// Save sends a request to the Gotenberg API and saves the resulting document
// to the specified local file path.
//
// This method is suitable for operations where the Gotenberg output is a single file
// (e.g., PDF, DOCX, etc.), and you want to store it directly on the local filesystem.
// Webhooks are not allowed when using this method, as Gotenberg would
// send the result elsewhere via the webhook mechanism instead of returning it.
func (c *Client) Save(ctx context.Context, r multipartRequest, dest string) error {
	if r.hasWebhook() {
		return errWebhookNotAllowed
	}

	return c.save(ctx, r, dest)
}

// Screenshot sends a screenshot request to the Gotenberg API and returns the raw HTTP response.
//
// This is the public entry point for capturing screenshots via the Gotenberg service.
// It wraps the internal screenshot method, providing a clear and simple API for clients
// that need to handle the *http.Response directly.
func (c *Client) Screenshot(ctx context.Context, r screenshotRequest) (*http.Response, error) {
	return c.screenshot(ctx, r)
}

// SaveScreenshot sends a screenshot request to the Gotenberg API and saves the resulting image
// to the specified local file path.
//
// This method is suitable for operations where the Gotenberg output is a single image file
// (e.g., PNG, JPEG) and you want to store it directly on the local filesystem.
// Webhooks are not allowed when using this method, as Gotenberg would
// send the result elsewhere via the webhook mechanism instead of returning it.
func (c *Client) SaveScreenshot(ctx context.Context, r screenshotRequest, dest string) error {
	if r.hasWebhook() {
		return errWebhookNotAllowed
	}

	return c.saveScreenshot(ctx, r, dest)
}

// Stream sends a request to the Gotenberg API and returns the raw HTTP response body
// as an io.ReadCloser.
//
// This method provides direct access to the response stream, allowing for
// efficient processing of large documents without saving them to disk first
// (e.g., streaming to a client, uploading to cloud storage).
// The caller is responsible for reading from and closing the io.ReadCloser
// to prevent resource leaks.
//
// Webhooks are not allowed when using this method, as the output would be
// handled by Gotenberg's webhook mechanism instead of streamed back to the client.
func (c *Client) Stream(ctx context.Context, r multipartRequest) (io.ReadCloser, error) {
	if r.hasWebhook() {
		return nil, errWebhookNotAllowed
	}

	return c.stream(ctx, r)
}
