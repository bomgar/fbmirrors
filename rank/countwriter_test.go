package rank

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type TargetUrl string

func (t TargetUrl) TargetUrl() string {
	return string(t)
}

func TestMeasureDownloadSpeed(t *testing.T) {

	testHttpServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_, err := io.WriteString(w, "Hello, world!")
		assert.NoError(t, err)
	}))
	defer testHttpServer.Close()

	speed, err := measureDownloadSpeed(testHttpServer.Client(), testHttpServer.URL)
	require.NoError(t, err)
	require.Greater(t, speed, 0.0)

}

func TestMeasureDownloadSpeedHttpError(t *testing.T) {

	testHttpServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer testHttpServer.Close()

	_, err := measureDownloadSpeed(testHttpServer.Client(), testHttpServer.URL)
	require.Error(t, err)
}

func TestMeasureTargets(t *testing.T) {
	testHttpServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_, err := io.WriteString(w, "Hello, world!")
		assert.NoError(t, err)
	}))
	defer testHttpServer.Close()

	targets := []TargetUrl{TargetUrl(testHttpServer.URL), TargetUrl(testHttpServer.URL), TargetUrl(testHttpServer.URL)}

	results, err := MeasureTargets(testHttpServer.Client(), targets)
	require.NoError(t, err)
	require.Len(t, results, 3)
	for _, result := range results {
		require.Greater(t, result.SpeedInMBps, 0.0)
	}
}
