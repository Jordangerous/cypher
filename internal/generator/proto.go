// Package generator provides implementations for generating schema files in various formats.
// This file contains the ProtoGenerator which generates Proto3 schemas.
package generator

import (
	"cypher/pkg/schema"
	"fmt"
	"os"
	"strings"
)

// ProtoGenerator is responsible for generating Proto3 schema files from the intermediate schema representation.
type ProtoGenerator struct{}

// Generate creates a Proto3 schema file based on the provided Schema.
// The generated file is written to the specified outputFile path.
func (g *ProtoGenerator) Generate(s *schema.Schema, outputFile string) error {
	var protoMessages []string

	for _, typeDef := range s.Types {
		if typeDef.TypeKind != schema.TypeKindObject {
			continue
		}
		messageDef := generateProtoMessage(typeDef)
		protoMessages = append(protoMessages, messageDef)
	}

	output := strings.Join(protoMessages, "\n\n")
	// Add syntax declaration at the top
	output = "syntax = \"proto3\";\n\n" + output

	err := os.WriteFile(outputFile, []byte(output), 0644)
	if err != nil {
		return fmt.Errorf("failed to write Proto3 schema to file: %v", err)
	}

	return nil
}

// Helper function to generate a Proto3 message definition
func generateProtoMessage(typeDef schema.TypeDefinition) string {
	var fields []string
	fieldNumber := 1
	for _, field := range typeDef.Fields {
		protoType := mapSchemaTypeToProtoType(field.Type)
		// Handle repeated fields if necessary
		repeated := ""
		if strings.HasPrefix(protoType, "repeated ") {
			repeated = "repeated "
			protoType = strings.TrimPrefix(protoType, "repeated ")
		}
		fieldDef := fmt.Sprintf("    %s%s %s = %d;", repeated, protoType, field.Name, fieldNumber)
		fields = append(fields, fieldDef)
		fieldNumber++
	}
	messageDef := fmt.Sprintf("message %s {\n%s\n}", typeDef.Name, strings.Join(fields, "\n"))
	return messageDef
}

// Helper function to map schema types to Proto3 data types
func mapSchemaTypeToProtoType(schemaType string) string {
	switch schemaType {
	case "string":
		return "string"
	case "int", "integer":
		return "int32"
	case "float", "double":
		return "double"
	case "bool", "boolean":
		return "bool"
	case "[]string":
		return "repeated string"
	case "[]int", "[]integer":
		return "repeated int32"
	// Add more mappings as needed
	default:
		return "string"
	}
}
