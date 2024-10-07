package parser

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// SQLTable represents a table in the SQL schema
type SQLTable struct {
	Name    string
	Columns []SQLColumn
}

// SQLColumn represents a column in an SQL table
type SQLColumn struct {
	Name        string
	DataType    string
	Constraints []string
}

// SQLSchema represents the structure of a parsed SQL schema
type SQLSchema struct {
	Tables []SQLTable
}

// Regex patterns to match SQL components
var tableRegex = regexp.MustCompile(`(?i)CREATE\s+TABLE\s+(\w+)`)
var columnRegex = regexp.MustCompile(`\s*(\w+)\s+(\w+)(.*),?`)

// ParseSQLSchema reads an SQL schema file and parses it into an SQLSchema struct
func ParseSQLSchema(filePath string) (*SQLSchema, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open SQL schema file: %w", err)
	}
	defer file.Close()

	schema := &SQLSchema{}
	scanner := bufio.NewScanner(file)
	var currentTable *SQLTable

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Skip empty lines and comments
		if line == "" || strings.HasPrefix(line, "--") || strings.HasPrefix(line, "/*") || strings.HasPrefix(line, "*") {
			continue
		}

		// Detect the start of a table definition
		if tableMatch := tableRegex.FindStringSubmatch(line); len(tableMatch) > 1 {
			tableName := tableMatch[1]
			currentTable = &SQLTable{Name: tableName}
			schema.Tables = append(schema.Tables, *currentTable)
		} else if currentTable != nil && strings.HasPrefix(line, ");") {
			// End of the current table definition
			currentTable = nil
		} else if currentTable != nil {
			// Parse columns within the table
			if columnMatch := columnRegex.FindStringSubmatch(line); len(columnMatch) > 2 {
				columnName := columnMatch[1]
				dataType := columnMatch[2]
				constraints := strings.Fields(strings.TrimSpace(columnMatch[3]))

				sqlColumn := SQLColumn{
					Name:        columnName,
					DataType:    dataType,
					Constraints: constraints,
				}
				currentTable.Columns = append(currentTable.Columns, sqlColumn)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error while reading SQL schema file: %w", err)
	}

	return schema, nil
}
