package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/bomgar/fbmirrors/mirrorservice"
	"github.com/spf13/cobra"
)

var fetchAvailableMirrorsCmd = &cobra.Command{
	Use:   "fetch-available-mirrors",
	Short: "Updates the cached list of available mirrors.",
	Long:  `Updates the mirror cache`,
	Run: func(cmd *cobra.Command, _ []string) {
		url, _ := cmd.Flags().GetString("url")

		cacheDir, err := os.UserCacheDir()
		if err != nil {
			log.Fatal(err)
		}

		client := &http.Client{}

		path, err := mirrorservice.FetchAvailableMirrors(client, cacheDir, url)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Mirrors list updated: %s\n", path)
	},
}

func init() {
	rootCmd.AddCommand(fetchAvailableMirrorsCmd)

	fetchAvailableMirrorsCmd.Flags().StringP("url", "u", "https://archlinux.org/mirrors/status/json/", "Arch json mirror list url")
}
