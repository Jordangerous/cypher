package schema

// ProtoSchema represents the structure of a parsed Proto3 schema
type ProtoSchema struct {
	Messages []ProtoMessage
}

// ProtoMessage represents a message in the Proto3 schema
type ProtoMessage struct {
	Name   string
	Fields []ProtoField
}

// ProtoField represents a field within a Proto3 message
type ProtoField struct {
	Name     string
	DataType string
	Number   int
	Label    string // Can be 'optional', 'required', or 'repeated' in older versions of Proto
}
