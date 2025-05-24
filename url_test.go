package gotenberg

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/starwalkn/gotenberg-go-client/v8/document"
	"github.com/starwalkn/gotenberg-go-client/v8/testutil"
)

func TestURL(t *testing.T) {
	c, err := NewClient("http://localhost:3000", http.DefaultClient, nil)
	require.NoError(t, err)

	req := NewURLRequest("http://example.com")
	req.Trace("testURL")
	req.UseBasicAuth("foo", "bar")
	dirPath := t.TempDir()
	dest := fmt.Sprintf("%s/foo.pdf", dirPath)
	err = c.Save(context.Background(), req, dest)
	require.NoError(t, err)
	assert.FileExists(t, dest)

	isPDF, err := testutil.IsPDF(dest)
	require.NoError(t, err)
	assert.True(t, isPDF)
}

func TestURLComplete(t *testing.T) {
	c, err := NewClient("http://localhost:3000", http.DefaultClient, nil)
	require.NoError(t, err)

	req := NewURLRequest("http://example.com")
	req.Trace("testURLComplete")
	req.UseBasicAuth("foo", "bar")
	header, err := document.FromPath("header.html", testutil.HTMLTestFilePath(t, "header.html"))
	require.NoError(t, err)
	req.Header(header)
	footer, err := document.FromPath("footer.html", testutil.HTMLTestFilePath(t, "footer.html"))
	require.NoError(t, err)
	req.Footer(footer)
	req.OutputFilename("foo.pdf")
	req.WaitDelay(1 * time.Second)
	req.PaperSize(A4)
	req.Margins(NormalMargins)
	dirPath := t.TempDir()
	dest := fmt.Sprintf("%s/foo.pdf", dirPath)
	err = c.Save(context.Background(), req, dest)
	require.NoError(t, err)
	assert.FileExists(t, dest)

	isPDF, err := testutil.IsPDF(dest)
	require.NoError(t, err)
	assert.True(t, isPDF)
}

func TestURLPageRanges(t *testing.T) {
	c, err := NewClient("http://localhost:3000", http.DefaultClient, nil)
	require.NoError(t, err)

	req := NewURLRequest("http://example.com")
	req.Trace("testURLPageRanges")
	req.UseBasicAuth("foo", "bar")
	req.NativePageRanges("1-1")
	resp, err := c.Send(context.Background(), req)
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestURLScreenshot(t *testing.T) {
	c, err := NewClient("http://localhost:3000", http.DefaultClient, nil)
	require.NoError(t, err)

	req := NewURLRequest("https://example.com")
	req.Trace("testURLScreenshot")
	req.UseBasicAuth("foo", "bar")
	dirPath := t.TempDir()
	req.ScreenshotFormat(JPEG)
	dest := fmt.Sprintf("%s/foo.jpeg", dirPath)
	err = c.SaveScreenshot(context.Background(), req, dest)
	require.NoError(t, err)
	assert.FileExists(t, dest)

	isValidJPEG, err := testutil.IsValidJPEG(dest)
	require.NoError(t, err)
	assert.True(t, isValidJPEG)
}
