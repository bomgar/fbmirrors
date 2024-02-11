package mirrorservice

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func download(client *http.Client, url string, w io.Writer) error {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("fetch available mirrors: %w", err)
	}
	response, err := client.Do(request)
	if err != nil {
		return fmt.Errorf("fetch available mirrors: %w", err)
	}

	defer response.Body.Close()

	_, err = io.Copy(w, response.Body)
	if err != nil {
		return fmt.Errorf("fetch available mirrors: %w", err)
	}

	return nil
}

// FetchAvailableMirrors fetches the available mirrors from the archlinux.org
// and returns the path to the cached file.
func FetchAvailableMirrors(client *http.Client, cacheDirectory, url string) (string, error) {

	fbmirrorsCacheDir := filepath.Join(cacheDirectory, "fbmirros")
	cacheFile := filepath.Join(fbmirrorsCacheDir, "mirrors.json")

	err := os.MkdirAll(fbmirrorsCacheDir, 0755)
	if err != nil {
		return "", fmt.Errorf("fetch available mirrors: %w", err)
	}

	file, err := os.Create(cacheFile)
	if err != nil {
		return "", fmt.Errorf("fetch available mirrors: %w", err)
	}
	defer file.Close()

	err = download(client, url, file)
	if err != nil {
		return "", fmt.Errorf("fetch available mirrors: %w", err)
	}

	return cacheFile, nil
}
