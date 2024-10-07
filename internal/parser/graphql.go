package parser

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// GraphQLSchema represents the structure of a parsed GraphQL schema
type GraphQLSchema struct {
	Types []GraphQLType
}

// GraphQLType represents a single type in a GraphQL schema
type GraphQLType struct {
	Name     string
	Fields   []GraphQLField
	TypeKind string // Could be object, interface, enum, etc.
}

// GraphQLField represents a field within a GraphQL type
type GraphQLField struct {
	Name string
	Type string
}

// ParseGraphQLSchema reads a GraphQL schema file and parses it into a GraphQLSchema struct
func ParseGraphQLSchema(filePath string) (*GraphQLSchema, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read GraphQL schema file: %w", err)
	}

	schema := &GraphQLSchema{}
	schemaString := string(content)
	types := strings.Split(schemaString, "type ")

	for _, typeDef := range types[1:] { // Skipping the first entry as it will be an empty string
		typeParts := strings.Split(typeDef, "{")
		if len(typeParts) < 2 {
			continue
		}

		typeName := strings.TrimSpace(typeParts[0])
		fieldDefinitions := strings.Split(typeParts[1], "}")
		if len(fieldDefinitions) < 1 {
			continue
		}

		fields := parseGraphQLFields(fieldDefinitions[0])
		schema.Types = append(schema.Types, GraphQLType{
			Name:     typeName,
			Fields:   fields,
			TypeKind: "object", // Assuming these are object types for now
		})
	}

	return schema, nil
}

// parseGraphQLFields parses the fields of a GraphQL type from its string representation
func parseGraphQLFields(fieldsString string) []GraphQLField {
	lines := strings.Split(fieldsString, "\n")
	var fields []GraphQLField

	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)
		if trimmedLine == "" || strings.HasPrefix(trimmedLine, "#") { // Ignore empty lines or comments
			continue
		}

		fieldParts := strings.Fields(trimmedLine)
		if len(fieldParts) < 2 {
			continue
		}

		fieldName := fieldParts[0]
		fieldType := fieldParts[1]

		fields = append(fields, GraphQLField{
			Name: fieldName,
			Type: fieldType,
		})
	}

	return fields
}
