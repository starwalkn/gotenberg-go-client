package gotenberg

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/runatal/gotenberg-go-client/v8/document"
	"github.com/runatal/gotenberg-go-client/v8/test"
)

func TestReadWriteMetadata(t *testing.T) {
	c, err := NewClient("http://localhost:3000", http.DefaultClient)
	require.NoError(t, err)

	// Writing metadata.
	pdf1, err := document.FromPath("gotenberg1.pdf", test.PDFTestFilePath(t, "gotenberg.pdf"))
	require.NoError(t, err)
	reqWrite := NewWriteMetadataRequest(pdf1)
	reqWrite.Trace("testWriteMetadata")
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

	dirPath := t.TempDir()
	dest := fmt.Sprintf("%s/foo.pdf", dirPath)
	err = c.Store(context.Background(), reqWrite, dest)
	require.NoError(t, err)
	assert.FileExists(t, dest)

	// Reading metadata.
	pdf2, err := document.FromPath("foo.pdf", dest)
	require.NoError(t, err)
	reqRead := NewReadMetadataRequest(pdf2)
	reqRead.Trace("testReadMetadata")
	reqRead.UseBasicAuth("foo", "bar")
	reqRead.OutputFilename("foo.pdf")
	respRead, err := c.Send(context.Background(), reqRead)
	require.NoError(t, err)
	assert.Equal(t, 200, respRead.StatusCode)

	var readData exifData
	err = json.NewDecoder(respRead.Body).Decode(&readData)
	require.NoError(t, err)
	expected := exifData{
		FooPdf: writeDataStruct,
	}
	assert.Equal(t, expected, readData)
}

type exifData struct {
	FooPdf struct {
		Author    string `json:"Author"`
		Copyright string `json:"Copyright"`
	} `json:"foo.pdf"`
}
