package parser

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v3"
)

// YAMLSchema represents the structure of a parsed YAML schema
type YAMLSchema struct {
	Properties map[string]YAMLProperty `yaml:"properties"`
	Type       string                  `yaml:"type"`
	Title      string                  `yaml:"title,omitempty"`
}

// YAMLProperty represents a property in the YAML schema
type YAMLProperty struct {
	Type        string                  `yaml:"type"`
	Description string                  `yaml:"description,omitempty"`
	Properties  map[string]YAMLProperty `yaml:"properties,omitempty"`
	Items       *YAMLProperty           `yaml:"items,omitempty"` // For array types
}

// ParseYAMLSchema reads a YAML schema file and parses it into a YAMLSchema struct
func ParseYAMLSchema(filePath string) (*YAMLSchema, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open YAML schema file: %w", err)
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read YAML schema file: %w", err)
	}

	var schema YAMLSchema
	err = yaml.Unmarshal(content, &schema)
	if err != nil {
		return nil, fmt.Errorf("failed to parse YAML schema: %w", err)
	}

	return &schema, nil
}
