package svgo

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// Main SVG parent element `<svg></svg>`. SVG elements have a
// fixed `xmlns="http://www.w3.org/2000/svg"` property.
type SVG struct {
	VboxWidth      int32             // SVG Viewbox Width
	VboxHeight     int32             // SVG Viewbox Height
	Width          int32             // SVG Width
	Height         int32             // SVG Height
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

// Sets the svg's width and height to a
// non-zero value. If width/height are set
// and viewbox is not set, viewbox will be
// set to match width / height. Otherwise
// they will be set to `100`. If width/height
// aren't set, they will be set to `100`.
func (svg *SVG) validateSvgShape() {
	// Set ViewBox params
	if svg.VboxWidth <= 0 {
		if svg.Width > 0 {
			svg.VboxWidth = svg.Width
		} else {
			svg.VboxWidth = 100
		}
	}
	if svg.VboxHeight <= 0 {
		if svg.Height > 0 {
			svg.VboxHeight = svg.Height
		} else {
			svg.VboxHeight = 100
		}
	}
	// Set Size params
	if svg.Width <= 0 {
		svg.Width = 100
	}

	if svg.Height <= 0 {
		svg.Height = 100
	}

}

// Convert SVG element to a string
func (svg SVG) ToString() string {
	svg.validateSvgShape()
	sattrs := formatAttributes(svg.Attributes)
	schildren := ""
	for _, c := range svg.Children {
		schildren += c.ToString()
	}
	return fmt.Sprintf("<svg width=\"%d\" height=\"%d\" viewbox=\"0 0 %d %d\" xmlns=\"http://www.w3.org/2000/svg\" %s>%s</svg>",
		svg.Width, svg.Height, svg.VboxWidth, svg.VboxHeight, sattrs, schildren)
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

// Set svg element's attribute `key` to `value`.
// If `key == "width" or key == "height"`, the svg's
// Width/Height attribute will be changed.
func (svg *SVG) SetAttr(key string, value string) error {
	if strings.ToLower(key) == "width" {
		w, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		svg.Width = int32(w)
	} else if  strings.ToLower(key) == "height" {
		h, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		svg.Height = int32(h)
	} else {
		svg.initAttributes()
		svg.Attributes[key] = value
	}
	return nil
}

// Delete attribute `key` from svg's elements
// if it exists
func (svg *SVG) DelAttr(key string) {
	svg.initAttributes()
	delete(svg.Attributes, key)
}
