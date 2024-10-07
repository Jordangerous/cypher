package main

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cypher",
	Short: "Cypher is a schema conversion CLI tool",
	Long:  "Cypher allows you to convert XSD, SQL, GraphQL, JSON, YAML, and Proto3 schemas to other formats.",
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
