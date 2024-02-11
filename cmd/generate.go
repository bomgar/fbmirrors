package cmd

import (
	"log"
	"net/http"
	"os"

	"github.com/bomgar/fbmirrors/generate"
	"github.com/bomgar/fbmirrors/mirrorlist"
	"github.com/bomgar/fbmirrors/mirrorservice"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generates a mirror list",
	Long:  `Fetches available mirrors and generates a mirror list. Available mirrors are cached for 24 hours.`,
	Run: func(cmd *cobra.Command, _ []string) {
		url, _ := cmd.Flags().GetString("url")

		cacheDir, err := os.UserCacheDir()
		if err != nil {
			log.Fatal(err)
		}

		client := &http.Client{}

		path, err := mirrorservice.RefreshIfNecessary(client, cacheDir, url)
		if err != nil {
			log.Fatal(err)
		}

		mirrorList, err := mirrorlist.LoadMirrorList(path)
		if err != nil {
			log.Fatal(err)
		}
		err = generate.Generate(mirrorList)
		if err != nil {
			log.Fatal(err)
		}

	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().StringP("url", "u", "https://archlinux.org/mirrors/status/json/", "Arch json mirror list url")

}
