package gotenberg

import (
	"archive/zip"
	"context"
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/dcaraxes/gotenberg-go-client/document"
	"github.com/dcaraxes/gotenberg-go-client/test"
)

func TestOffice(t *testing.T) {
	c, err := NewClient("http://localhost:3000", &http.Client{})
	require.NoError(t, err)

	doc, err := document.FromPath("document.docx", test.OfficeTestFilePath(t, "document.docx"))
	require.NoError(t, err)
	req := NewOfficeRequest(doc)
	req.UseBasicAuth("foo", "bar")
	req.OutputFilename("foo.pdf")
	dirPath, err := test.Rand()
	require.NoError(t, err)
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
	err = os.RemoveAll(dirPath)
	require.NoError(t, err)
}

func TestOfficePageRanges(t *testing.T) {
	c, err := NewClient("http://localhost:3000", &http.Client{})
	require.NoError(t, err)

	doc, err := document.FromPath("document.docx", test.OfficeTestFilePath(t, "document.docx"))
	require.NoError(t, err)
	req := NewOfficeRequest(doc)
	req.UseBasicAuth("foo", "bar")
	req.NativePageRanges("1-1")
	resp, err := c.Send(context.Background(), req)
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestOfficeLosslessCompression(t *testing.T) {
	c, err := NewClient("http://localhost:3000", &http.Client{})
	require.NoError(t, err)

	doc, err := document.FromPath("document.docx", test.OfficeTestFilePath(t, "document.docx"))
	require.NoError(t, err)
	req := NewOfficeRequest(doc)
	req.UseBasicAuth("foo", "bar")
	req.OutputFilename("foo.pdf")
	req.LosslessImageCompression()
	dirPath, err := test.Rand()
	require.NoError(t, err)
	dest := fmt.Sprintf("%s/foo.pdf", dirPath)
	err = c.Store(context.Background(), req, dest)
	require.NoError(t, err)
	assert.FileExists(t, dest)
	isPDF, err := test.IsPDF(dest)
	require.NoError(t, err)
	assert.True(t, isPDF)
	err = os.RemoveAll(dirPath)
	require.NoError(t, err)
}

func TestOfficeCompression(t *testing.T) {
	c, err := NewClient("http://localhost:3000", &http.Client{})
	require.NoError(t, err)

	doc, err := document.FromPath("document.docx", test.OfficeTestFilePath(t, "document.docx"))
	require.NoError(t, err)
	req := NewOfficeRequest(doc)
	req.UseBasicAuth("foo", "bar")
	req.OutputFilename("foo.pdf")
	req.Quality(1)
	req.ReduceImageResolution()
	req.MaxImageResolution(75)
	dirPath, err := test.Rand()
	require.NoError(t, err)
	dest := fmt.Sprintf("%s/foo.pdf", dirPath)
	err = c.Store(context.Background(), req, dest)
	require.NoError(t, err)
	assert.FileExists(t, dest)
	isPDF, err := test.IsPDF(dest)
	require.NoError(t, err)
	assert.True(t, isPDF)
	err = os.RemoveAll(dirPath)
	require.NoError(t, err)
}

func TestOfficeMultipleWithoutMerge(t *testing.T) {
	c, err := NewClient("http://localhost:3000", &http.Client{})
	require.NoError(t, err)

	doc1, err := document.FromPath("document1.docx", test.OfficeTestFilePath(t, "document.docx"))
	require.NoError(t, err)
	doc2, err := document.FromPath("document2.docx", test.OfficeTestFilePath(t, "document.docx"))
	require.NoError(t, err)
	req := NewOfficeRequest(doc1, doc2)
	req.UseBasicAuth("foo", "bar")
	req.OutputFilename("foo.zip")
	dirPath, err := test.Rand()
	require.NoError(t, err)
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
	err = os.RemoveAll(dirPath)
	require.NoError(t, err)
}

func TestOfficeMultipleWithMerge(t *testing.T) {
	c, err := NewClient("http://localhost:3000", &http.Client{})
	require.NoError(t, err)

	doc1, err := document.FromPath("document1.docx", test.OfficeTestFilePath(t, "document.docx"))
	require.NoError(t, err)
	doc2, err := document.FromPath("document2.docx", test.OfficeTestFilePath(t, "document.docx"))
	require.NoError(t, err)
	req := NewOfficeRequest(doc1, doc2)
	req.UseBasicAuth("foo", "bar")
	req.OutputFilename("foo.pdf")
	req.Merge()
	dirPath, err := test.Rand()
	require.NoError(t, err)
	dest := fmt.Sprintf("%s/foo.pdf", dirPath)
	err = c.Store(context.Background(), req, dest)
	require.NoError(t, err)
	assert.FileExists(t, dest)
	isPDF, err := test.IsPDF(dest)
	require.NoError(t, err)
	assert.True(t, isPDF)
	err = os.RemoveAll(dirPath)
	require.NoError(t, err)
}

func TestOfficePdfA(t *testing.T) {
	c, err := NewClient("http://localhost:3000", &http.Client{})
	require.NoError(t, err)

	doc, err := document.FromPath("document.docx", test.OfficeTestFilePath(t, "document.docx"))
	require.NoError(t, err)
	req := NewOfficeRequest(doc)
	req.UseBasicAuth("foo", "bar")
	req.OutputFilename("foo.pdf")
	req.PdfA(PdfA3b)
	dirPath, err := test.Rand()
	require.NoError(t, err)
	dest := fmt.Sprintf("%s/foo.pdf", dirPath)
	err = c.Store(context.Background(), req, dest)
	require.NoError(t, err)
	assert.FileExists(t, dest)
	isPDFA, err := test.IsPDFA(dest)
	require.NoError(t, err)
	assert.True(t, isPDFA)
	err = os.RemoveAll(dirPath)
	require.NoError(t, err)
}

func TestOfficePdfUA(t *testing.T) {
	c, err := NewClient("http://localhost:3000", &http.Client{})
	require.NoError(t, err)

	doc, err := document.FromPath("document.docx", test.OfficeTestFilePath(t, "document.docx"))
	require.NoError(t, err)
	req := NewOfficeRequest(doc)
	req.UseBasicAuth("foo", "bar")
	req.OutputFilename("foo.pdf")
	req.PdfUA()
	dirPath, err := test.Rand()
	require.NoError(t, err)
	dest := fmt.Sprintf("%s/foo.pdf", dirPath)
	err = c.Store(context.Background(), req, dest)
	require.NoError(t, err)
	assert.FileExists(t, dest)
	isPDFUA, err := test.IsPDFUA(dest)
	require.NoError(t, err)
	assert.True(t, isPDFUA)
	err = os.RemoveAll(dirPath)
	require.NoError(t, err)
}