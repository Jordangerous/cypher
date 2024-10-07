// Package generator provides implementations for generating schema files in various formats.
// This file contains the XSDGenerator which generates XSD schemas.
package generator

import (
	"bytes"
	"cypher/pkg/schema"
	"fmt"
	"os"
)

// XSDGenerator is responsible for generating XSD schema files from the intermediate schema representation.
type XSDGenerator struct{}

// Generate creates an XSD schema file based on the provided Schema.
// The generated file is written to the specified outputFile path.
func (g *XSDGenerator) Generate(s *schema.Schema, outputFile string) error {
	// Generate XSD content from the Schema struct
	var buffer bytes.Buffer

	// Write the XML header and schema opening tag
	buffer.WriteString(`<?xml version="1.0" encoding="UTF-8"?>` + "\n")
	buffer.WriteString(`<xs:schema xmlns:xs="http://www.w3.org/2001/XMLSchema">` + "\n")

	// Generate XSD elements for each type in the schema
	for _, typeDef := range s.Types {
		if typeDef.TypeKind != schema.TypeKindObject {
			continue
		}

		element, err := generateXSDElement(typeDef)
		if err != nil {
			return fmt.Errorf("failed to generate XSD element: %v", err)
		}
		buffer.WriteString(element)
	}

	// Close the schema tag
	buffer.WriteString(`</xs:schema>`)

	// Write the XSD content to the output file
	err := os.WriteFile(outputFile, buffer.Bytes(), 0644)
	if err != nil {
		return fmt.Errorf("failed to write XSD schema to file: %v", err)
	}

	return nil
}

// Helper function to generate an XSD element for a given type definition
func generateXSDElement(typeDef schema.TypeDefinition) (string, error) {
	var buffer bytes.Buffer

	buffer.WriteString(fmt.Sprintf(`  <xs:element name="%s">`+"\n", typeDef.Name))
	buffer.WriteString(`    <xs:complexType>` + "\n")
	buffer.WriteString(`      <xs:sequence>` + "\n")

	for _, field := range typeDef.Fields {
		xsdType := mapSchemaTypeToXSDType(field.Type)
		elementLine := fmt.Sprintf(`        <xs:element name="%s" type="%s" minOccurs="%d" maxOccurs="1"/>`+"\n",
			field.Name, xsdType, boolToInt(!field.Required))
		buffer.WriteString(elementLine)
	}

	buffer.WriteString(`      </xs:sequence>` + "\n")
	buffer.WriteString(`    </xs:complexType>` + "\n")
	buffer.WriteString(`  </xs:element>` + "\n")

	return buffer.String(), nil
}

// Helper function to map schema types to XSD types
func mapSchemaTypeToXSDType(schemaType string) string {
	switch schemaType {
	case "string":
		return "xs:string"
	case "int", "integer":
		return "xs:int"
	case "float", "double":
		return "xs:double"
	case "bool", "boolean":
		return "xs:boolean"
	// Add more mappings as needed
	default:
		return "xs:string"
	}
}

// Helper function to convert a boolean to an integer (0 or 1)
func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}
