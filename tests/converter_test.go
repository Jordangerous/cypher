package test

import (
	"cypher/internal/converter"
	"testing"
)

// Test JSON to SQL conversion
func TestConvertJSONToSQL(t *testing.T) {
	err := converter.Convert("json", "sql", "example.json", "output.sql")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

// Test XSD to GraphQL conversion
func TestConvertXSDToGraphQL(t *testing.T) {
	err := converter.Convert("xsd", "graphql", "example.xsd", "output.graphql")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

// Test YAML to Proto3 conversion
func TestConvertYAMLToProto3(t *testing.T) {
	err := converter.Convert("yaml", "proto3", "example.yaml", "output.proto3")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

// Test unsupported conversion
func TestUnsupportedConversion(t *testing.T) {
	err := converter.Convert("json", "unsupported_format", "example.json", "output.txt")
	if err == nil {
		t.Fatalf("Expected an error for unsupported conversion, got none")
	}

	expectedError := "Conversion from json to unsupported_format is not supported"
	if err.Error() != expectedError {
		t.Errorf("Expected error message %q, got %q", expectedError, err.Error())
	}
}
