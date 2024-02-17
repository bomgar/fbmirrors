package mirrorlist

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type MirrorList struct {
	LastCheck      time.Time `json:"last_check"`
	Mirrors        []Mirror  `json:"urls"`
	Cutoff         int64     `json:"cutoff"`
	NumChecks      int       `json:"num_checks"`
	CheckFrequency int       `json:"check_frequency"`
}

type Mirror struct {
	Score          *float64   `json:"score"`
	LastSync       *time.Time `json:"last_sync"`
	Country        string     `json:"country"`
	Protocol       string     `json:"protocol"`
	URL            string     `json:"url"`
	Details        string     `json:"details"`
	CountryCode    string     `json:"country_code"`
	CompletionPct  float64    `json:"completion_pct"`
	DurationStddev float64    `json:"duration_stddev"`
	DurationAvg    float64    `json:"duration_avg"`
	Delay          int        `json:"delay"`
	Active         bool       `json:"active"`
	Isos           bool       `json:"isos"`
	IPv4           bool       `json:"ipv4"`
	IPv6           bool       `json:"ipv6"`
}

func (m *MirrorList) GetCountries() map[string]string {
	countries := make(map[string]string)
	for _, mirror := range m.Mirrors {
		if mirror.CountryCode == "" {
			continue
		}
		countries[mirror.Country] = mirror.CountryCode
	}
	return countries
}

func (m Mirror) PlaceHolderUrl() string {
	return fmt.Sprintf("%s$repo/os/$arch", m.URL)
}

func LoadMirrorList(filepath string) (*MirrorList, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("load mirror list: %w", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	mirrorList := &MirrorList{}
	err = decoder.Decode(mirrorList)
	if err != nil {
		return nil, fmt.Errorf("load mirror list: %w", err)
	}

	return mirrorList, nil
}
