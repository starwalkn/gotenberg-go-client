package gotenberg

import (
	"context"
	"fmt"
	"net/http"
)

type screenshotRequester interface {
	screenshotEndpoint() string

	multipartRequester
}

func (c *Client) Screenshot(ctx context.Context, scr screenshotRequester) (*http.Response, error) {
	return c.screenshot(ctx, scr)
}

func (c *Client) screenshot(ctx context.Context, scr screenshotRequester) (*http.Response, error) {
	req, err := c.createRequest(ctx, scr, scr.screenshotEndpoint())
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", errSendRequestFailed, err)
	}

	return resp, nil
}

func (c *Client) StoreScreenshot(ctx context.Context, req screenshotRequester, dest string) error {
	return c.storeScreenshot(ctx, req, dest)
}

func (c *Client) storeScreenshot(ctx context.Context, scr screenshotRequester, dest string) error {
	if hasWebhook(scr) {
		return errWebhookNotAllowed
	}

	resp, err := c.screenshot(ctx, scr)
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
