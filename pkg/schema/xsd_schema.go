package schema

// XSDSchema represents the entire structure of an XSD schema
type XSDSchema struct {
	Elements []XSDElement
}

// XSDElement represents a single element in the XSD schema
type XSDElement struct {
	Name       string
	Type       string
	Attributes []XSDAttribute
	Children   []XSDElement
}

// XSDAttribute represents an attribute of an XSD element
type XSDAttribute struct {
	Name string
	Type string
}

// NewXSDSchema creates a new instance of XSDSchema
func NewXSDSchema(elements []XSDElement) *XSDSchema {
	return &XSDSchema{Elements: elements}
}

// AddElement adds a new element to the XSDSchema
func (schema *XSDSchema) AddElement(element XSDElement) {
	schema.Elements = append(schema.Elements, element)
}

// NewXSDElement creates a new instance of XSDElement
func NewXSDElement(name, elementType string, attributes []XSDAttribute, children []XSDElement) XSDElement {
	return XSDElement{
		Name:       name,
		Type:       elementType,
		Attributes: attributes,
		Children:   children,
	}
}

// AddChildElement adds a new child element to an existing XSDElement
func (element *XSDElement) AddChildElement(child XSDElement) {
	element.Children = append(element.Children, child)
}

// NewXSDAttribute creates a new instance of XSDAttribute
func NewXSDAttribute(name, attributeType string) XSDAttribute {
	return XSDAttribute{
		Name: name,
		Type: attributeType,
	}
}
