/*
Copyright © 2024 Danny Franca me@dannyfranca.com
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "contextify",
	Short: "A command-line interface (CLI) for converting your codebase into markdown format that can be easily comprehended by Artificial Intelligence (AI) systems.",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {}
