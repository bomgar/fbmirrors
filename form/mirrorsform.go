package form

import (
	"fmt"

	"github.com/bomgar/fbmirrors/mirrorlist"
	"github.com/charmbracelet/huh"
)

func SelectMirrors(mirrors []mirrorlist.Mirror) ([]mirrorlist.Mirror, error) {
	selectedMirrors := []mirrorlist.Mirror{}
	mirrorOptions := []huh.Option[mirrorlist.Mirror]{}

	for _, mirror := range mirrors {
		mirrorOptions = append(mirrorOptions, huh.NewOption(
			fmt.Sprintf("%s (score: %f, last_synced: %v)", mirror.URL, *mirror.Score, *mirror.LastSync),
			mirror,
		))
	}

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewMultiSelect[mirrorlist.Mirror]().
				Title("Select mirrors").
				Options(mirrorOptions...).
				Value(&selectedMirrors),
		),
	)

	err := form.Run()
	if err != nil {
		return selectedMirrors, fmt.Errorf("select mirrors: %w", err)
	}

	return selectedMirrors, nil
}
