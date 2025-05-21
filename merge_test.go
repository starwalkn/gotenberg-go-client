package gotenberg

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/starwalkn/gotenberg-go-client/v8/document"
	"github.com/starwalkn/gotenberg-go-client/v8/testutil"
)

func TestMerge(t *testing.T) {
	c, err := NewClient("http://localhost:3000", http.DefaultClient)
	require.NoError(t, err)

	pdf1, err := document.FromPath("gotenberg1.pdf", testutil.PDFTestFilePath(t, "gotenberg.pdf"))
	require.NoError(t, err)
	pdf2, err := document.FromPath("gotenberg2.pdf", testutil.PDFTestFilePath(t, "gotenberg.pdf"))
	require.NoError(t, err)
	req := NewMergeRequest(pdf1, pdf2)
	req.Trace("testMerge")
	req.UseBasicAuth("foo", "bar")
	req.OutputFilename("foo.pdf")
	dirPath := t.TempDir()
	dest := fmt.Sprintf("%s/foo.pdf", dirPath)
	err = c.Save(context.Background(), req, dest)
	require.NoError(t, err)
	assert.FileExists(t, dest)

	isPDF, err := testutil.IsPDF(dest)
	require.NoError(t, err)
	assert.True(t, isPDF)

	count, err := testutil.GetPDFPageCount(dest)
	require.NoError(t, err)
	assert.Equal(t, 6, count)
}
