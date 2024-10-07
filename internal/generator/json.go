// Package generator provides implementations for generating schema files in various formats.
// This file contains the JSONGenerator which generates JSON schemas.
package generator

import (
	"cypher/pkg/schema"
	"encoding/json"
	"fmt"
	"os"
)

// Generate creates a JSON schema file based on the provided Schema.
// The generated file is written to the specified outputFile path.
func (g *JSONGenerator) Generate(s *schema.Schema, outputFile string) error {
	// Convert the Schema struct to a JSON representation
	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal schema to JSON: %v", err)
	}

	// Write the JSON data to the output file
	err = os.WriteFile(outputFile, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write JSON schema to file: %v", err)
	}

	return nil
}
