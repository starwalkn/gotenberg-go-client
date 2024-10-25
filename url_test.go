package gotenberg

import (
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/dcaraxes/gotenberg-go-client/document"
	"github.com/dcaraxes/gotenberg-go-client/test"
)

func TestURL(t *testing.T) {
	c, err := NewClient("http://localhost:3000", &http.Client{})
	require.NoError(t, err)
	req := NewURLRequest("http://example.com")
	req.UseBasicAuth("foo", "bar")
	dirPath, err := test.Rand()
	require.NoError(t, err)
	dest := fmt.Sprintf("%s/foo.pdf", dirPath)
	err = c.Store(req, dest)
	require.NoError(t, err)
	assert.FileExists(t, dest)
	err = os.RemoveAll(dirPath)
	require.NoError(t, err)
}

func TestURLComplete(t *testing.T) {
	c, err := NewClient("http://localhost:3000", &http.Client{})

	require.NoError(t, err)
	req := NewURLRequest("http://example.com")
	req.UseBasicAuth("foo", "bar")
	header, err := document.FromPath("header.html", test.HTMLTestFilePath(t, "header.html"))
	require.NoError(t, err)
	req.Header(header)
	footer, err := document.FromPath("footer.html", test.HTMLTestFilePath(t, "footer.html"))
	require.NoError(t, err)
	req.Footer(footer)
	req.OutputFilename("foo.pdf")
	req.WaitDelay(1 * time.Second)
	req.PaperSize(A4)
	req.Margins(NormalMargins)
	dirPath, err := test.Rand()
	require.NoError(t, err)
	dest := fmt.Sprintf("%s/foo.pdf", dirPath)
	err = c.Store(req, dest)
	require.NoError(t, err)
	assert.FileExists(t, dest)
	err = os.RemoveAll(dirPath)
	require.NoError(t, err)
}

func TestURLPageRanges(t *testing.T) {
	c, err := NewClient("http://localhost:3000", &http.Client{})

	require.NoError(t, err)
	req := NewURLRequest("http://example.com")
	req.UseBasicAuth("foo", "bar")
	req.NativePageRanges("1-1")
	resp, err := c.Post(req)
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestURLScreenshot(t *testing.T) {
	c, err := NewClient("http://localhost:3000", &http.Client{})

	require.NoError(t, err)
	req := NewURLRequest("https://example.com")
	req.UseBasicAuth("foo", "bar")
	dirPath, err := test.Rand()
	require.NoError(t, err)
	req.Format(JPEG)
	dest := fmt.Sprintf("%s/foo.jpeg", dirPath)
	err = c.StoreScreenshot(req, dest)
	require.NoError(t, err)
	assert.FileExists(t, dest)
	err = os.RemoveAll(dirPath)
	require.NoError(t, err)
}