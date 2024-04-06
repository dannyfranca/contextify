package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	version = "0.0.0"
	commit  = "unknown"
	date    = "unknown"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Contextify",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Contextify version %s\n", getVersion())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

func getVersion() string {
	return fmt.Sprintf("%s, commit %s, built at %s", version, commit, date)
}
