package gotenberg

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/runatal/gotenberg-go-client/v8/document"
	"github.com/runatal/gotenberg-go-client/v8/test"
)

func TestHTML(t *testing.T) {
	c, err := NewClient("http://localhost:3000", &http.Client{})
	require.NoError(t, err)

	index, err := document.FromPath("index.html", test.HTMLTestFilePath(t, "index.html"))
	require.NoError(t, err)
	req := NewHTMLRequest(index)
	req.Trace("testHTML")
	req.UseBasicAuth("foo", "bar")
	dirPath, err := test.Rand()
	require.NoError(t, err)
	dest := fmt.Sprintf("%s/foo.pdf", dirPath)
	err = c.Store(context.Background(), req, dest)
	require.NoError(t, err)
	assert.FileExists(t, dest)
	isPDF, err := test.IsPDF(dest)
	require.NoError(t, err)
	assert.True(t, isPDF)
	err = os.RemoveAll(dirPath)
	require.NoError(t, err)
}

func TestHTMLFromString(t *testing.T) {
	c, err := NewClient("http://localhost:3000", &http.Client{})
	require.NoError(t, err)

	index, err := document.FromString("index.html", "<html>Foo</html>")
	require.NoError(t, err)
	req := NewHTMLRequest(index)
	req.Trace("testHTMLFromString")
	req.UseBasicAuth("foo", "bar")
	dirPath, err := test.Rand()
	require.NoError(t, err)
	dest := fmt.Sprintf("%s/foo.pdf", dirPath)
	err = c.Store(context.Background(), req, dest)
	require.NoError(t, err)
	assert.FileExists(t, dest)
	isPDF, err := test.IsPDF(dest)
	require.NoError(t, err)
	assert.True(t, isPDF)
	err = os.RemoveAll(dirPath)
	require.NoError(t, err)
}

func TestHTMLFromBytes(t *testing.T) {
	c, err := NewClient("http://localhost:3000", &http.Client{})
	require.NoError(t, err)

	index, err := document.FromBytes("index.html", []byte("<html>Foo</html>"))
	require.NoError(t, err)
	req := NewHTMLRequest(index)
	req.Trace("testHTMLFromBytes")
	req.UseBasicAuth("foo", "bar")
	dirPath, err := test.Rand()
	require.NoError(t, err)
	dest := fmt.Sprintf("%s/foo.pdf", dirPath)
	err = c.Store(context.Background(), req, dest)
	require.NoError(t, err)
	assert.FileExists(t, dest)
	isPDF, err := test.IsPDF(dest)
	require.NoError(t, err)
	assert.True(t, isPDF)
	err = os.RemoveAll(dirPath)
	require.NoError(t, err)
}

func TestHTMLFromReader(t *testing.T) {
	c, err := NewClient("http://localhost:3000", &http.Client{})
	require.NoError(t, err)

	r, err := os.Open(test.HTMLTestFilePath(t, "index.html"))
	require.NoError(t, err)
	index, err := document.FromReader("index.html", r)
	require.NoError(t, err)
	req := NewHTMLRequest(index)
	req.Trace("testHTMLFromReader")
	req.UseBasicAuth("foo", "bar")
	dirPath, err := test.Rand()
	require.NoError(t, err)
	dest := fmt.Sprintf("%s/foo.pdf", dirPath)
	err = c.Store(context.Background(), req, dest)
	require.NoError(t, err)
	assert.FileExists(t, dest)
	isPDF, err := test.IsPDF(dest)
	require.NoError(t, err)
	assert.True(t, isPDF)
	err = os.RemoveAll(dirPath)
	require.NoError(t, err)
}

