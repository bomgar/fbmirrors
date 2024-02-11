package mirrorservice

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFetchAvailableMirrors(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "fbmirrors-test")
	require.NoError(t, err)

	data := `{"mock": "data"}`

	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, data)
	}))
	defer mockServer.Close()

	cacheFile, err := FetchAvailableMirrors(mockServer.Client(), tempDir, mockServer.URL)
	require.NoError(t, err)

	bytes, err := os.ReadFile(cacheFile)
	require.NoError(t, err)

	content := string(bytes)

	require.Equal(t, data, content)

	defer os.RemoveAll(tempDir)
}
