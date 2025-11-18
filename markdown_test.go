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
	"github.com/starwalkn/gotenberg-go-client/v8/test"
)

func TestMarkdown(t *testing.T) {
	c := NewClient("http://localhost:3000", http.DefaultClient)

	index, err := document.FromPath("index.html", test.MarkdownTestFilePath(t, "index.html"))
	require.NoError(t, err)
	markdown1, err := document.FromPath("paragraph1.md", test.MarkdownTestFilePath(t, "paragraph1.md"))
	require.NoError(t, err)
	markdown2, err := document.FromPath("paragraph2.md", test.MarkdownTestFilePath(t, "paragraph2.md"))
	require.NoError(t, err)
	markdown3, err := document.FromPath("paragraph3.md", test.MarkdownTestFilePath(t, "paragraph3.md"))
	require.NoError(t, err)

	header, err := document.FromPath("header.html", test.MarkdownTestFilePath(t, "header.html"))
	require.NoError(t, err)
	footer, err := document.FromPath("footer.html", test.MarkdownTestFilePath(t, "footer.html"))
	require.NoError(t, err)

	font, err := document.FromPath("font.woff", test.MarkdownTestFilePath(t, "font.woff"))
	require.NoError(t, err)
	img, err := document.FromPath("img.gif", test.MarkdownTestFilePath(t, "img.gif"))
	require.NoError(t, err)
	style, err := document.FromPath("style.css", test.MarkdownTestFilePath(t, "style.css"))
	require.NoError(t, err)

	dest := fmt.Sprintf("%s/foo.pdf", t.TempDir())

	err = c.Chromium().Markdown(index, markdown1, markdown2, markdown3).
		Trace("testMarkdown").
		BasicAuth("foo", "bar").
		ExtraHTTPHeaders(map[string]string{
			"X-Header":        "Value",
			"X-Scoped-Header": `value;scope=https?:\\/\\/([a-zA-Z0-9-]+\\.)*domain\\.com\\/.*`,
		}).
		Header(header).
		Footer(footer).
		Assets(font, img, style).
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

	count, err := test.GetPDFPageCount(dest)
	require.NoError(t, err)
	assert.Equal(t, 2, count)
}

func TestMarkdownPageRanges(t *testing.T) {
	c := NewClient("http://localhost:3000", http.DefaultClient)

	index, err := document.FromPath("index.html", test.MarkdownTestFilePath(t, "index.html"))
	require.NoError(t, err)
	markdown1, err := document.FromPath("paragraph1.md", test.MarkdownTestFilePath(t, "paragraph1.md"))
	require.NoError(t, err)
	markdown2, err := document.FromPath("paragraph2.md", test.MarkdownTestFilePath(t, "paragraph2.md"))
	require.NoError(t, err)
	markdown3, err := document.FromPath("paragraph3.md", test.MarkdownTestFilePath(t, "paragraph3.md"))
	require.NoError(t, err)

	resp, err := c.Chromium().Markdown(index, markdown1, markdown2, markdown3).
		Trace("testMarkdownPageRanges").
		BasicAuth("foo", "bar").
		ExtraHTTPHeaders(map[string]string{
			"X-Header":        "Value",
			"X-Scoped-Header": `value;scope=https?:\\/\\/([a-zA-Z0-9-]+\\.)*domain\\.com\\/.*`,
		}).
		NativePageRanges("1-1").
		Send(context.Background())

	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestMarkdownScreenshot(t *testing.T) {
	c := NewClient("http://localhost:3000", http.DefaultClient)

	index, err := document.FromPath("index.html", test.MarkdownTestFilePath(t, "index.html"))
	require.NoError(t, err)
	markdown1, err := document.FromPath("paragraph1.md", test.MarkdownTestFilePath(t, "paragraph1.md"))
	require.NoError(t, err)
	markdown2, err := document.FromPath("paragraph2.md", test.MarkdownTestFilePath(t, "paragraph2.md"))
	require.NoError(t, err)
	markdown3, err := document.FromPath("paragraph3.md", test.MarkdownTestFilePath(t, "paragraph3.md"))
	require.NoError(t, err)

	dest := fmt.Sprintf("%s/foo.jpeg", t.TempDir())

	err = c.Chromium().Markdown(index, markdown1, markdown2, markdown3).
		Trace("testMarkdownScreenshot").
		BasicAuth("foo", "bar").
		ExtraHTTPHeaders(map[string]string{
			"X-Header":        "Value",
			"X-Scoped-Header": `value;scope=https?:\\/\\/([a-zA-Z0-9-]+\\.)*domain\\.com\\/.*`,
		}).
		Format(JPEG).
		StoreScreenshot(context.Background(), dest)

	require.NoError(t, err)
	assert.FileExists(t, dest)

	isValidJPEG, err := test.IsValidJPEG(dest)
	require.NoError(t, err)
	assert.True(t, isValidJPEG)
}
