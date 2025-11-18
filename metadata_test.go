package gotenberg

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/starwalkn/gotenberg-go-client/v8/document"
	"github.com/starwalkn/gotenberg-go-client/v8/test"
)

func TestReadWriteMetadata(t *testing.T) {
	c := NewClient("http://localhost:3000", http.DefaultClient)

	pdf1, err := document.FromPath("gotenberg1.pdf", test.PDFTestFilePath(t, "gotenberg.pdf"))
	require.NoError(t, err)

	dest := fmt.Sprintf("%s/foo.pdf", t.TempDir())

	writeDataStruct := struct {
		Author    string `json:"Author"`
		Copyright string `json:"Copyright"`
	}{
		Author:    "Alexander Pikeev",
		Copyright: "Alexander Pikeev",
	}

	writeData, err := json.Marshal(writeDataStruct)
	require.NoError(t, err)

	err = c.PDFEngines().WriteMetadata(pdf1).
		Trace("testWriteMetadata").
		BasicAuth("foo", "bar").
		OutputFilename("foo.pdf").
		Metadata(writeData).
		Store(context.Background(), dest)

	require.NoError(t, err)
	assert.FileExists(t, dest)

	pdf2, err := document.FromPath("foo.pdf", dest)
	require.NoError(t, err)

	resp, err := c.PDFEngines().ReadMetadata(pdf2).
		Trace("testReadMetadata").
		BasicAuth("foo", "bar").
		OutputFilename("foo.pdf").
		Send(context.Background())

	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	var readData exifData
	err = json.NewDecoder(resp.Body).Decode(&readData)
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
