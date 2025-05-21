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

func TestFlatten(t *testing.T) {
	c, err := NewClient("http://localhost:3000", http.DefaultClient)
	require.NoError(t, err)

	doc, err := document.FromPath("gotenberg1.pdf", testutil.PDFTestFilePath(t, "gotenberg.pdf"))
	require.NoError(t, err)

	r := NewFlattenRequest(doc)
	r.Trace("testFlatten")
	r.UseBasicAuth("foo", "bar")

	dirPath := t.TempDir()
	dest := fmt.Sprintf("%s/foo.pdf", dirPath)

	err = c.Store(context.Background(), r, dest)
	require.NoError(t, err)
	assert.FileExists(t, dest)

	isPDF, err := testutil.IsPDF(dest)
	require.NoError(t, err)
	assert.True(t, isPDF)
}
