package cli

import (
	"fmt"
	"os"

	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "linkmap",
	Short: "Linkmap - an open-source URL shortener written in Go.",
	Long: `A minimal URL shortener.
Simple. Open Source. Self-hosted. Written in Go.

Built with 🩵
by casali-dev`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("Linkmap, your links in O(1)")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
