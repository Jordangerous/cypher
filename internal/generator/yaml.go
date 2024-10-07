// Package generator provides implementations for generating schema files in various formats.
// This file contains the YAMLGenerator which generates YAML schemas.
package generator

import (
	"cypher/pkg/schema"
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

// YAMLGenerator is responsible for generating YAML schema files from the intermediate schema representation.
type YAMLGenerator struct{}

// Generate creates a YAML schema file based on the provided Schema.
// The generated file is written to the specified outputFile path.
func (g *YAMLGenerator) Generate(s *schema.Schema, outputFile string) error {
	// Convert the Schema struct to a YAML representation
	data, err := yaml.Marshal(s)
	if err != nil {
		return fmt.Errorf("failed to marshal schema to YAML: %v", err)
	}

	// Write the YAML data to the output file
	err = os.WriteFile(outputFile, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write YAML schema to file: %v", err)
	}

	return nil
}
