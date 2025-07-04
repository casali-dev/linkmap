package cli

import (
	"github.com/casali-dev/linkmap/internal/db"
	"github.com/casali-dev/linkmap/internal/services"
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(shortenCmd)
}

var shortenCmd = &cobra.Command{
	Use:   "shorten",
	Short: "Creates the shortened version of the given URL",
	Long:  `Register the given URL, creating the shortened version of the url.`,
	Run: func(cmd *cobra.Command, args []string) {
		db := db.Init()

		if len(args) < 1 {
			log.Error("You must provide a URL.")
			return
		}

		url := args[0]

		slug, err := services.ShortenUrl(db, url)
		if err != nil {
			log.Error("Failed to shorten URL", "err", err)
			return
		}

		log.Infof("Url shortened! Slug: %s", slug)
	},
}
