package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Linkmap",
	Long:  `All software has versions. This is Linkmap's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Linkmap v0.1.0")
	},
}
