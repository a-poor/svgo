package svgo

import (
	"fmt"
	"io/ioutil"
)

type SVG struct {
	Width      int32
	Height     int32
	Children   []SVGElement
	Attributes map[string]string
}

func (svg SVG) ToFile(path string) error {
	svgString := svg.ToString()
	svgBytes := []byte(svgString)
	err := ioutil.WriteFile(path, svgBytes, 0644)
	return err
}

func (svg SVG) ToString() string {
	sattrs := formatAttributes(svg.Attributes)
	schildren := ""
	for _, c := range svg.Children {
		schildren += c.ToString()
	}
	return fmt.Sprintf("<svg viewbox=\"0 0 %d %d\" xmlns=\"http://www.w3.org/2000/svg\" %s>%s</svg>",
		svg.Width, svg.Height, sattrs, schildren)
}

func (svg *SVG) AddChild(child SVGElement) {
	svg.Children = append(svg.Children, child)
}

func (svg *SVG) initAttributes() {
	if svg.Attributes == nil {
		svg.Attributes = make(map[string]string)
	}
}

func (svg *SVG) GetAttr(key string) (string, bool) {
	svg.initAttributes()
	val, inAttrs := svg.Attributes[key]
	return val, inAttrs
}

func (svg *SVG) SetAttr(key string, value string) {
	svg.initAttributes()
	svg.Attributes[key] = value
}

func (svg *SVG) DelAttr(key string) {
	svg.initAttributes()
	delete(svg.Attributes, key)
}
