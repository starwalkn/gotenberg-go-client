package gotenberg

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/starwalkn/gotenberg-go-client/v8/document"
	"github.com/starwalkn/gotenberg-go-client/v8/test"
)

func TestEmbed(t *testing.T) {
	c, err := NewClient("http://localhost:3000", http.DefaultClient)
	require.NoError(t, err)

	var docs []document.Document

	doc1, err := document.FromPath("gotenberg1.pdf", test.PDFTestFilePath(t, "gotenberg.pdf"))
	require.NoError(t, err)
	docs = append(docs, doc1)

	var embeds []document.Document

	doc2, err := document.FromPath("gotenberg2.pdf", test.PDFTestFilePath(t, "gotenberg_bis.pdf"))
	require.NoError(t, err)
	embeds = append(embeds, doc2)

	r := NewEmbedRequest([]document.Document{doc1}, embeds)
	r.Trace("testEmbed")
	r.UseBasicAuth("foo", "bar")

	dest := fmt.Sprintf("%s/foo.pdf", t.TempDir())
	err = c.Store(context.Background(), r, dest)
	require.NoError(t, err)
	assert.FileExists(t, dest)

	isPDF, err := test.IsPDF(dest)
	require.NoError(t, err)
	assert.True(t, isPDF)

	hasEmbeds, err := test.HasEmbeds(dest)
	require.NoError(t, err)
	assert.True(t, hasEmbeds)
}
