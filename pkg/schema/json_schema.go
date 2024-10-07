package schema

// JSONSchema represents the overall structure of a JSON schema
type JSONSchema struct {
	Title      string                  `json:"title,omitempty"`
	Type       string                  `json:"type,omitempty"`
	Properties map[string]JSONProperty `json:"properties,omitempty"`
}

// JSONProperty represents a property in a JSON schema
type JSONProperty struct {
	Type        string                  `json:"type,omitempty"`
	Description string                  `json:"description,omitempty"`
	Properties  map[string]JSONProperty `json:"properties,omitempty"`
	Items       *JSONProperty           `json:"items,omitempty"` // For array types
}
