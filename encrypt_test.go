package gotenberg

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/starwalkn/gotenberg-go-client/v9/document"
	"github.com/starwalkn/gotenberg-go-client/v9/test"
)

func TestEncrypt(t *testing.T) {
	c := NewClient("http://localhost:3000", http.DefaultClient)

	doc, err := document.FromPath("gotenberg1.pdf", test.PDFTestFilePath(t, "gotenberg.pdf"))
	require.NoError(t, err)

	dest := fmt.Sprintf("%s/foo.pdf", t.TempDir())

	const (
		userPassword  = "abc"
		ownerPassword = "def"
	)

	err = c.PDFEngines().Encrypt(userPassword, ownerPassword, doc).
		Trace("testEncrypt").
		BasicAuth("foo", "bar").
		Store(context.Background(), dest)

	require.NoError(t, err)
	assert.FileExists(t, dest)

	hasPassword, err := test.HasPassword(dest)
	require.NoError(t, err)
	assert.True(t, hasPassword)
}
