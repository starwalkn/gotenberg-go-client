package gotenberg

import (
	"context"
	"fmt"
	"net/http"
)

type ScreenshotRequester interface {
	screenshotEndpoint() string

	baseRequester
}

func (c *Client) Screenshot(ctx context.Context, scr ScreenshotRequester) (*http.Response, error) {
	return c.screenshot(ctx, scr)
}

func (c *Client) screenshot(ctx context.Context, scr ScreenshotRequester) (*http.Response, error) {
	req, err := c.createRequest(ctx, scr, scr.screenshotEndpoint())
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, errSendRequestFailed
	}

	return resp, nil
}

func (c *Client) StoreScreenshot(ctx context.Context, req ScreenshotRequester, dest string) error {
	return c.storeScreenshot(ctx, req, dest)
}

func (c *Client) storeScreenshot(ctx context.Context, scr ScreenshotRequester, dest string) error {
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

	return c.writeNewFile(dest, resp.Body)
}
