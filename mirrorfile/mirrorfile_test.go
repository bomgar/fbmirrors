package mirrorfile

import (
	_ "embed"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

//go:embed testdata/mirrorlist
var content string

func TestParseContent(t *testing.T) {
	mirrors, err := parseContent(strings.NewReader(content))
	require.NoError(t, err)

	require.Equal(t, 6, len(mirrors))
	require.Contains(t, mirrors, MirrorEntry("https://archlinux.thaller.ws/$repo/os/$arch"))

}
