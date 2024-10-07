package schema

type Schema struct {
	Types []TypeDefinition
}

type TypeDefinition struct {
	Name     string
	Fields   []FieldDefinition
	TypeKind TypeKind // e.g., Object, Enum, etc.
}

type FieldDefinition struct {
	Name     string
	Type     string
	Required bool
}
