package mirrorlist

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoadMirrorList(t *testing.T) {

	mirrorList, err := LoadMirrorList("testdata/mirrors.json")
	require.NoError(t, err)

	require.Equal(t, 1165, len(mirrorList.Mirrors))

	mirror := mirrorList.Mirrors[0]
	require.Equal(t, "https://mirror.aarnet.edu.au/pub/archlinux/", mirror.URL)
}

func TestCountries(t *testing.T) {
	mirrorList, err := LoadMirrorList("testdata/mirrors.json")
	require.NoError(t, err)

	countries := mirrorList.GetCountries()
	require.Contains(t, countries, "Australia")
	require.Equal(t, 70, len(countries))
}
