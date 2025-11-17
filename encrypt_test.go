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

func TestEncrypt(t *testing.T) {
	c, err := NewClient("http://localhost:3000", http.DefaultClient)
	require.NoError(t, err)

	doc, err := document.FromPath("gotenberg1.pdf", test.PDFTestFilePath(t, "gotenberg.pdf"))
	require.NoError(t, err)

	const (
		userPassword  = "abc"
		ownerPassword = "def"
	)

	r := NewEncryptRequest(userPassword, ownerPassword, doc)
	r.Trace("testEncrypt")
	r.UseBasicAuth("foo", "bar")

	dest := fmt.Sprintf("%s/foo.pdf", t.TempDir())
	err = c.Store(context.Background(), r, dest)
	require.NoError(t, err)
	assert.FileExists(t, dest)

	hasPassword, err := test.HasPassword(dest)
	require.NoError(t, err)
	assert.True(t, hasPassword)
}
