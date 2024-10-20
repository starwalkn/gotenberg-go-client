package gotenberg

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/dcaraxes/gotenberg-go-client/v8/test"
)

func TestHTML(t *testing.T) {
	c := &Client{Hostname: "http://localhost:3000"}
	index, err := NewDocumentFromPath("index.html", test.HTMLTestFilePath(t, "index.html"))
	require.Nil(t, err)
	req := NewHTMLRequest(index)
	req.SetBasicAuth("foo", "bar")
	dirPath, err := test.Rand()
	require.Nil(t, err)
	dest := fmt.Sprintf("%s/foo.pdf", dirPath)
	err = c.Store(req, dest)
	assert.Nil(t, err)
	assert.FileExists(t, dest)
	isPDF, err := test.IsPDF(dest)
	assert.Nil(t, err)
	assert.True(t, isPDF)
	err = os.RemoveAll(dirPath)
	assert.Nil(t, err)
}

func TestHTMLFromString(t *testing.T) {
	c := &Client{Hostname: "http://localhost:3000"}
	index, err := NewDocumentFromString("index.html", "<html>Foo</html>")
	require.Nil(t, err)
	req := NewHTMLRequest(index)
	req.SetBasicAuth("foo", "bar")
	dirPath, err := test.Rand()
	require.Nil(t, err)
	dest := fmt.Sprintf("%s/foo.pdf", dirPath)
	err = c.Store(req, dest)
	assert.Nil(t, err)
	assert.FileExists(t, dest)
	isPDF, err := test.IsPDF(dest)
	assert.Nil(t, err)
	assert.True(t, isPDF)
	err = os.RemoveAll(dirPath)
	assert.Nil(t, err)
}

func TestHTMLFromBytes(t *testing.T) {
	c := &Client{Hostname: "http://localhost:3000"}
	index, err := NewDocumentFromBytes("index.html", []byte("<html>Foo</html>"))
	require.Nil(t, err)
	req := NewHTMLRequest(index)
	req.SetBasicAuth("foo", "bar")
	dirPath, err := test.Rand()
	require.Nil(t, err)
	dest := fmt.Sprintf("%s/foo.pdf", dirPath)
	err = c.Store(req, dest)
	assert.Nil(t, err)
	assert.FileExists(t, dest)
	isPDF, err := test.IsPDF(dest)
	assert.Nil(t, err)
	assert.True(t, isPDF)
	err = os.RemoveAll(dirPath)
	assert.Nil(t, err)
}

func TestHTMLFromReader(t *testing.T) {
	c := &Client{Hostname: "http://localhost:3000"}
	r, err := os.Open(test.HTMLTestFilePath(t, "index.html"))
	index, err := NewDocumentFromReader("index.html", r)
	require.Nil(t, err)
	req := NewHTMLRequest(index)
	req.SetBasicAuth("foo", "bar")
	dirPath, err := test.Rand()
	require.Nil(t, err)
	dest := fmt.Sprintf("%s/foo.pdf", dirPath)
	err = c.Store(req, dest)
	assert.Nil(t, err)
	assert.FileExists(t, dest)
	isPDF, err := test.IsPDF(dest)
	assert.Nil(t, err)
	assert.True(t, isPDF)
	err = os.RemoveAll(dirPath)
	assert.Nil(t, err)
}

func TestHTMLComplete(t *testing.T) {
	c := &Client{Hostname: "http://localhost:3000"}
	index, err := NewDocumentFromPath("index.html", test.HTMLTestFilePath(t, "index.html"))
	require.Nil(t, err)
	req := NewHTMLRequest(index)
	req.SetBasicAuth("foo", "bar")
	header, err := NewDocumentFromPath("header.html", test.HTMLTestFilePath(t, "header.html"))
	require.Nil(t, err)
	req.Header(header)
	footer, err := NewDocumentFromPath("footer.html", test.HTMLTestFilePath(t, "footer.html"))
	require.Nil(t, err)
	req.Footer(footer)
	font, err := NewDocumentFromPath("font.woff", test.HTMLTestFilePath(t, "font.woff"))
	require.Nil(t, err)
	img, err := NewDocumentFromPath("img.gif", test.HTMLTestFilePath(t, "img.gif"))
	require.Nil(t, err)
	style, err := NewDocumentFromPath("style.css", test.HTMLTestFilePath(t, "style.css"))
	require.Nil(t, err)
	req.Assets(font, img, style)
	req.ResultFilename("foo.pdf")
	req.WaitTimeout(5)
	req.WaitDelay(1 * time.Second)
	req.PaperSize(A4)
	req.Margins(NormalMargins)
	req.Landscape(false)
	req.Scale(1.5)
	dirPath, err := test.Rand()
	require.Nil(t, err)
	dest := fmt.Sprintf("%s/foo.pdf", dirPath)
	err = c.Store(req, dest)
	assert.Nil(t, err)
	assert.FileExists(t, dest)
	isPDF, err := test.IsPDF(dest)
	assert.Nil(t, err)
	assert.True(t, isPDF)
	err = os.RemoveAll(dirPath)
	assert.Nil(t, err)
}

