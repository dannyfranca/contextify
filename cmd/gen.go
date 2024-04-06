/*
Copyright © 2024 Danny Franca me@dannyfranca.com
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generates a markdown file from the codebase provided as input.",
	Long: `Generates a markdown file from the codebase provided as input.
	
	The generated markdown file can be easily comprehended by Artificial Intelligence (AI) systems.
	
	You can provide multiple directory paths and extensions as input to generate a combined markdown file.`,
	Run: runGen,
}

var (
	outputFile string
	extensions string
)

func init() {
	rootCmd.AddCommand(genCmd)

	genCmd.Flags().StringVarP(&outputFile, "output", "o", "codebase.md", "Output file path")
	genCmd.MarkFlagRequired("output")
	genCmd.Flags().StringVarP(&extensions, "extensions", "e", "go,md,json,html,templ", "File extensions to include (comma-separated)")
	genCmd.MarkFlagRequired("extensions")
}

func runGen(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Println("Please provide at least one directory path")
		os.Exit(1)
	}

	directoryPaths := args
	extensionList := strings.Split(extensions, ",")

	err := cleanOutputFile(outputFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, directoryPath := range directoryPaths {
		err = processDirectory(directoryPath, outputFile, extensionList)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	fmt.Printf("Combined files saved to %s\n", outputFile)
}

func cleanOutputFile(outputPath string) error {
	file, err := os.OpenFile(outputPath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	return nil
}

func processDirectory(directoryPath, outputFile string, extensionList []string) error {
	file, err := os.OpenFile(outputFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	err = filepath.Walk(directoryPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if shouldProcessFile(info, extensionList) {
			err := processFile(path, directoryPath, writer)
			if err != nil {
				return err
			}
		}

		return nil
	})

	return err
}

func shouldProcessFile(info os.FileInfo, extensionList []string) bool {
	if info.IsDir() {
		return false
	}

	for _, ext := range extensionList {
		if strings.HasSuffix(info.Name(), "."+ext) {
			return true
		}
	}

	return false
}

func processFile(filePath, directoryPath string, writer *bufio.Writer) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	relativePath := strings.TrimPrefix(filePath, directoryPath+string(os.PathSeparator))
	entryPoint := filepath.Base(directoryPath)
	writer.WriteString(fmt.Sprintf("# %s/%s\n\n", entryPoint, relativePath))

	extension := filepath.Ext(filePath)
	writer.WriteString(fmt.Sprintf("```%s\n", extension[1:]))

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		writer.WriteString(scanner.Text() + "\n")
	}

	writer.WriteString("```\n\n")

	return scanner.Err()
}
