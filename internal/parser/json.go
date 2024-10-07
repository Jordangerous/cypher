package parser

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// JSONSchema represents the structure of a parsed JSON schema
type JSONSchema struct {
	Title      string                  `json:"title,omitempty"`
	Type       string                  `json:"type,omitempty"`
	Properties map[string]JSONProperty `json:"properties,omitempty"`
}

// JSONProperty represents a property in a JSON schema
type JSONProperty struct {
	Type        string                  `json:"type,omitempty"`
	Description string                  `json:"description,omitempty"`
	Properties  map[string]JSONProperty `json:"properties,omitempty"`
	Items       *JSONProperty           `json:"items,omitempty"` // For array types
}

// ParseJSONSchema reads a JSON schema file and parses it into a JSONSchema struct
func ParseJSONSchema(filePath string) (*JSONSchema, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read JSON schema file: %w", err)
	}

	var schema JSONSchema
	err = json.Unmarshal(content, &schema)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JSON schema: %w", err)
	}

	return &schema, nil
}
