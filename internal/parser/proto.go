package parser

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ProtoMessage represents a Protocol Buffers message with its fields
type ProtoMessage struct {
	Name   string
	Fields []ProtoField
}

// ProtoField represents a field in a Protocol Buffers message
type ProtoField struct {
	Name     string
	DataType string
	Number   int
}

// ProtoSchema represents the structure of a parsed Proto3 schema
type ProtoSchema struct {
	Messages []ProtoMessage
}

// ParseProtoSchema reads a Proto3 schema file and parses it into a ProtoSchema struct
func ParseProtoSchema(filePath string) (*ProtoSchema, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open Proto3 schema file: %w", err)
	}
	defer file.Close()

	schema := &ProtoSchema{}
	scanner := bufio.NewScanner(file)

	var currentMessage *ProtoMessage

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Skip empty lines and comments
		if line == "" || strings.HasPrefix(line, "//") || strings.HasPrefix(line, "/*") || strings.HasPrefix(line, "*") {
			continue
		}

		// Detect the start of a message definition
		if strings.HasPrefix(line, "message ") {
			messageName := strings.TrimSpace(strings.TrimPrefix(line, "message "))
			currentMessage = &ProtoMessage{Name: messageName}
			schema.Messages = append(schema.Messages, *currentMessage)
		} else if currentMessage != nil && strings.HasPrefix(line, "}") {
			// End of the current message
			currentMessage = nil
		} else if currentMessage != nil {
			// Parse fields within a message
			fieldParts := strings.Fields(line)
			if len(fieldParts) >= 3 {
				fieldType := fieldParts[0]
				fieldName := fieldParts[1]
				fieldNumber := strings.TrimSuffix(fieldParts[2], ";")

				protoField := ProtoField{
					Name:     fieldName,
					DataType: fieldType,
				}

				fmt.Sscanf(fieldNumber, "%d", &protoField.Number)
				currentMessage.Fields = append(currentMessage.Fields, protoField)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error while reading Proto3 schema file: %w", err)
	}

	return schema, nil
}
