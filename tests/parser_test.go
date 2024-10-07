package test

import (
	"cypher/internal/parser"
	"os"
	"testing"
)

// Helper function to create test files
func createTestFile(fileName, content string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(content)
	return err
}

// Test GraphQL parser
func TestParseGraphQLSchema(t *testing.T) {
	testContent := `
        type User {
            id: ID
            name: String
            email: String
        }

        type Post {
            id: ID
            title: String
            content: String
            author: User
        }
    `
	fileName := "test.graphql"
	err := createTestFile(fileName, testContent)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	defer os.Remove(fileName)

	schema, err := parser.ParseGraphQLSchema(fileName)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(schema.Types) != 2 {
		t.Errorf("Expected 2 types, got %d", len(schema.Types))
	}

	if schema.Types[0].Name != "User" || len(schema.Types[0].Fields) != 3 {
		t.Errorf("Expected type 'User' with 3 fields, got type '%s' with %d fields", schema.Types[0].Name, len(schema.Types[0].Fields))
	}
}

// Test JSON parser
func TestParseJSONSchema(t *testing.T) {
	testContent := `{
        "title": "Person",
        "type": "object",
        "properties": {
            "id": {
                "type": "string",
                "description": "The unique identifier for a person"
            },
            "name": {
                "type": "string",
                "description": "Name of the person"
            },
            "age": {
                "type": "integer",
                "description": "Age of the person"
            }
        }
    }`
	fileName := "test.json"
	err := createTestFile(fileName, testContent)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	defer os.Remove(fileName)

	schema, err := parser.ParseJSONSchema(fileName)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if schema.Title != "Person" {
		t.Errorf("Expected schema title 'Person', got '%s'", schema.Title)
	}

	if len(schema.Properties) != 3 {
		t.Errorf("Expected 3 properties, got %d", len(schema.Properties))
	}

	if schema.Properties["id"].Type != "string" {
		t.Errorf("Expected property 'id' of type 'string', got '%s'", schema.Properties["id"].Type)
	}
}
