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

func TestLibreOffice(t *testing.T) {
	c, err := NewClient("http://localhost:3000", http.DefaultClient, nil)
	require.NoError(t, err)

	doc, err := document.FromPath("document.docx", testutil.LibreOfficeTestFilePath(t, "document.docx"))
	require.NoError(t, err)
	req := NewLibreOfficeRequest(doc)
	req.Trace("testLibreOffice")
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
	isPDFA, err := testutil.IsPDFA(dest)
	require.NoError(t, err)
	assert.False(t, isPDFA)
}

func TestLibreOfficePageRanges(t *testing.T) {
	c, err := NewClient("http://localhost:3000", http.DefaultClient, nil)
	require.NoError(t, err)

	doc, err := document.FromPath("document.docx", testutil.LibreOfficeTestFilePath(t, "document.docx"))
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
	c, err := NewClient("http://localhost:3000", http.DefaultClient, nil)
	require.NoError(t, err)

	doc, err := document.FromPath("document.docx", testutil.LibreOfficeTestFilePath(t, "document.docx"))
	require.NoError(t, err)
	req := NewLibreOfficeRequest(doc)
	req.Trace("testLibreOfficeLosslessCompression")
	req.UseBasicAuth("foo", "bar")
	req.OutputFilename("foo.pdf")
	req.LosslessImageCompression()
	dirPath := t.TempDir()
	dest := fmt.Sprintf("%s/foo.pdf", dirPath)
	err = c.Save(context.Background(), req, dest)
	require.NoError(t, err)
	assert.FileExists(t, dest)
	isPDF, err := testutil.IsPDF(dest)
	require.NoError(t, err)
	assert.True(t, isPDF)
}

func TestLibreOfficeCompression(t *testing.T) {
	c, err := NewClient("http://localhost:3000", http.DefaultClient, nil)
	require.NoError(t, err)

	doc, err := document.FromPath("document.docx", testutil.LibreOfficeTestFilePath(t, "document.docx"))
	require.NoError(t, err)
	req := NewLibreOfficeRequest(doc)
	req.Trace("testLibreOfficeCompression")
	req.UseBasicAuth("foo", "bar")
	req.OutputFilename("foo.pdf")
	req.Quality(1)
	req.ReduceImageResolution()
	req.MaxImageResolution(75)
	dirPath := t.TempDir()
	dest := fmt.Sprintf("%s/foo.pdf", dirPath)
	err = c.Save(context.Background(), req, dest)
	require.NoError(t, err)
	assert.FileExists(t, dest)
	isPDF, err := testutil.IsPDF(dest)
	require.NoError(t, err)
	assert.True(t, isPDF)
}

func TestLibreOfficeMultipleWithoutMerge(t *testing.T) {
	c, err := NewClient("http://localhost:3000", http.DefaultClient, nil)
	require.NoError(t, err)

	doc1, err := document.FromPath("document1.docx", testutil.LibreOfficeTestFilePath(t, "document.docx"))
	require.NoError(t, err)
	doc2, err := document.FromPath("document2.docx", testutil.LibreOfficeTestFilePath(t, "document.docx"))
	require.NoError(t, err)
	req := NewLibreOfficeRequest(doc1, doc2)
	req.Trace("testLibreOfficeMultipleWithoutMerge")
	req.UseBasicAuth("foo", "bar")
	req.OutputFilename("foo.zip")
	dirPath := t.TempDir()
	dest := fmt.Sprintf("%s/foo.zip", dirPath)
	err = c.Save(context.Background(), req, dest)
	require.NoError(t, err)
	assert.FileExists(t, dest)

	count, isPDFs, err := testutil.IsPDFsInArchive(t, dest)
	require.NoError(t, err)
	assert.Equal(t, 2, count)
	assert.True(t, isPDFs)
}

func TestLibreOfficeMultipleWithMerge(t *testing.T) {
	c, err := NewClient("http://localhost:3000", http.DefaultClient, nil)
	require.NoError(t, err)

	doc1, err := document.FromPath("document1.docx", testutil.LibreOfficeTestFilePath(t, "document.docx"))
	require.NoError(t, err)
	doc2, err := document.FromPath("document2.docx", testutil.LibreOfficeTestFilePath(t, "document.docx"))
	require.NoError(t, err)
	req := NewLibreOfficeRequest(doc1, doc2)
	req.Trace("testLibreOfficeMultipleWithMerge")
	req.UseBasicAuth("foo", "bar")
	req.OutputFilename("foo.pdf")
	req.Merge()
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
	assert.Equal(t, 4, count)
}

func TestLibreOfficePdfA(t *testing.T) {
	c, err := NewClient("http://localhost:3000", http.DefaultClient, nil)
	require.NoError(t, err)

	doc, err := document.FromPath("document.docx", testutil.LibreOfficeTestFilePath(t, "document.docx"))
	require.NoError(t, err)
	req := NewLibreOfficeRequest(doc)
	req.Trace("testLibreOfficePdfA")
	req.UseBasicAuth("foo", "bar")
	req.OutputFilename("foo.pdf")
	req.PdfA(PdfA3b)
	dirPath := t.TempDir()
	dest := fmt.Sprintf("%s/foo.pdf", dirPath)
	err = c.Save(context.Background(), req, dest)
	require.NoError(t, err)
	assert.FileExists(t, dest)
	isPDFA, err := testutil.IsPDFA(dest)
	require.NoError(t, err)
	assert.True(t, isPDFA)
}

func TestLibreOfficePdfUA(t *testing.T) {
	c, err := NewClient("http://localhost:3000", http.DefaultClient, nil)
	require.NoError(t, err)

	doc, err := document.FromPath("document.docx", testutil.LibreOfficeTestFilePath(t, "document.docx"))
	require.NoError(t, err)
	req := NewLibreOfficeRequest(doc)
	req.Trace("testLibreOfficePdfUA")
	req.UseBasicAuth("foo", "bar")
	req.OutputFilename("foo.pdf")
	req.PdfUA()
	dirPath := t.TempDir()
	dest := fmt.Sprintf("%s/foo.pdf", dirPath)
	err = c.Save(context.Background(), req, dest)
	require.NoError(t, err)
	assert.FileExists(t, dest)
	isPDFUA, err := testutil.IsPDFUA(dest)
	require.NoError(t, err)
	assert.True(t, isPDFUA)
}
