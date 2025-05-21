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

func TestHTML(t *testing.T) {
	c, err := NewClient("http://localhost:3000", http.DefaultClient)
	require.NoError(t, err)

	index, err := document.FromPath("index.html", testutil.HTMLTestFilePath(t, "index.html"))
	require.NoError(t, err)
	req := NewHTMLRequest(index)
	req.Trace("testHTML")
	req.UseBasicAuth("foo", "bar")

	cks := []Cookie{{Name: "foo", Value: "bar", Domain: "mydomain.com"}}
	err = req.Cookies(cks)
	require.NoError(t, err)

	header, err := document.FromPath("header.html", testutil.HTMLTestFilePath(t, "header.html"))
	require.NoError(t, err)
	req.Header(header)
	footer, err := document.FromPath("footer.html", testutil.HTMLTestFilePath(t, "footer.html"))
	require.NoError(t, err)
	req.Footer(footer)
	font, err := document.FromPath("font.woff", testutil.HTMLTestFilePath(t, "font.woff"))
	require.NoError(t, err)
	img, err := document.FromPath("img.gif", testutil.HTMLTestFilePath(t, "img.gif"))
	require.NoError(t, err)
	style, err := document.FromPath("style.css", testutil.HTMLTestFilePath(t, "style.css"))
	require.NoError(t, err)
	req.Assets(font, img, style)
	req.OutputFilename("foo.pdf")
	req.WaitDelay(1 * time.Second)
	req.PaperSize(A4)
	req.Margins(NormalMargins)
	req.Scale(1.5)
	dirPath := t.TempDir()
	dest := fmt.Sprintf("%s/foo.pdf", dirPath)
	err = c.Store(context.Background(), req, dest)
	require.NoError(t, err)
	assert.FileExists(t, dest)
	isPDF, err := testutil.IsPDF(dest)
	require.NoError(t, err)
	assert.True(t, isPDF)
}

func TestHTMLPageRanges(t *testing.T) {
	c, err := NewClient("http://localhost:3000", nil)
	require.NoError(t, err)

	index, err := document.FromPath("index.html", testutil.HTMLTestFilePath(t, "index.html"))
	require.NoError(t, err)
	req := NewHTMLRequest(index)
	req.Trace("testHTMLPageRanges")
	req.UseBasicAuth("foo", "bar")

	cks := []Cookie{{Name: "foo", Value: "bar", Domain: "mydomain.com"}}
	err = req.Cookies(cks)
	require.NoError(t, err)

	req.NativePageRanges("1-1")
	resp, err := c.Send(context.Background(), req)
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestHTMLScreenshot(t *testing.T) {
	c, err := NewClient("http://localhost:3000", http.DefaultClient)
	require.NoError(t, err)

	index, err := document.FromPath("index.html", testutil.HTMLTestFilePath(t, "index.html"))
	require.NoError(t, err)
	req := NewHTMLRequest(index)
	req.Trace("testHTMLScreenshot")
	req.UseBasicAuth("foo", "bar")

	cks := []Cookie{{Name: "foo", Value: "bar", Domain: "mydomain.com"}}
	err = req.Cookies(cks)
	require.NoError(t, err)

	dirPath := t.TempDir()
	req.Format(JPEG)
	dest := fmt.Sprintf("%s/foo.jpeg", dirPath)
	err = c.StoreScreenshot(context.Background(), req, dest)
	require.NoError(t, err)
	assert.FileExists(t, dest)

	isValidJPEG, err := testutil.IsValidJPEG(dest)
	require.NoError(t, err)
	assert.True(t, isValidJPEG)
}

func TestHTMLPdfA(t *testing.T) {
	c, err := NewClient("http://localhost:3000", http.DefaultClient)
	require.NoError(t, err)

	index, err := document.FromPath("index.html", testutil.HTMLTestFilePath(t, "index.html"))
	require.NoError(t, err)
	req := NewHTMLRequest(index)
	req.Trace("testHTMLPdfA")
	req.UseBasicAuth("foo", "bar")

	cks := []Cookie{{Name: "foo", Value: "bar", Domain: "mydomain.com"}}
	err = req.Cookies(cks)
	require.NoError(t, err)

	req.PdfA(PdfA3b)
	dirPath := t.TempDir()
	dest := fmt.Sprintf("%s/foo.pdf", dirPath)
	err = c.Store(context.Background(), req, dest)
	require.NoError(t, err)
	assert.FileExists(t, dest)
	isPDFA, err := testutil.IsPDFA(dest)
	require.NoError(t, err)
	assert.True(t, isPDFA)
}

func TestHTMLPdfUA(t *testing.T) {
	c, err := NewClient("http://localhost:3000", http.DefaultClient)
	require.NoError(t, err)

	index, err := document.FromPath("index.html", testutil.HTMLTestFilePath(t, "index.html"))
	require.NoError(t, err)
	req := NewHTMLRequest(index)
	req.Trace("testHTMLPdfUA")
	req.UseBasicAuth("foo", "bar")

	cks := []Cookie{{Name: "foo", Value: "bar", Domain: "mydomain.com"}}
	err = req.Cookies(cks)
	require.NoError(t, err)

	req.PdfUA()
	dirPath := t.TempDir()
	dest := fmt.Sprintf("%s/foo.pdf", dirPath)
	err = c.Store(context.Background(), req, dest)
	require.NoError(t, err)
	assert.FileExists(t, dest)
	isPDFUA, err := testutil.IsPDFUA(dest)
	require.NoError(t, err)
	assert.True(t, isPDFUA)
}
