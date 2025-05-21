package gotenberg

import (
	"context"
	"fmt"
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
}

// NewClient creates a new gotenberg.Client.
//
// The 'hostname' parameter is the base URL of your Gotenberg API instance (e.g., "http://localhost:3000").
// If 'httpClient' is passed as 'nil', then 'http.DefaultClient' will be used.
func NewClient(hostname string, httpClient *http.Client) (*Client, error) {
	if hostname == "" {
		return nil, fmt.Errorf("hostname is empty")
	}

	client := &Client{hostname: hostname, Client: httpClient}

	if httpClient == nil {
		client.Client = http.DefaultClient
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
