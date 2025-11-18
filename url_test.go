package gotenberg

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/starwalkn/gotenberg-go-client/v9/document"
	"github.com/starwalkn/gotenberg-go-client/v9/test"
)

func TestURL(t *testing.T) {
	c := NewClient("http://localhost:3000", http.DefaultClient)

	dest := fmt.Sprintf("%s/foo.pdf", t.TempDir())

	err := c.Chromium().URL("http://example.com").
		Trace("testURL").
		BasicAuth("foo", "bar").
		Store(context.Background(), dest)

	require.NoError(t, err)
	assert.FileExists(t, dest)

	isPDF, err := test.IsPDF(dest)
	require.NoError(t, err)
	assert.True(t, isPDF)
}

func TestURLComplete(t *testing.T) {
	c := NewClient("http://localhost:3000", http.DefaultClient)

	header, err := document.FromPath("header.html", test.HTMLTestFilePath(t, "header.html"))
	require.NoError(t, err)
	footer, err := document.FromPath("footer.html", test.HTMLTestFilePath(t, "footer.html"))
	require.NoError(t, err)

	dest := fmt.Sprintf("%s/foo.pdf", t.TempDir())

	err = c.Chromium().URL("http://example.com").
		Trace("testURLComplete").
		BasicAuth("foo", "bar").
		Header(header).
		Footer(footer).
		OutputFilename("foo.pdf").
		WaitDelay(1*time.Second).
		PaperSize(A4).
		Margins(NormalMargins).
		Store(context.Background(), dest)

	require.NoError(t, err)
	assert.FileExists(t, dest)

	isPDF, err := test.IsPDF(dest)
	require.NoError(t, err)
	assert.True(t, isPDF)
}

func TestURLPageRanges(t *testing.T) {
	c := NewClient("http://localhost:3000", http.DefaultClient)

	resp, err := c.Chromium().URL("http://example.com").
		Trace("testURLPageRanges").
		BasicAuth("foo", "bar").
		NativePageRanges("1-1").
		Send(context.Background())

	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestURLScreenshot(t *testing.T) {
	c := NewClient("http://localhost:3000", http.DefaultClient)

	dest := fmt.Sprintf("%s/foo.jpeg", t.TempDir())

	err := c.Chromium().URL("http://example.com").
		Trace("testURLScreenshot").
		BasicAuth("foo", "bar").
		Format(JPEG).
		StoreScreenshot(context.Background(), dest)

	require.NoError(t, err)
	assert.FileExists(t, dest)

	isValidJPEG, err := test.IsValidJPEG(dest)
	require.NoError(t, err)
	assert.True(t, isValidJPEG)
}
