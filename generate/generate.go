package generate

import (
	"cmp"
	"fmt"
	"slices"

	"github.com/bomgar/fbmirrors/form"
	"github.com/bomgar/fbmirrors/mirrorlist"
)

func Generate(mirrorList *mirrorlist.MirrorList) error {
	options, err := form.GetGenerateOptions(mirrorList)
	if err != nil {
		return fmt.Errorf("generate: %w", err)
	}

	filteredMirrors := filterMirrors(mirrorList.Mirrors, options)

	slices.SortFunc(filteredMirrors, func(a, b mirrorlist.Mirror) int {
		return cmp.Compare(*a.Score, *b.Score)
	})

	selectedMirrors, err := form.SelectMirrors(filteredMirrors)
	if err != nil {
		return fmt.Errorf("generate: %w", err)
	}

	printMirrors(selectedMirrors)

	return nil
}

func printMirrors(mirrors []mirrorlist.Mirror) {
	for _, mirror := range mirrors {
		fmt.Printf("Server = %s\n", mirror.PlaceHolderUrl())
	}
}

func filterMirrors(mirrorList []mirrorlist.Mirror, options form.GenerateOptions) []mirrorlist.Mirror {
	mirrors := []mirrorlist.Mirror{}

	for _, mirror := range mirrorList {
		if mirror.Score == nil {
			continue
		}
		if mirror.LastSync == nil {
			continue
		}
		if !mirror.Active {
			continue
		}

		if mirror.CountryCode != options.CountryCode {
			continue
		}

		if !slices.Contains(options.Protocols, mirror.Protocol) {
			continue
		}

		if options.IpVersion == form.Ipv4 && !mirror.IPv4 {
			continue
		}

		if options.IpVersion == form.Ipv6 && !mirror.IPv6 {
			continue
		}

		mirrors = append(mirrors, mirror)
	}

	return mirrors
}
