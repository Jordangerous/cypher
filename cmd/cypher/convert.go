package main

import (
	"cypher/internal/converter"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	inputFormat  string
	outputFormat string
	inputFile    string
	outputFile   string
)

var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert schema files from one format to another",
	Long: `Convert schema files such as XSD, SQL, GraphQL, JSON, YAML, and Proto3 to another schema format.
    
Example usage:
  cypher convert --input-format=json --output-format=sql --input-file=data.json --output-file=data.sql`,
	Run: func(cmd *cobra.Command, args []string) {
		if inputFile == "" || outputFile == "" || inputFormat == "" || outputFormat == "" {
			fmt.Println("Error: input-file, output-file, input-format, and output-format are required parameters.")
			cmd.Usage()
			os.Exit(1)
		}

		fmt.Printf("Converting from %s to %s...\n", inputFormat, outputFormat)

		err := converter.Convert(inputFormat, outputFormat, inputFile, outputFile)
		if err != nil {
			fmt.Printf("Error during conversion: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Conversion complete! Output file: %s\n", outputFile)
	},
}

func init() {
	convertCmd.Flags().StringVarP(&inputFormat, "input-format", "i", "", "Format of the input file (e.g., json, xsd, sql, graphql, yaml, proto3)")
	convertCmd.Flags().StringVarP(&outputFormat, "output-format", "o", "", "Desired output format (e.g., json, xsd, sql, graphql, yaml, proto3)")
	convertCmd.Flags().StringVarP(&inputFile, "input-file", "f", "", "Path to the input schema file")
	convertCmd.Flags().StringVarP(&outputFile, "output-file", "d", "", "Path to the output schema file")
}
