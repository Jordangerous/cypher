package parser

import (
	"encoding/xml"
	"fmt"
	"os"
)

// XSDSchema represents the structure of a parsed XSD schema
type XSDSchema struct {
	Elements []XSDElement
}

// XSDElement represents an element in the XSD schema
type XSDElement struct {
	Name       string
	Type       string
	Attributes []XSDAttribute
	Children   []XSDElement
}

// XSDAttribute represents an attribute of an element in the XSD schema
type XSDAttribute struct {
	Name string
	Type string
}

// ParseXSDSchema reads an XSD schema file and parses it into an XSDSchema struct
func ParseXSDSchema(filePath string) (*XSDSchema, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open XSD schema file: %w", err)
	}
	defer file.Close()

	var xsdRoot struct {
		Elements []XSDElement `xml:"element"`
	}

	decoder := xml.NewDecoder(file)
	if err := decoder.Decode(&xsdRoot); err != nil {
		return nil, fmt.Errorf("failed to parse XSD schema: %w", err)
	}

	schema := &XSDSchema{
		Elements: xsdRoot.Elements,
	}

	return schema, nil
}

// parseXMLElement recursively parses XML elements into XSDElement structs
func parseXMLElement(element xml.StartElement) XSDElement {
	var xsdElement XSDElement
	xsdElement.Name = getAttribute(element.Attr, "name")
	xsdElement.Type = getAttribute(element.Attr, "type")

	for _, attr := range element.Attr {
		xsdElement.Attributes = append(xsdElement.Attributes, XSDAttribute{
			Name: attr.Name.Local,
			Type: attr.Value,
		})
	}

	return xsdElement
}

// getAttribute is a helper function that retrieves the value of a specific attribute from a list of XML attributes
func getAttribute(attrs []xml.Attr, name string) string {
	for _, attr := range attrs {
		if attr.Name.Local == name {
			return attr.Value
		}
	}
	return ""
}
