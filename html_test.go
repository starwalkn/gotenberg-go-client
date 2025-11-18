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

type testDocs struct {
	htmlIndex  document.Document
	htmlHeader document.Document
	htmlFooter document.Document
	htmlAssets []document.Document
}

func getTestDocs(t *testing.T) *testDocs {
	index, err := document.FromPath("index.html", test.HTMLTestFilePath(t, "index.html"))
	require.NoError(t, err)

	header, err := document.FromPath("header.html", test.HTMLTestFilePath(t, "header.html"))
	require.NoError(t, err)

	footer, err := document.FromPath("footer.html", test.HTMLTestFilePath(t, "footer.html"))
	require.NoError(t, err)

	font, err := document.FromPath("font.woff", test.HTMLTestFilePath(t, "font.woff"))
	require.NoError(t, err)

	img, err := document.FromPath("img.gif", test.HTMLTestFilePath(t, "img.gif"))
	require.NoError(t, err)

	style, err := document.FromPath("style.css", test.HTMLTestFilePath(t, "style.css"))
	require.NoError(t, err)

	return &testDocs{
		htmlIndex:  index,
		htmlHeader: header,
		htmlFooter: footer,
		htmlAssets: []document.Document{font, img, style},
	}
}

func TestHTML(t *testing.T) {
	c := NewClient("http://localhost:3000", http.DefaultClient)

	docs := getTestDocs(t)

	cks := []Cookie{{Name: "foo", Value: "bar", Domain: "mydomain.com"}}
	dest := fmt.Sprintf("%s/foo.pdf", t.TempDir())

	err := c.Chromium().HTML(docs.htmlIndex).
		Trace("testHTML").
		BasicAuth("foo", "bar").
		Cookies(cks).
		Header(docs.htmlHeader).
		Footer(docs.htmlFooter).
		Assets(docs.htmlAssets...).
		OutputFilename("foo.pdf").
		WaitDelay(1*time.Second).
		PaperSize(A4).
		Margins(NormalMargins).
		Scale(1.5).
		Store(context.Background(), dest)

	require.NoError(t, err)
	assert.FileExists(t, dest)
	isPDF, err := test.IsPDF(dest)
	require.NoError(t, err)
	assert.True(t, isPDF)
}

func TestHTMLPageRanges(t *testing.T) {
	c := NewClient("http://localhost:3000", http.DefaultClient)

	docs := getTestDocs(t)
	cks := []Cookie{{Name: "foo", Value: "bar", Domain: "mydomain.com"}}

	resp, err := c.Chromium().HTML(docs.htmlIndex).
		Trace("testHTMLPageRanges").
		BasicAuth("foo", "bar").
		Cookies(cks).
		NativePageRanges("1-1").
		Send(context.Background())

	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestHTMLScreenshot(t *testing.T) {
	c := NewClient("http://localhost:3000", http.DefaultClient)

	docs := getTestDocs(t)
	dest := fmt.Sprintf("%s/foo.jpeg", t.TempDir())

	err := c.Chromium().HTML(docs.htmlIndex).
		Trace("testHTMLScreenshot").
		BasicAuth("foo", "bar").
		Format(JPEG).
		StoreScreenshot(context.Background(), dest)

	require.NoError(t, err)
	assert.FileExists(t, dest)

	isValidJPEG, err := test.IsValidJPEG(dest)
	require.NoError(t, err)
	assert.True(t, isValidJPEG)
}

func TestHTMLPdfA(t *testing.T) {
	c := NewClient("http://localhost:3000", http.DefaultClient)

	docs := getTestDocs(t)
	dest := fmt.Sprintf("%s/foo.pdf", t.TempDir())

	err := c.Chromium().HTML(docs.htmlIndex).
		Trace("testHTMLPdfA").
		BasicAuth("foo", "bar").
		PdfA(PdfA3b).
		Store(context.Background(), dest)

	require.NoError(t, err)
	assert.FileExists(t, dest)
	isPDFA, err := test.IsPDFA(dest)
	require.NoError(t, err)
	assert.True(t, isPDFA)
}

func TestHTMLPdfUA(t *testing.T) {
	c := NewClient("http://localhost:3000", http.DefaultClient)

	docs := getTestDocs(t)
	dest := fmt.Sprintf("%s/foo.pdf", t.TempDir())

	err := c.Chromium().HTML(docs.htmlIndex).
		Trace("testHTMLPdfUA").
		BasicAuth("foo", "bar").
		PdfUA().
		Store(context.Background(), dest)

	require.NoError(t, err)
	assert.FileExists(t, dest)
	isPDFUA, err := test.IsPDFUA(dest)
	require.NoError(t, err)
	assert.True(t, isPDFUA)
}

func TestHTMLEmbeds(t *testing.T) {
	c := NewClient("http://localhost:3000", http.DefaultClient)

	docs := getTestDocs(t)
	dest := fmt.Sprintf("%s/foo.html", t.TempDir())
	embeds := docs.htmlAssets

	err := c.Chromium().HTML(docs.htmlIndex).
		Trace("testHTMLEmbeds").
		BasicAuth("foo", "bar").
		Embeds(embeds...).
		Store(context.Background(), dest)

	require.NoError(t, err)
	assert.FileExists(t, dest)

	isPDF, err := test.IsPDF(dest)
	require.NoError(t, err)
	assert.True(t, isPDF)

	hasEmbeds, err := test.HasEmbeds(dest)
	require.NoError(t, err)
	assert.True(t, hasEmbeds)
}
