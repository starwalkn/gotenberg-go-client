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

func TestSplitIntervals(t *testing.T) {
	c, err := NewClient("http://localhost:3000", http.DefaultClient)
	require.NoError(t, err)

	doc, err := document.FromPath("gotenberg1.pdf", testutil.PDFTestFilePath(t, "gotenberg.pdf"))
	require.NoError(t, err)

	r, err := NewSplitRequest(SplitModeIntervals, doc)
	require.NoError(t, err)

	r.Trace("testSplitIntervals")
	r.UseBasicAuth("foo", "bar")

	var (
		span          = 1
		expectedCount = 3
	)

	r.SpanIntervals(span)
	r.OutputFilename("splitted.zip")

	dirPath := t.TempDir()
	dest := fmt.Sprintf("%s/splitted.zip", dirPath)
	err = c.Save(context.Background(), r, dest)
	require.NoError(t, err)
	assert.FileExists(t, dest)

	count, isPDFs, err := testutil.IsPDFsInArchive(t, dest)
	require.NoError(t, err)

	require.Equal(t, expectedCount, count)
	require.True(t, isPDFs)
}

func TestSplitIntervalsOnePage(t *testing.T) {
	c, err := NewClient("http://localhost:3000", http.DefaultClient)
	require.NoError(t, err)

	doc, err := document.FromPath("gotenberg1.pdf", testutil.PDFTestFilePath(t, "gotenberg.pdf"))
	require.NoError(t, err)

	r, err := NewSplitRequest(SplitModeIntervals, doc)
	require.NoError(t, err)

	r.Trace("testSplitIntervalsOnePage")
	r.UseBasicAuth("foo", "bar")

	r.SpanIntervals(3)
	r.OutputFilename("splitted.pdf")

	dirPath := t.TempDir()
	dest := fmt.Sprintf("%s/splitted.pdf", dirPath)
	err = c.Save(context.Background(), r, dest)
	require.NoError(t, err)
	assert.FileExists(t, dest)

	isPDF, err := testutil.IsPDF(dest)
	require.NoError(t, err)
	require.True(t, isPDF)
}

func TestSplitPages(t *testing.T) {
	c, err := NewClient("http://localhost:3000", http.DefaultClient)
	require.NoError(t, err)

	doc, err := document.FromPath("gotenberg1.pdf", testutil.PDFTestFilePath(t, "gotenberg.pdf"))
	require.NoError(t, err)

	r, err := NewSplitRequest(SplitModePages, doc)
	require.NoError(t, err)

	r.Trace("testSplitPages")
	r.UseBasicAuth("foo", "bar")

	var (
		span          = "1-2"
		expectedCount = 2
	)

	r.SpanPages(span)
	r.SplitUnify(false)
	r.OutputFilename("splitted.zip")

	dirPath := t.TempDir()
	dest := fmt.Sprintf("%s/splitted.zip", dirPath)
	err = c.Save(context.Background(), r, dest)
	require.NoError(t, err)
	assert.FileExists(t, dest)

	count, isPDFs, err := testutil.IsPDFsInArchive(t, dest)
	require.NoError(t, err)

	require.Equal(t, expectedCount, count)
	require.True(t, isPDFs)
}

func TestSplitPagesOnePage(t *testing.T) {
	c, err := NewClient("http://localhost:3000", http.DefaultClient)
	require.NoError(t, err)

	doc, err := document.FromPath("gotenberg1.pdf", testutil.PDFTestFilePath(t, "gotenberg.pdf"))
	require.NoError(t, err)

	r, err := NewSplitRequest(SplitModePages, doc)
	require.NoError(t, err)

	r.Trace("testSplitPagesOnePage")
	r.UseBasicAuth("foo", "bar")

	r.SpanPages("1-1")
	r.SplitUnify(false)
	r.OutputFilename("splitted.pdf")

	dirPath := t.TempDir()
	dest := fmt.Sprintf("%s/splitted.pdf", dirPath)
	err = c.Save(context.Background(), r, dest)
	require.NoError(t, err)
	assert.FileExists(t, dest)

	isPDF, err := testutil.IsPDF(dest)
	require.NoError(t, err)
	require.True(t, isPDF)
}

func TestSplitPagesUnify(t *testing.T) {
	c, err := NewClient("http://localhost:3000", http.DefaultClient)
	require.NoError(t, err)

	doc, err := document.FromPath("gotenberg1.pdf", testutil.PDFTestFilePath(t, "gotenberg.pdf"))
	require.NoError(t, err)

	r, err := NewSplitRequest(SplitModePages, doc)
	require.NoError(t, err)

	r.Trace("testSplitPagesUnify")
	r.UseBasicAuth("foo", "bar")

	r.SpanPages("1-2")
	r.SplitUnify(true)
	r.OutputFilename("splitted.pdf")

	dirPath := t.TempDir()
	dest := fmt.Sprintf("%s/splitted.pdf", dirPath)
	err = c.Save(context.Background(), r, dest)
	require.NoError(t, err)
	assert.FileExists(t, dest)

	isPDF, err := testutil.IsPDF(dest)
	require.NoError(t, err)
	require.True(t, isPDF)
}

func TestSplitIncorrectMode(t *testing.T) {
	doc, err := document.FromPath("gotenberg1.pdf", testutil.PDFTestFilePath(t, "gotenberg.pdf"))
	require.NoError(t, err)

	_, err = NewSplitRequest("IncorrectMode", doc)
	require.Error(t, err)
}
