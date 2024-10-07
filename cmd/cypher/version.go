package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Define the version of the Cypher tool
const version = "0.1.0"

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Cypher",
	Long:  "Print the version number of the Cypher schema conversion CLI tool.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Cypher CLI tool version: %s\n", version)
	},
}

func init() {
	// Register the version command to the root command
	rootCmd.AddCommand(versionCmd)
}
