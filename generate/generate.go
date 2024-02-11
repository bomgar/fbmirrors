package generate

import (
	"fmt"

	"github.com/bomgar/fbmirrors/form"
	"github.com/bomgar/fbmirrors/mirrorlist"
)

func Generate(mirrorList *mirrorlist.MirrorList) error {
	options, err := form.GetGenerateOptions(mirrorList)
	if err != nil {
		return fmt.Errorf("generate: %w", err)
	}

	fmt.Println(options)

	return nil
}
