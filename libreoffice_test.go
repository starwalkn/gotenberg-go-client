package gotenberg

import (
	"archive/zip"
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
	dirPath := t.TempDir()
	dest := fmt.Sprintf("%s/foo.pdf", dirPath)
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
	dirPath := t.TempDir()
	dest := fmt.Sprintf("%s/foo.pdf", dirPath)
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
	dirPath := t.TempDir()
	dest := fmt.Sprintf("%s/foo.pdf", dirPath)
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
	dirPath := t.TempDir()
	dest := fmt.Sprintf("%s/foo.zip", dirPath)
	err = c.Store(context.Background(), req, dest)
	require.NoError(t, err)
	assert.FileExists(t, dest)

	zipReader, err := zip.OpenReader(dest)
	require.NoError(t, err)

	expectedFiles := map[string]bool{
		"document1.docx.pdf": false,
		"document2.docx.pdf": false,
	}

	for _, file := range zipReader.File {
		if _, ok := expectedFiles[file.Name]; ok {
			expectedFiles[file.Name] = true
		}
	}

	for fileName, found := range expectedFiles {
		assert.True(t, found, "File %s not found in zip", fileName)
	}
	err = zipReader.Close()
	require.NoError(t, err)
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
	dirPath := t.TempDir()
	dest := fmt.Sprintf("%s/foo.pdf", dirPath)
	err = c.Store(context.Background(), req, dest)
	require.NoError(t, err)
	assert.FileExists(t, dest)
	isPDF, err := test.IsPDF(dest)
	require.NoError(t, err)
	assert.True(t, isPDF)
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
	dirPath := t.TempDir()
	dest := fmt.Sprintf("%s/foo.pdf", dirPath)
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
	dirPath := t.TempDir()
	dest := fmt.Sprintf("%s/foo.pdf", dirPath)
	err = c.Store(context.Background(), req, dest)
	require.NoError(t, err)
	assert.FileExists(t, dest)
	isPDFUA, err := test.IsPDFUA(dest)
	require.NoError(t, err)
	assert.True(t, isPDFUA)
}
