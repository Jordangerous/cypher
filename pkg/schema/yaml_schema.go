package schema

// YAMLSchema represents the structure of a parsed YAML schema
type YAMLSchema struct {
	Title      string                  `yaml:"title,omitempty"`
	Type       string                  `yaml:"type,omitempty"`
	Properties map[string]YAMLProperty `yaml:"properties,omitempty"`
}

// YAMLProperty represents a property in a YAML schema
type YAMLProperty struct {
	Type        string                  `yaml:"type,omitempty"`
	Description string                  `yaml:"description,omitempty"`
	Properties  map[string]YAMLProperty `yaml:"properties,omitempty"`
	Items       *YAMLProperty           `yaml:"items,omitempty"` // For array types
}
