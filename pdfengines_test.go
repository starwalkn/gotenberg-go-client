package gotenberg

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/dcaraxes/gotenberg-go-client/document"
	"github.com/dcaraxes/gotenberg-go-client/test"
)

func TestMerge(t *testing.T) {
	c, err := NewClient("http://localhost:3000", &http.Client{})

	require.NoError(t, err)
	pdf1, err := document.FromPath("gotenberg1.pdf", test.PDFTestFilePath(t, "gotenberg.pdf"))
	require.NoError(t, err)
	pdf2, err := document.FromPath("gotenberg2.pdf", test.PDFTestFilePath(t, "gotenberg.pdf"))
	require.NoError(t, err)
	req := NewMergeRequest(pdf1, pdf2)
	req.UseBasicAuth("foo", "bar")
	req.OutputFilename("foo.pdf")
	dirPath, err := test.Rand()
	require.NoError(t, err)
	dest := fmt.Sprintf("%s/foo.pdf", dirPath)
	err = c.Store(req, dest)
	require.NoError(t, err)
	assert.FileExists(t, dest)
	err = os.RemoveAll(dirPath)
	require.NoError(t, err)
}

func TestReadWriteMetadata(t *testing.T) {
	c, err := NewClient("http://localhost:3000", &http.Client{})

	require.NoError(t, err)
	// WRITE
	pdf1, err := document.FromPath("gotenberg1.pdf", test.PDFTestFilePath(t, "gotenberg.pdf"))
	require.NoError(t, err)
	reqWrite := NewWriteMetadataRequest(pdf1)
	reqWrite.UseBasicAuth("foo", "bar")
	reqWrite.OutputFilename("foo.pdf")

	writeDataStruct := struct {
		Author    string `json:"Author"`
		Copyright string `json:"Copyright"`
	}{
		Author:    "Alexander Pikeev",
		Copyright: "Alexander Pikeev",
	}

	writeData, err := json.Marshal(writeDataStruct)
	require.NoError(t, err)
	reqWrite.Metadata(writeData)

	dirPath, err := test.Rand()
	require.NoError(t, err)
	dest := fmt.Sprintf("%s/foo.pdf", dirPath)
	err = c.Store(reqWrite, dest)
	require.NoError(t, err)
	assert.FileExists(t, dest)

	// READ
	pdf2, err := document.FromPath("foo.pdf", dest)
	require.NoError(t, err)
	reqRead := NewReadMetadataRequest(pdf2)
	reqRead.UseBasicAuth("foo", "bar")
	reqRead.OutputFilename("foo.pdf")
	respRead, err := c.Send(reqRead)
	require.NoError(t, err)
	assert.Equal(t, 200, respRead.StatusCode)

	var readData exifData
	err = json.NewDecoder(respRead.Body).Decode(&readData)
	require.NoError(t, err)
	expected := exifData{
		FooPdf: writeDataStruct,
	}
	assert.Equal(t, expected, readData)
	err = os.RemoveAll(dirPath)
	require.NoError(t, err)
}

type exifData struct {
	FooPdf struct {
		Author    string `json:"Author"`
		Copyright string `json:"Copyright"`
	} `json:"foo.pdf"`
}