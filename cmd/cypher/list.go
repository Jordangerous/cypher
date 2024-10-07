package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Supported schema formats
var supportedFormats = []string{
	"json",
	"xsd",
	"sql",
	"graphql",
	"yaml",
	"proto3",
}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all supported schema formats",
	Long:  "List all schema formats that can be used for input and output conversions in the Cypher CLI tool.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Supported schema formats:")
		for _, format := range supportedFormats {
			fmt.Printf(" - %s\n", format)
		}
	},
}

func init() {
	// Register the list command to the root command
	rootCmd.AddCommand(listCmd)
}
