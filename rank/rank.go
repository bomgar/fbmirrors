package rank

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

type Target interface {
	TargetUrl() string
}

type MeasureResult[T Target] struct {
	Target      T
	SpeedInMBps float64
}

func MeasureTargets[T Target](client *http.Client, list []T) ([]MeasureResult[T], error) {
	results := make([]MeasureResult[T], len(list))

	for i, target := range list {
		speed, err := measureDownloadSpeed(client, target.TargetUrl())
		if err != nil {
			return nil, fmt.Errorf("measure targets: %w", err)
		}
		results[i] = MeasureResult[T]{Target: target, SpeedInMBps: speed}
	}

	return results, nil
}

func measureDownloadSpeed(client *http.Client, url string) (float64, error) {
	start := time.Now()

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0, fmt.Errorf("measure download speed: %w", err)
	}

	response, err := client.Do(request)
	if err != nil {
		return 0, fmt.Errorf("measure download speed: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("measure download speed: status code: %d", response.StatusCode)
	}

	countWriter := &countWriter{}

	_, err = io.Copy(countWriter, response.Body)
	if err != nil {
		return 0, fmt.Errorf("measure download speed: %w", err)
	}

	bytesPerNanoSecond := float64(countWriter.count) / float64(time.Since(start).Nanoseconds())
	megaBytesPerSecond := bytesPerNanoSecond / 1e6

	return megaBytesPerSecond, nil
}
