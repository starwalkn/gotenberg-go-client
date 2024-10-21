package gotenberg

import (
	"archive/zip"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/dcaraxes/gotenberg-go-client/v8/test"
)

func TestOffice(t *testing.T) {
	c := &Client{Hostname: "http://localhost:3000"}
	doc, err := NewDocumentFromPath("document.docx", test.OfficeTestFilePath(t, "document.docx"))
	require.Nil(t, err)
	req := NewOfficeRequest(doc)
	req.SetBasicAuth("foo", "bar")
	req.ResultFilename("foo.pdf")
	req.WaitTimeout(5)
	req.Landscape(false)
	dirPath, err := test.Rand()
	require.Nil(t, err)
	dest := fmt.Sprintf("%s/foo.pdf", dirPath)
	err = c.Store(req, dest)
	assert.Nil(t, err)
	assert.FileExists(t, dest)
	isPDF, err := test.IsPDF(dest)
	assert.Nil(t, err)
	assert.True(t, isPDF)
	isPDFA, err := test.IsPDFA(dest)
	assert.Nil(t, err)
	assert.False(t, isPDFA)
	err = os.RemoveAll(dirPath)
	assert.Nil(t, err)
}

func TestOfficePageRanges(t *testing.T) {
	c := &Client{Hostname: "http://localhost:3000"}
	doc, err := NewDocumentFromPath("document.docx", test.OfficeTestFilePath(t, "document.docx"))
	require.Nil(t, err)
	req := NewOfficeRequest(doc)
	req.SetBasicAuth("foo", "bar")
	req.PageRanges("1-1")
	resp, err := c.Post(req)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestOfficeLosslessCompression(t *testing.T) {
	c := &Client{Hostname: "http://localhost:3000"}
	doc, err := NewDocumentFromPath("document.docx", test.OfficeTestFilePath(t, "document.docx"))
	require.Nil(t, err)
	req := NewOfficeRequest(doc)
	req.SetBasicAuth("foo", "bar")
	req.ResultFilename("foo.pdf")
	req.WaitTimeout(5)
	req.Landscape(false)
	req.LosslessImageCompression()
	dirPath, err := test.Rand()
	require.Nil(t, err)
	dest := fmt.Sprintf("%s/foo.pdf", dirPath)
	err = c.Store(req, dest)
	assert.Nil(t, err)
	assert.FileExists(t, dest)
	isPDF, err := test.IsPDF(dest)
	assert.Nil(t, err)
	assert.True(t, isPDF)
	err = os.RemoveAll(dirPath)
	assert.Nil(t, err)
}

func TestOfficeCompression(t *testing.T) {
	c := &Client{Hostname: "http://localhost:3000"}
	doc, err := NewDocumentFromPath("document.docx", test.OfficeTestFilePath(t, "document.docx"))
	require.Nil(t, err)
	req := NewOfficeRequest(doc)
	req.SetBasicAuth("foo", "bar")
	req.ResultFilename("foo.pdf")
	req.WaitTimeout(5)
	req.Landscape(false)
	req.Quality(1)
	req.ReduceImageResolution()
	req.MaxImageResolution(75)
	dirPath, err := test.Rand()
	require.Nil(t, err)
	dest := fmt.Sprintf("%s/foo.pdf", dirPath)
	err = c.Store(req, dest)
	assert.Nil(t, err)
	assert.FileExists(t, dest)
	isPDF, err := test.IsPDF(dest)
	assert.Nil(t, err)
	assert.True(t, isPDF)
	err = os.RemoveAll(dirPath)
	assert.Nil(t, err)
}

func TestOfficeWebhook(t *testing.T) {
	c := &Client{Hostname: "http://localhost:3000"}
	doc, err := NewDocumentFromPath("document.docx", test.OfficeTestFilePath(t, "document.docx"))
	require.Nil(t, err)
	req := NewOfficeRequest(doc)
	req.SetBasicAuth("foo", "bar")
	req.WebhookURL("https://google.com")
	req.WebhookURLTimeout(5.0)
	req.AddWebhookURLHTTPHeader("A-Header", "Foo")
	resp, err := c.Post(req)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestOfficeMultipleWithoutMerge(t *testing.T) {
	c := &Client{Hostname: "http://localhost:3000"}
	doc1, err := NewDocumentFromPath("document1.docx", test.OfficeTestFilePath(t, "document.docx"))
	require.Nil(t, err)
	doc2, err := NewDocumentFromPath("document2.docx", test.OfficeTestFilePath(t, "document.docx"))
	require.Nil(t, err)
	req := NewOfficeRequest(doc1, doc2)
	req.SetBasicAuth("foo", "bar")
	req.ResultFilename("foo.zip")
	req.WaitTimeout(5)
	req.Landscape(false)
	dirPath, err := test.Rand()
	require.Nil(t, err)
	dest := fmt.Sprintf("%s/foo.zip", dirPath)
	err = c.Store(req, dest)
	assert.Nil(t, err)
	assert.FileExists(t, dest)

	zipReader, err := zip.OpenReader(dest)
	require.Nil(t, err)

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
		assert.True(t, found, fmt.Sprintf("File %s not found in zip", fileName))
	}
	err = zipReader.Close()
	assert.Nil(t, err)
	err = os.RemoveAll(dirPath)
	assert.Nil(t, err)
}

func TestOfficeMultipleWithMerge(t *testing.T) {
	c := &Client{Hostname: "http://localhost:3000"}
	doc1, err := NewDocumentFromPath("document1.docx", test.OfficeTestFilePath(t, "document.docx"))
	require.Nil(t, err)
	doc2, err := NewDocumentFromPath("document2.docx", test.OfficeTestFilePath(t, "document.docx"))
	require.Nil(t, err)
	req := NewOfficeRequest(doc1, doc2)
	req.SetBasicAuth("foo", "bar")
	req.ResultFilename("foo.pdf")
	req.WaitTimeout(5)
	req.Landscape(false)
	req.Merge()
	dirPath, err := test.Rand()
	require.Nil(t, err)
	dest := fmt.Sprintf("%s/foo.pdf", dirPath)
	err = c.Store(req, dest)
	assert.Nil(t, err)
	assert.FileExists(t, dest)
	isPDF, err := test.IsPDF(dest)
	assert.Nil(t, err)
	assert.True(t, isPDF)
	err = os.RemoveAll(dirPath)
	assert.Nil(t, err)
}

func TestOfficePdfA(t *testing.T) {
	c := &Client{Hostname: "http://localhost:3000"}
	doc, err := NewDocumentFromPath("document.docx", test.OfficeTestFilePath(t, "document.docx"))
	require.Nil(t, err)
	req := NewOfficeRequest(doc)
	req.SetBasicAuth("foo", "bar")
	req.ResultFilename("foo.pdf")
	req.WaitTimeout(5)
	req.Landscape(false)
	req.PdfA(PdfA3b)
	dirPath, err := test.Rand()
	require.Nil(t, err)
	dest := fmt.Sprintf("%s/foo.pdf", dirPath)
	err = c.Store(req, dest)
	assert.Nil(t, err)
	assert.FileExists(t, dest)
	isPDFA, err := test.IsPDFA(dest)
	assert.Nil(t, err)
	assert.True(t, isPDFA)
	err = os.RemoveAll(dirPath)
	assert.Nil(t, err)
}

func TestOfficePdfUA(t *testing.T) {
	c := &Client{Hostname: "http://localhost:3000"}
	doc, err := NewDocumentFromPath("document.docx", test.OfficeTestFilePath(t, "document.docx"))
	require.Nil(t, err)
	req := NewOfficeRequest(doc)
	req.SetBasicAuth("foo", "bar")
	req.ResultFilename("foo.pdf")
	req.WaitTimeout(5)
	req.Landscape(false)
	req.PdfUA()
	dirPath, err := test.Rand()
	require.Nil(t, err)
	dest := fmt.Sprintf("%s/foo.pdf", dirPath)
	err = c.Store(req, dest)
	assert.Nil(t, err)
	assert.FileExists(t, dest)
	isPDFUA, err := test.IsPDFUA(dest)
	assert.Nil(t, err)
	assert.True(t, isPDFUA)
	err = os.RemoveAll(dirPath)
	assert.Nil(t, err)
}
