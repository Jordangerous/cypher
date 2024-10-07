package schema

// GraphQLSchema represents the structure of a parsed GraphQL schema
type GraphQLSchema struct {
	Types []GraphQLType
}

// GraphQLType represents a single type in a GraphQL schema
type GraphQLType struct {
	Name       string
	Kind       string // Could be object, interface, enum, etc.
	Fields     []GraphQLField
	Interfaces []string
}

// GraphQLField represents a field within a GraphQL type
type GraphQLField struct {
	Name       string
	Type       string
	Arguments  []GraphQLArgument
	IsNullable bool
}

// GraphQLArgument represents an argument in a GraphQL field
type GraphQLArgument struct {
	Name       string
	Type       string
	IsOptional bool
}
