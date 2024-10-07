// Package schema provides type definitions and constants used across the Cypher project.
package schema

// TypeKind represents the kind of a type definition (e.g., Object, Enum).
type TypeKind string

const (
	// TypeKindObject represents a standard object type with fields.
	TypeKindObject TypeKind = "Object"
	// TypeKindEnum represents an enumeration type with predefined values.
	TypeKindEnum TypeKind = "Enum"
	// TypeKindInterface represents an interface type.
	TypeKindInterface TypeKind = "Interface"
	// TypeKindUnion represents a union type.
	TypeKindUnion TypeKind = "Union"
	// TypeKindScalar represents a scalar type.
	TypeKindScalar TypeKind = "Scalar"
)

// ScalarTypes is a set of recognized scalar types.
var ScalarTypes = map[string]bool{
	"string":  true,
	"int":     true,
	"float":   true,
	"boolean": true,
	"ID":      true,
	"date":    true,
	"time":    true,
	"decimal": true,
	"byte":    true,
	"short":   true,
	"long":    true,
	"char":    true,
	"binary":  true,
	"uuid":    true,
	"json":    true,
	"object":  true,
	"array":   true,
	"map":     true,
	"set":     true,
	"list":    true,
	"tuple":   true,
	"any":     true,
	"void":    true,
	"null":    true,
	"none":    true,
	"unknown": true,
	"custom":  true,
	"complex": true,
	"simple":  true,
	"mixed":   true,
	"variant": true,
}

// MapSchemaType maps a type string to a standardized type name.
func MapSchemaType(typeStr string) string {
	switch typeStr {
	case "string", "xs:string":
		return "string"
	case "int", "integer", "xs:int":
		return "int"
	case "float", "double", "xs:float", "xs:double":
		return "float"
	case "bool", "boolean", "xs:boolean":
		return "boolean"
	default:
		// For complex or unrecognized types, return the type string as is.
		return typeStr
	}
}

// IsScalarType checks if a given type is a scalar type.
func IsScalarType(typeStr string) bool {
	return ScalarTypes[typeStr]
}

// Additional helper functions and constants can be added here as needed.
