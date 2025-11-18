package gotenberg

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/starwalkn/gotenberg-go-client/v9/document"
	"github.com/starwalkn/gotenberg-go-client/v9/test"
)

func TestSplitIntervals(t *testing.T) {
	c := NewClient("http://localhost:3000", http.DefaultClient)

	doc, err := document.FromPath("gotenberg1.pdf", test.PDFTestFilePath(t, "gotenberg.pdf"))
	require.NoError(t, err)

	dest := fmt.Sprintf("%s/splitted.zip", t.TempDir())

	var (
		span          = 1
		expectedCount = 3
	)

	err = c.PDFEngines().SplitIntervals(doc).
		Trace("testSplitIntervals").
		BasicAuth("foo", "bar").
		Span(span).
		OutputFilename("splitted.zip").
		Store(context.Background(), dest)

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

	dest := fmt.Sprintf("%s/splitted.pdf", t.TempDir())

	err = c.PDFEngines().SplitIntervals(doc).
		Trace("testSplitIntervalsOnePage").
		BasicAuth("foo", "bar").
		Span(3).
		OutputFilename("splitted.pdf").
		Store(context.Background(), dest)

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

	dest := fmt.Sprintf("%s/splitted.zip", t.TempDir())

	var (
		span          = "1-2"
		expectedCount = 2
	)

	err = c.PDFEngines().SplitPages(doc).
		Trace("testSplitPages").
		BasicAuth("foo", "bar").
		Span(span).
		Unify(false).
		OutputFilename("splitted.zip").
		Store(context.Background(), dest)

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

	dest := fmt.Sprintf("%s/splitted.pdf", t.TempDir())

	err = c.PDFEngines().SplitPages(doc).
		Trace("testSplitPagesOnePage").
		BasicAuth("foo", "bar").
		Span("1-1").
		Unify(false).
		OutputFilename("splitted.pdf").
		Store(context.Background(), dest)

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

	dest := fmt.Sprintf("%s/splitted.pdf", t.TempDir())

	err = c.PDFEngines().SplitPages(doc).
		Trace("testSplitPagesUnify").
		BasicAuth("foo", "bar").
		Span("1-2").
		Unify(true).
		OutputFilename("splitted.pdf").
		Store(context.Background(), dest)

	require.NoError(t, err)
	assert.FileExists(t, dest)

	isPDF, err := test.IsPDF(dest)
	require.NoError(t, err)
	require.True(t, isPDF)
}
