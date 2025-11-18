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

func TestSplitIntervals(t *testing.T) {
	c := NewClient("http://localhost:3000", http.DefaultClient)

	doc, err := document.FromPath("gotenberg1.pdf", test.PDFTestFilePath(t, "gotenberg.pdf"))
	require.NoError(t, err)

	r := NewSplitIntervalsRequest(doc)
	r.Trace("testSplitIntervals")
	r.UseBasicAuth("foo", "bar")

	var (
		span          = 1
		expectedCount = 3
	)

	r.SplitSpan(span)
	r.OutputFilename("splitted.zip")

	dirPath := t.TempDir()
	dest := fmt.Sprintf("%s/splitted.zip", dirPath)
	err = c.Store(context.Background(), r, dest)
	require.NoError(t, err)
	assert.FileExists(t, dest)

	count, isPDFs, err := test.IsPDFsInArchive(t, dest)
	require.NoError(t, err)

	require.Equal(t, expectedCount, count)
	require.True(t, isPDFs)
}

func TestSplitIntervalsOnePage(t *testing.T) {
	c := NewClient("http://localhost:3000", http.DefaultClient)

	doc, err := document.FromPath("gotenberg1.pdf", test.PDFTestFilePath(t, "gotenberg.pdf"))
	require.NoError(t, err)

	r := NewSplitIntervalsRequest(doc)
	r.Trace("testSplitIntervalsOnePage")
	r.UseBasicAuth("foo", "bar")

	r.SplitSpan(3)
	r.OutputFilename("splitted.pdf")

	dirPath := t.TempDir()
	dest := fmt.Sprintf("%s/splitted.pdf", dirPath)
	err = c.Store(context.Background(), r, dest)
	require.NoError(t, err)
	assert.FileExists(t, dest)

	isPDF, err := test.IsPDF(dest)
	require.NoError(t, err)
	require.True(t, isPDF)
}

func TestSplitPages(t *testing.T) {
	c := NewClient("http://localhost:3000", http.DefaultClient)

	doc, err := document.FromPath("gotenberg1.pdf", test.PDFTestFilePath(t, "gotenberg.pdf"))
	require.NoError(t, err)

	r := NewSplitPagesRequest(doc)
	r.Trace("testSplitPages")
	r.UseBasicAuth("foo", "bar")

	var (
		span          = "1-2"
		expectedCount = 2
	)

	r.SplitSpan(span)
	r.SplitUnify(false)
	r.OutputFilename("splitted.zip")

	dirPath := t.TempDir()
	dest := fmt.Sprintf("%s/splitted.zip", dirPath)
	err = c.Store(context.Background(), r, dest)
	require.NoError(t, err)
	assert.FileExists(t, dest)

	count, isPDFs, err := test.IsPDFsInArchive(t, dest)
	require.NoError(t, err)

	require.Equal(t, expectedCount, count)
	require.True(t, isPDFs)
}

func TestSplitPagesOnePage(t *testing.T) {
	c := NewClient("http://localhost:3000", http.DefaultClient)

	doc, err := document.FromPath("gotenberg1.pdf", test.PDFTestFilePath(t, "gotenberg.pdf"))
	require.NoError(t, err)

	r := NewSplitPagesRequest(doc)
	r.Trace("testSplitPagesOnePage")
	r.UseBasicAuth("foo", "bar")

	r.SplitSpan("1-1")
	r.SplitUnify(false)
	r.OutputFilename("splitted.pdf")

	dirPath := t.TempDir()
	dest := fmt.Sprintf("%s/splitted.pdf", dirPath)
	err = c.Store(context.Background(), r, dest)
	require.NoError(t, err)
	assert.FileExists(t, dest)

	isPDF, err := test.IsPDF(dest)
	require.NoError(t, err)
	require.True(t, isPDF)
}

func TestSplitPagesUnify(t *testing.T) {
	c := NewClient("http://localhost:3000", http.DefaultClient)

	doc, err := document.FromPath("gotenberg1.pdf", test.PDFTestFilePath(t, "gotenberg.pdf"))
	require.NoError(t, err)

	r := NewSplitPagesRequest(doc)
	r.Trace("testSplitPagesUnify")
	r.UseBasicAuth("foo", "bar")

	r.SplitSpan("1-2")
	r.SplitUnify(true)
	r.OutputFilename("splitted.pdf")

	dirPath := t.TempDir()
	dest := fmt.Sprintf("%s/splitted.pdf", dirPath)
	err = c.Store(context.Background(), r, dest)
	require.NoError(t, err)
	assert.FileExists(t, dest)

	isPDF, err := test.IsPDF(dest)
	require.NoError(t, err)
	require.True(t, isPDF)
}
