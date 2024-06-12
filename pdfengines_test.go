package gotenberg

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/dcaraxes/gotenberg-go-client/v8/test"
)

func TestMerge(t *testing.T) {
	c := &Client{Hostname: "http://localhost:3000"}
	pdf1, err := NewDocumentFromPath("gotenberg1.pdf", test.PDFTestFilePath(t, "gotenberg.pdf"))
	require.Nil(t, err)
	pdf2, err := NewDocumentFromPath("gotenberg2.pdf", test.PDFTestFilePath(t, "gotenberg.pdf"))
	require.Nil(t, err)
	req := NewMergeRequest(pdf1, pdf2)
	req.SetBasicAuth("foo", "bar")
	req.ResultFilename("foo.pdf")
	req.WaitTimeout(5)
	dirPath, err := test.Rand()
	require.Nil(t, err)
	dest := fmt.Sprintf("%s/foo.pdf", dirPath)
	err = c.Store(req, dest)
	assert.Nil(t, err)
	assert.FileExists(t, dest)
	err = os.RemoveAll(dirPath)
	assert.Nil(t, err)
}

func TestMergeWebhook(t *testing.T) {
	c := &Client{Hostname: "http://localhost:3000"}
	pdf1, err := NewDocumentFromPath("gotenberg1.pdf", test.PDFTestFilePath(t, "gotenberg.pdf"))
	require.Nil(t, err)
	pdf2, err := NewDocumentFromPath("gotenberg2.pdf", test.PDFTestFilePath(t, "gotenberg.pdf"))
	require.Nil(t, err)
	req := NewMergeRequest(pdf1, pdf2)
	req.SetBasicAuth("foo", "bar")
	req.WebhookURL("https://google.com")
	req.WebhookURLTimeout(5.0)
	req.AddWebhookURLHTTPHeader("A-Header", "Foo")
	resp, err := c.Post(req)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestReadWriteMetadata(t *testing.T) {
	c := &Client{Hostname: "http://localhost:3000"}
	// WRITE
	pdf1, err := NewDocumentFromPath("gotenberg1.pdf", test.PDFTestFilePath(t, "gotenberg.pdf"))
	require.Nil(t, err)
	reqWrite := NewWriteMetadataRequest(pdf1)
	reqWrite.SetBasicAuth("foo", "bar")
	reqWrite.ResultFilename("foo.pdf")
	reqWrite.WaitTimeout(5)

	writeDataStruct := struct {
		Author    string `json:"Author"`
		Copyright string `json:"Copyright"`
	}{
		Author:    "Alexander Pikeev",
		Copyright: "Alexander Pikeev",
	}

	writeData, err := json.Marshal(writeDataStruct)
	assert.Nil(t, err)
	reqWrite.Metadata(writeData)

	dirPath, err := test.Rand()
	require.Nil(t, err)
	dest := fmt.Sprintf("%s/foo.pdf", dirPath)
	err = c.Store(reqWrite, dest)
	assert.Nil(t, err)
	assert.FileExists(t, dest)

	// READ
	pdf2, err := NewDocumentFromPath("foo.pdf", dest)
	require.Nil(t, err)
	reqRead := NewReadMetadataRequest(pdf2)
	reqRead.SetBasicAuth("foo", "bar")
	reqRead.ResultFilename("foo.pdf")
	reqRead.WaitTimeout(5)
	respRead, err := c.Post(reqRead)
	assert.Nil(t, err)
	assert.Equal(t, 200, respRead.StatusCode)

	var readData exifData
	err = json.NewDecoder(respRead.Body).Decode(&readData)
	assert.Nil(t, err)
	expected := exifData{
		FooPdf: writeDataStruct,
	}
	assert.Equal(t, expected, readData)
	err = os.RemoveAll(dirPath)
	assert.Nil(t, err)
}

type exifData struct {
	FooPdf struct {
		Author    string `json:"Author"`
		Copyright string `json:"Copyright"`
	} `json:"foo.pdf"`
}
