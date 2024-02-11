package form

import (
	"cmp"
	"fmt"
	"slices"

	"github.com/bomgar/fbmirrors/mirrorlist"
	"github.com/charmbracelet/huh"
)

type GenerateOptions struct {
	CountryCode string
	IpVersion   []string
	Protocols   []string
}

func GetGenerateOptions(mirrorList *mirrorlist.MirrorList) (GenerateOptions, error) {

	countries := toHuhOptions(mirrorList.GetCountries())

	options := GenerateOptions{}

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Pick a country.").
				Options(countries...).
				Value(&options.CountryCode),
		),
		huh.NewGroup(
			huh.NewMultiSelect[string]().
				Title("IP version").
				Options(
					huh.NewOption("IPv4", "IPv4").Selected(true),
					huh.NewOption("IPv6", "IPv6").Selected(true),
				).
				Value(&options.IpVersion),
			huh.NewMultiSelect[string]().
				Title("Protocol").
				Options(
					huh.NewOption("HTTPS", "https").Selected(true),
					huh.NewOption("HTTP", "http").Selected(false),
					huh.NewOption("RSYNC", "rsync").Selected(false),
				).
				Value(&options.Protocols),
		),
	)

	err := form.Run()
	if err != nil {
		return options, fmt.Errorf("run form: %w", err)
	}

	return options, nil
}

func toHuhOptions(input map[string]string) []huh.Option[string] {
	options := make([]huh.Option[string], len(input))
	i := 0
	for k, v := range input {
		options[i] = huh.NewOption(k, v)
		i++
	}

	slices.SortFunc(options, func(a, b huh.Option[string]) int {
		return cmp.Compare(a.Key, b.Key)
	})

	return options
}
