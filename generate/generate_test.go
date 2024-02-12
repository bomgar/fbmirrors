package generate

import (
	"fmt"
	"testing"
	"time"

	"github.com/bomgar/fbmirrors/form"
	"github.com/bomgar/fbmirrors/mirrorlist"
	"github.com/stretchr/testify/require"
)

func TestFilterMirrors(t *testing.T) {
	score := 1.0
	lastSync := time.Now()
	mirror1 := mirrorlist.Mirror{
		LastSync:    &lastSync,
		Score:       &score,
		Active:      true,
		CountryCode: "US",
		Protocol:    "https",
		IPv4:        true,
		IPv6:        false,
	}
	mirror2 := mirrorlist.Mirror{
		LastSync:    &lastSync,
		Score:       &score,
		Active:      true,
		CountryCode: "DE",
		Protocol:    "https",
		IPv4:        true,
		IPv6:        false,
	}
	mirror3 := mirrorlist.Mirror{
		LastSync:    &lastSync,
		Score:       &score,
		Active:      true,
		CountryCode: "US",
		Protocol:    "https",
		IPv4:        false,
		IPv6:        true,
	}
	mirror4 := mirrorlist.Mirror{
		LastSync:    &lastSync,
		Score:       &score,
		Active:      true,
		CountryCode: "US",
		Protocol:    "rsync",
		IPv4:        true,
		IPv6:        false,
	}

	mirrors := []mirrorlist.Mirror{
		mirror1,
		mirror2,
		mirror3,
		mirror4,
	}

	filteredMirrors := filterMirrors(mirrors, form.GenerateOptions{
		CountryCode: "US",
		IpVersion:   form.Ipv4,
		Protocols:   []string{"https"},
	})

	fmt.Println(filteredMirrors)

	require.Contains(t, filteredMirrors, mirror1)
}
