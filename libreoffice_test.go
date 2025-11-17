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
	c, err := NewClient("http://localhost:3000", http.DefaultClient)
	require.NoError(t, err)

	doc, err := document.FromPath("document.docx", test.LibreOfficeTestFilePath(t, "document.docx"))
	require.NoError(t, err)
	req := NewLibreOfficeRequest(doc)
	req.Trace("testLibreOffice")
	req.UseBasicAuth("foo", "bar")
	req.OutputFilename("foo.pdf")

	dest := fmt.Sprintf("%s/foo.pdf", t.TempDir())
	err = c.Store(context.Background(), req, dest)
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
	c, err := NewClient("http://localhost:3000", http.DefaultClient)
	require.NoError(t, err)

	doc, err := document.FromPath("document.docx", test.LibreOfficeTestFilePath(t, "document.docx"))
	require.NoError(t, err)
	req := NewLibreOfficeRequest(doc)
	req.Trace("testLibreOfficePageRanges")
	req.UseBasicAuth("foo", "bar")
	req.NativePageRanges("1-1")
	resp, err := c.Send(context.Background(), req)
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestLibreOfficeLosslessCompression(t *testing.T) {
	c, err := NewClient("http://localhost:3000", http.DefaultClient)
	require.NoError(t, err)

	doc, err := document.FromPath("document.docx", test.LibreOfficeTestFilePath(t, "document.docx"))
	require.NoError(t, err)
	req := NewLibreOfficeRequest(doc)
	req.Trace("testLibreOfficeLosslessCompression")
	req.UseBasicAuth("foo", "bar")
	req.OutputFilename("foo.pdf")
	req.LosslessImageCompression()

	dest := fmt.Sprintf("%s/foo.pdf", t.TempDir())
	err = c.Store(context.Background(), req, dest)
	require.NoError(t, err)
	assert.FileExists(t, dest)
	isPDF, err := test.IsPDF(dest)
	require.NoError(t, err)
	assert.True(t, isPDF)
}

func TestLibreOfficeCompression(t *testing.T) {
	c, err := NewClient("http://localhost:3000", http.DefaultClient)
	require.NoError(t, err)

	doc, err := document.FromPath("document.docx", test.LibreOfficeTestFilePath(t, "document.docx"))
	require.NoError(t, err)
	req := NewLibreOfficeRequest(doc)
	req.Trace("testLibreOfficeCompression")
	req.UseBasicAuth("foo", "bar")
	req.OutputFilename("foo.pdf")
	req.Quality(1)
	req.ReduceImageResolution()
	req.MaxImageResolution(75)

	dest := fmt.Sprintf("%s/foo.pdf", t.TempDir())
	err = c.Store(context.Background(), req, dest)
	require.NoError(t, err)
	assert.FileExists(t, dest)
	isPDF, err := test.IsPDF(dest)
	require.NoError(t, err)
	assert.True(t, isPDF)
}

func TestLibreOfficeMultipleWithoutMerge(t *testing.T) {
	c, err := NewClient("http://localhost:3000", http.DefaultClient)
	require.NoError(t, err)

	doc1, err := document.FromPath("document1.docx", test.LibreOfficeTestFilePath(t, "document.docx"))
	require.NoError(t, err)
	doc2, err := document.FromPath("document2.docx", test.LibreOfficeTestFilePath(t, "document.docx"))
	require.NoError(t, err)
	req := NewLibreOfficeRequest(doc1, doc2)
	req.Trace("testLibreOfficeMultipleWithoutMerge")
	req.UseBasicAuth("foo", "bar")
	req.OutputFilename("foo.zip")

	dest := fmt.Sprintf("%s/foo.zip", t.TempDir())
	err = c.Store(context.Background(), req, dest)
	require.NoError(t, err)
	assert.FileExists(t, dest)

	count, isPDFs, err := test.IsPDFsInArchive(t, dest)
	require.NoError(t, err)
	assert.Equal(t, 2, count)
	assert.True(t, isPDFs)
}

func TestLibreOfficeMultipleWithMerge(t *testing.T) {
	c, err := NewClient("http://localhost:3000", http.DefaultClient)
	require.NoError(t, err)

	doc1, err := document.FromPath("document1.docx", test.LibreOfficeTestFilePath(t, "document.docx"))
	require.NoError(t, err)
	doc2, err := document.FromPath("document2.docx", test.LibreOfficeTestFilePath(t, "document.docx"))
	require.NoError(t, err)
	req := NewLibreOfficeRequest(doc1, doc2)
	req.Trace("testLibreOfficeMultipleWithMerge")
	req.UseBasicAuth("foo", "bar")
	req.OutputFilename("foo.pdf")
	req.Merge()

	dest := fmt.Sprintf("%s/foo.pdf", t.TempDir())
	err = c.Store(context.Background(), req, dest)
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
	c, err := NewClient("http://localhost:3000", http.DefaultClient)
	require.NoError(t, err)

	doc, err := document.FromPath("document.docx", test.LibreOfficeTestFilePath(t, "document.docx"))
	require.NoError(t, err)
	req := NewLibreOfficeRequest(doc)
	req.Trace("testLibreOfficePdfA")
	req.UseBasicAuth("foo", "bar")
	req.OutputFilename("foo.pdf")
	req.PdfA(PdfA3b)

	dest := fmt.Sprintf("%s/foo.pdf", t.TempDir())
	err = c.Store(context.Background(), req, dest)
	require.NoError(t, err)
	assert.FileExists(t, dest)
	isPDFA, err := test.IsPDFA(dest)
	require.NoError(t, err)
	assert.True(t, isPDFA)
}

func TestLibreOfficePdfUA(t *testing.T) {
	c, err := NewClient("http://localhost:3000", http.DefaultClient)
	require.NoError(t, err)

	doc, err := document.FromPath("document.docx", test.LibreOfficeTestFilePath(t, "document.docx"))
	require.NoError(t, err)
	req := NewLibreOfficeRequest(doc)
	req.Trace("testLibreOfficePdfUA")
	req.UseBasicAuth("foo", "bar")
	req.OutputFilename("foo.pdf")
	req.PdfUA()

	dest := fmt.Sprintf("%s/foo.pdf", t.TempDir())
	err = c.Store(context.Background(), req, dest)
	require.NoError(t, err)
	assert.FileExists(t, dest)
	isPDFUA, err := test.IsPDFUA(dest)
	require.NoError(t, err)
	assert.True(t, isPDFUA)
}

func TestLibreOfficeEmbeds(t *testing.T) {
	c, err := NewClient("http://localhost:3000", http.DefaultClient)
	require.NoError(t, err)

	doc1, err := document.FromPath("document1.docx", test.LibreOfficeTestFilePath(t, "document.docx"))
	require.NoError(t, err)

	req := NewLibreOfficeRequest(doc1)
	req.Trace("testLibreOfficeEmbeds")
	req.UseBasicAuth("foo", "bar")
	req.OutputFilename("foo.pdf")
	req.Merge()

	doc2, err := document.FromPath("document2.docx", test.LibreOfficeTestFilePath(t, "document.docx"))
	require.NoError(t, err)

	req.Embeds(doc2)

	dest := fmt.Sprintf("%s/foo.pdf", t.TempDir())
	err = c.Store(context.Background(), req, dest)
	require.NoError(t, err)
	assert.FileExists(t, dest)

	hasEmbeds, err := test.HasEmbeds(dest)
	require.NoError(t, err)
	assert.True(t, hasEmbeds)
}