func TestHTMLPageRanges(t *testing.T) {
	c := &Client{Hostname: "http://localhost:3000"}
	index, err := NewDocumentFromPath("index.html", test.HTMLTestFilePath(t, "index.html"))
	require.Nil(t, err)
	req := NewHTMLRequest(index)
	req.SetBasicAuth("foo", "bar")
	req.PageRanges("1-1")
	resp, err := c.Post(req)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestHTMLWebhook(t *testing.T) {
	c := &Client{Hostname: "http://localhost:3000"}
	index, err := NewDocumentFromPath("index.html", test.HTMLTestFilePath(t, "index.html"))
	require.Nil(t, err)
	req := NewHTMLRequest(index)
	req.SetBasicAuth("foo", "bar")
	req.WebhookURL("https://google.com")
	req.WebhookURLTimeout(5.0)
	req.AddWebhookURLHTTPHeader("A-Header", "Foo")
	resp, err := c.Post(req)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestHTMLScreenshot(t *testing.T) {
	c := &Client{Hostname: "http://localhost:3000"}
	index, err := NewDocumentFromPath("index.html", test.HTMLTestFilePath(t, "index.html"))
	require.Nil(t, err)
	req := NewHTMLRequest(index)
	req.SetBasicAuth("foo", "bar")
	req.AddWebhookURLHTTPHeader("A-Header", "Foo")
	dirPath, err := test.Rand()
	require.Nil(t, err)
	req.Format(JPEG)
	dest := fmt.Sprintf("%s/foo.jpeg", dirPath)
	err = c.StoreScreenshot(req, dest)
	assert.Nil(t, err)
	assert.FileExists(t, dest)
	err = os.RemoveAll(dirPath)
	assert.Nil(t, err)
}

func TestHTMLPdfA(t *testing.T) {
	c := &Client{Hostname: "http://localhost:3000"}
	index, err := NewDocumentFromPath("index.html", test.HTMLTestFilePath(t, "index.html"))
	require.Nil(t, err)
	req := NewHTMLRequest(index)
	req.SetBasicAuth("foo", "bar")
	req.PdfA(PdfA3b)
	dirPath, err := test.Rand()
	require.Nil(t, err)
	dest := fmt.Sprintf("%s/foo.pdf", dirPath)
	err = c.Store(req, dest)
	assert.Nil(t, err)
	assert.FileExists(t, dest)
	isPDFA, err := test.IsPDFA(dest)
	assert.Nil(t, err)
	assert.True(t, isPDFA)
	err = os.RemoveAll(dirPath)
	assert.Nil(t, err)
}

func TestHTMLPdfUA(t *testing.T) {
	c := &Client{Hostname: "http://localhost:3000"}
	index, err := NewDocumentFromPath("index.html", test.HTMLTestFilePath(t, "index.html"))
	require.Nil(t, err)
	req := NewHTMLRequest(index)
	req.SetBasicAuth("foo", "bar")
	req.PdfUA()
	dirPath, err := test.Rand()
	require.Nil(t, err)
	dest := fmt.Sprintf("%s/foo.pdf", dirPath)
	err = c.Store(req, dest)
	assert.Nil(t, err)
	assert.FileExists(t, dest)
	isPDFUA, err := test.IsPDFUA(dest)
	assert.Nil(t, err)
	assert.True(t, isPDFUA)
	err = os.RemoveAll(dirPath)
	assert.Nil(t, err)
}
