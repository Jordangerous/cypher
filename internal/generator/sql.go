package generator

import (
	"cypher/pkg/schema"
	"fmt"
	"os"
	"strings"
)

type SQLGenerator struct{}

func (g *SQLGenerator) Generate(s *schema.Schema, outputFile string) error {
	var sqlStatements []string

	for _, typeDef := range s.Types {
		if typeDef.TypeKind != schema.TypeKindObject {
			continue
		}
		createTableStmt := generateCreateTableStatement(typeDef)
		sqlStatements = append(sqlStatements, createTableStmt)
	}

	err := os.WriteFile(outputFile, []byte(strings.Join(sqlStatements, "\n")), 0644)
	if err != nil {
		return err
	}

	return nil
}

func generateCreateTableStatement(typeDef schema.TypeDefinition) string {
	var columns []string
	for _, field := range typeDef.Fields {
		sqlType := mapSchemaTypeToSQLType(field.Type)
		columnDef := fmt.Sprintf("    %s %s", field.Name, sqlType)
		columns = append(columns, columnDef)
	}
	createTableStmt := fmt.Sprintf("CREATE TABLE %s (\n%s\n);", typeDef.Name, strings.Join(columns, ",\n"))
	return createTableStmt
}

func mapSchemaTypeToSQLType(schemaType string) string {
	switch schemaType {
	case "string":
		return "VARCHAR(255)"
	case "int", "integer":
		return "INT"
	case "float", "double":
		return "FLOAT"
	case "bool", "boolean":
		return "BOOLEAN"
	case "id", "uuid":
		return "UUID"
	case "date":
		return "DATE"
	case "datetime", "timestamp":
		return "TIMESTAMP"
	case "time":
		return "TIME"
	case "binary":
		return "BLOB"
	default:
		return "TEXT"
	}
}