func TestHTMLComplete(t *testing.T) {
	c, err := NewClient("http://localhost:3000", &http.Client{})
	require.NoError(t, err)

	index, err := document.FromPath("index.html", test.HTMLTestFilePath(t, "index.html"))
	require.NoError(t, err)
	req := NewHTMLRequest(index)
	req.Trace("testHTMLComplete")
	req.UseBasicAuth("foo", "bar")
	header, err := document.FromPath("header.html", test.HTMLTestFilePath(t, "header.html"))
	require.NoError(t, err)
	req.Header(header)
	footer, err := document.FromPath("footer.html", test.HTMLTestFilePath(t, "footer.html"))
	require.NoError(t, err)
	req.Footer(footer)
	font, err := document.FromPath("font.woff", test.HTMLTestFilePath(t, "font.woff"))
	require.NoError(t, err)
	img, err := document.FromPath("img.gif", test.HTMLTestFilePath(t, "img.gif"))
	require.NoError(t, err)
	style, err := document.FromPath("style.css", test.HTMLTestFilePath(t, "style.css"))
	require.NoError(t, err)
	req.Assets(font, img, style)
	req.OutputFilename("foo.pdf")
	req.WaitDelay(1 * time.Second)
	req.PaperSize(A4)
	req.Margins(NormalMargins)
	req.Scale(1.5)
	dirPath, err := test.Rand()
	require.NoError(t, err)
	dest := fmt.Sprintf("%s/foo.pdf", dirPath)
	err = c.Store(context.Background(), req, dest)
	require.NoError(t, err)
	assert.FileExists(t, dest)
	isPDF, err := test.IsPDF(dest)
	require.NoError(t, err)
	assert.True(t, isPDF)
	err = os.RemoveAll(dirPath)
	require.NoError(t, err)
}

func TestHTMLPageRanges(t *testing.T) {
	c, err := NewClient("http://localhost:3000", nil)
	require.NoError(t, err)

	index, err := document.FromPath("index.html", test.HTMLTestFilePath(t, "index.html"))
	require.NoError(t, err)
	req := NewHTMLRequest(index)
	req.Trace("testHTMLPageRanges")
	req.UseBasicAuth("foo", "bar")
	req.NativePageRanges("1-1")
	resp, err := c.Send(context.Background(), req)
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestHTMLScreenshot(t *testing.T) {
	c, err := NewClient("http://localhost:3000", &http.Client{})
	require.NoError(t, err)

	index, err := document.FromPath("index.html", test.HTMLTestFilePath(t, "index.html"))
	require.NoError(t, err)
	req := NewHTMLRequest(index)
	req.Trace("testHTMLScreenshot")
	req.UseBasicAuth("foo", "bar")
	dirPath, err := test.Rand()
	require.NoError(t, err)
	req.Format(JPEG)
	dest := fmt.Sprintf("%s/foo.jpeg", dirPath)
	err = c.StoreScreenshot(context.Background(), req, dest)
	require.NoError(t, err)
	assert.FileExists(t, dest)
	err = os.RemoveAll(dirPath)
	require.NoError(t, err)
}

func TestHTMLPdfA(t *testing.T) {
	c, err := NewClient("http://localhost:3000", &http.Client{})
	require.NoError(t, err)

	index, err := document.FromPath("index.html", test.HTMLTestFilePath(t, "index.html"))
	require.NoError(t, err)
	req := NewHTMLRequest(index)
	req.Trace("testHTMLPdfA")
	req.UseBasicAuth("foo", "bar")
	req.PdfA(PdfA3b)
	dirPath, err := test.Rand()
	require.NoError(t, err)
	dest := fmt.Sprintf("%s/foo.pdf", dirPath)
	err = c.Store(context.Background(), req, dest)
	require.NoError(t, err)
	assert.FileExists(t, dest)
	isPDFA, err := test.IsPDFA(dest)
	require.NoError(t, err)
	assert.True(t, isPDFA)
	err = os.RemoveAll(dirPath)
	require.NoError(t, err)
}

func TestHTMLPdfUA(t *testing.T) {
	c, err := NewClient("http://localhost:3000", &http.Client{})
	require.NoError(t, err)

	index, err := document.FromPath("index.html", test.HTMLTestFilePath(t, "index.html"))
	require.NoError(t, err)
	req := NewHTMLRequest(index)
	req.Trace("testHTMLPdfUA")
	req.UseBasicAuth("foo", "bar")
	req.PdfUA()
	dirPath, err := test.Rand()
	require.NoError(t, err)
	dest := fmt.Sprintf("%s/foo.pdf", dirPath)
	err = c.Store(context.Background(), req, dest)
	require.NoError(t, err)
	assert.FileExists(t, dest)
	isPDFUA, err := test.IsPDFUA(dest)
	require.NoError(t, err)
	assert.True(t, isPDFUA)
	err = os.RemoveAll(dirPath)
	require.NoError(t, err)
}