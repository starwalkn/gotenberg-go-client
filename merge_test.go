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

func TestMerge(t *testing.T) {
	c := NewClient("http://localhost:3000", http.DefaultClient)

	pdf1, err := document.FromPath("gotenberg1.pdf", test.PDFTestFilePath(t, "gotenberg.pdf"))
	require.NoError(t, err)
	pdf2, err := document.FromPath("gotenberg2.pdf", test.PDFTestFilePath(t, "gotenberg.pdf"))
	require.NoError(t, err)

	dest := fmt.Sprintf("%s/foo.pdf", t.TempDir())

	err = c.PDFEngines().Merge(pdf1, pdf2).
		Trace("testMerge").
		BasicAuth("foo", "bar").
		OutputFilename("foo.pdf").
		Store(context.Background(), dest)

	require.NoError(t, err)
	assert.FileExists(t, dest)

	isPDF, err := test.IsPDF(dest)
	require.NoError(t, err)
	assert.True(t, isPDF)

	count, err := test.GetPDFPageCount(dest)
	require.NoError(t, err)
	assert.Equal(t, 6, count)
}
