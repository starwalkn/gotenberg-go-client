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

func TestLibreOffice(t *testing.T) {
	c := NewClient("http://localhost:3000", http.DefaultClient)

	doc, err := document.FromPath("document.docx", test.LibreOfficeTestFilePath(t, "document.docx"))
	require.NoError(t, err)

	dest := fmt.Sprintf("%s/foo.pdf", t.TempDir())

	err = c.LibreOffice().HTML(doc).
		Trace("testLibreOffice").
		BasicAuth("foo", "bar").
		OutputFilename("foo.pdf").
		Store(context.Background(), dest)

	require.NoError(t, err)
	assert.FileExists(t, dest)
	isPDF, err := test.IsPDF(dest)
	require.NoError(t, err)
	assert.True(t, isPDF)
	isPDFA, err := test.IsPDFA(dest)
	require.NoError(t, err)
	assert.False(t, isPDFA)
}

func TestLibreOfficePageRanges(t *testing.T) {
	c := NewClient("http://localhost:3000", http.DefaultClient)

	doc, err := document.FromPath("document.docx", test.LibreOfficeTestFilePath(t, "document.docx"))
	require.NoError(t, err)

	resp, err := c.LibreOffice().HTML(doc).
		Trace("testLibreOfficePageRanges").
		BasicAuth("foo", "bar").
		NativePageRanges("1-1").
		Send(context.Background())

	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestLibreOfficeLosslessCompression(t *testing.T) {
	c := NewClient("http://localhost:3000", http.DefaultClient)

	doc, err := document.FromPath("document.docx", test.LibreOfficeTestFilePath(t, "document.docx"))
	require.NoError(t, err)

	dest := fmt.Sprintf("%s/foo.pdf", t.TempDir())

	err = c.LibreOffice().HTML(doc).
		Trace("testLibreOfficeLosslessCompression").
		BasicAuth("foo", "bar").
		OutputFilename("foo.pdf").
		LosslessImageCompression().
		Store(context.Background(), dest)

	require.NoError(t, err)
	assert.FileExists(t, dest)
	isPDF, err := test.IsPDF(dest)
	require.NoError(t, err)
	assert.True(t, isPDF)
}

func TestLibreOfficeCompression(t *testing.T) {
	c := NewClient("http://localhost:3000", http.DefaultClient)

	doc, err := document.FromPath("document.docx", test.LibreOfficeTestFilePath(t, "document.docx"))
	require.NoError(t, err)

	dest := fmt.Sprintf("%s/foo.pdf", t.TempDir())

	err = c.LibreOffice().HTML(doc).
		Trace("testLibreOfficeCompression").
		BasicAuth("foo", "bar").
		OutputFilename("foo.pdf").
		Quality(1).
		ReduceImageResolution().
		MaxImageResolution(75).
		Store(context.Background(), dest)

	require.NoError(t, err)
	assert.FileExists(t, dest)
	isPDF, err := test.IsPDF(dest)
	require.NoError(t, err)
	assert.True(t, isPDF)
}

func TestLibreOfficeMultipleWithoutMerge(t *testing.T) {
	c := NewClient("http://localhost:3000", http.DefaultClient)

	doc1, err := document.FromPath("document1.docx", test.LibreOfficeTestFilePath(t, "document.docx"))
	require.NoError(t, err)
	doc2, err := document.FromPath("document2.docx", test.LibreOfficeTestFilePath(t, "document.docx"))
	require.NoError(t, err)

	dest := fmt.Sprintf("%s/foo.zip", t.TempDir())

	err = c.LibreOffice().HTML(doc1, doc2).
		Trace("testLibreOfficeMultipleWithoutMerge").
		BasicAuth("foo", "bar").
		OutputFilename("foo.zip").
		Store(context.Background(), dest)

	require.NoError(t, err)
	assert.FileExists(t, dest)

	count, isPDFs, err := test.IsPDFsInArchive(t, dest)
	require.NoError(t, err)
	assert.Equal(t, 2, count)
	assert.True(t, isPDFs)
}

func TestLibreOfficeMultipleWithMerge(t *testing.T) {
	c := NewClient("http://localhost:3000", http.DefaultClient)

	doc1, err := document.FromPath("document1.docx", test.LibreOfficeTestFilePath(t, "document.docx"))
	require.NoError(t, err)
	doc2, err := document.FromPath("document2.docx", test.LibreOfficeTestFilePath(t, "document.docx"))
	require.NoError(t, err)

	dest := fmt.Sprintf("%s/foo.pdf", t.TempDir())

	err = c.LibreOffice().HTML(doc1, doc2).
		Trace("testLibreOfficeMultipleWithMerge").
		BasicAuth("foo", "bar").
		OutputFilename("foo.pdf").
		Merge().
		Store(context.Background(), dest)

	require.NoError(t, err)
	assert.FileExists(t, dest)
	isPDF, err := test.IsPDF(dest)
	require.NoError(t, err)
	assert.True(t, isPDF)

	count, err := test.GetPDFPageCount(dest)
	require.NoError(t, err)
	assert.Equal(t, 4, count)
}

func TestLibreOfficePdfA(t *testing.T) {
	c := NewClient("http://localhost:3000", http.DefaultClient)

	doc, err := document.FromPath("document.docx", test.LibreOfficeTestFilePath(t, "document.docx"))
	require.NoError(t, err)

	dest := fmt.Sprintf("%s/foo.pdf", t.TempDir())

	err = c.LibreOffice().HTML(doc).
		Trace("testLibreOfficePdfA").
		BasicAuth("foo", "bar").
		OutputFilename("foo.pdf").
		PdfA(PdfA3b).
		Store(context.Background(), dest)

	require.NoError(t, err)
	assert.FileExists(t, dest)
	isPDFA, err := test.IsPDFA(dest)
	require.NoError(t, err)
	assert.True(t, isPDFA)
}

func TestLibreOfficePdfUA(t *testing.T) {
	c := NewClient("http://localhost:3000", http.DefaultClient)

	doc, err := document.FromPath("document.docx", test.LibreOfficeTestFilePath(t, "document.docx"))
	require.NoError(t, err)

	dest := fmt.Sprintf("%s/foo.pdf", t.TempDir())

	err = c.LibreOffice().HTML(doc).
		Trace("testLibreOfficePdfUA").
		BasicAuth("foo", "bar").
		OutputFilename("foo.pdf").
		PdfUA().
		Store(context.Background(), dest)

	require.NoError(t, err)
	assert.FileExists(t, dest)
	isPDFUA, err := test.IsPDFUA(dest)
	require.NoError(t, err)
	assert.True(t, isPDFUA)
}

func TestLibreOfficeEmbeds(t *testing.T) {
	c := NewClient("http://localhost:3000", http.DefaultClient)

	doc1, err := document.FromPath("document1.docx", test.LibreOfficeTestFilePath(t, "document.docx"))
	require.NoError(t, err)
	doc2, err := document.FromPath("document2.docx", test.LibreOfficeTestFilePath(t, "document.docx"))
	require.NoError(t, err)

	dest := fmt.Sprintf("%s/foo.pdf", t.TempDir())

	err = c.LibreOffice().HTML(doc1).
		Trace("testLibreOfficeEmbeds").
		BasicAuth("foo", "bar").
		OutputFilename("foo.pdf").
		Merge().
		Embeds(doc2).
		Store(context.Background(), dest)

	require.NoError(t, err)
	assert.FileExists(t, dest)

	hasEmbeds, err := test.HasEmbeds(dest)
	require.NoError(t, err)
	assert.True(t, hasEmbeds)
}
