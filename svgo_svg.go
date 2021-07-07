package svgo

import (
	"fmt"
	"io/ioutil"
)

// Main SVG parent element <svg></svg>
type SVG struct {
	Width      int32             // SVG Viewbox Width
	Height     int32             // SVG Viewbox Height
	Children   []SVGElement      // List of SVG child elements
	Attributes map[string]string // SVG element's attributes
}

// Save SVG element to file as text
func (svg SVG) ToFile(path string) error {
	svgString := svg.ToString()
	svgBytes := []byte(svgString)
	err := ioutil.WriteFile(path, svgBytes, 0644)
	return err
}

// Convert SVG element to a string
func (svg SVG) ToString() string {
	sattrs := formatAttributes(svg.Attributes)
	schildren := ""
	for _, c := range svg.Children {
		schildren += c.ToString()
	}
	return fmt.Sprintf("<svg viewbox=\"0 0 %d %d\" xmlns=\"http://www.w3.org/2000/svg\" %s>%s</svg>",
		svg.Width, svg.Height, sattrs, schildren)
}

// Add an SVGElement as a child of svg
func (svg *SVG) AddChild(child SVGElement) {
	svg.Children = append(svg.Children, child)
}

// Initialize the svg element's attributes map
// if it isn't already initialized
func (svg *SVG) initAttributes() {
	if svg.Attributes == nil {
		svg.Attributes = make(map[string]string)
	}
}

// Get an attribute from svg's attributes
func (svg *SVG) GetAttr(key string) string {
	svg.initAttributes()
	val := svg.Attributes[key]
	return val
}

// Check if the svg element has attribute `key`
func (svg *SVG) HasAttr(key string) bool {
	svg.initAttributes()
	_, has := svg.Attributes[key]
	return has
}

// Set svg element's attribute `key` to `value`
func (svg *SVG) SetAttr(key string, value string) {
	svg.initAttributes()
	svg.Attributes[key] = value
}

// Delete attribute `key` from svg's elements
// if it exists
func (svg *SVG) DelAttr(key string) {
	svg.initAttributes()
	delete(svg.Attributes, key)
}
