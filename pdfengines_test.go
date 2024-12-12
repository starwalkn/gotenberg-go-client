package gotenberg

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/runatal/gotenberg-go-client/v8/document"
	"github.com/runatal/gotenberg-go-client/v8/test"
)

func TestMerge(t *testing.T) {
	c, err := NewClient("http://localhost:3000", &http.Client{})
	require.NoError(t, err)

	pdf1, err := document.FromPath("gotenberg1.pdf", test.PDFTestFilePath(t, "gotenberg.pdf"))
	require.NoError(t, err)
	pdf2, err := document.FromPath("gotenberg2.pdf", test.PDFTestFilePath(t, "gotenberg.pdf"))
	require.NoError(t, err)
	req := NewMergeRequest(pdf1, pdf2)
	req.Trace("testMerge")
	req.UseBasicAuth("foo", "bar")
	req.OutputFilename("foo.pdf")
	tempDir := t.TempDir()
	dest := fmt.Sprintf("%s/foo.pdf", tempDir)
	err = c.Store(context.Background(), req, dest)
	require.NoError(t, err)
	assert.FileExists(t, dest)
}
