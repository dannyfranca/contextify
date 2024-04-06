package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	version = "0.1.0"
	commit  = "9677bacb5c249f43a86a1ecdc049c1fde02b9f7f"
	date    = "2024-04-06"
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
