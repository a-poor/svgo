package main

import "fmt"

type Ellipse struct {
	Cx         float64
	Cy         float64
	Rx         float64
	Ry         float64
	Attributes map[string]string
}

func (e Ellipse) ToString() string {
	sattrs := formatAttributes(e.Attributes)
	return fmt.Sprintf("<ellipse cx=\"%f\" cy=\"%f\" rx=\"%f\" rx=\"%f\" %s />",
		e.Cx, e.Cy, e.Rx, e.Ry, sattrs)
}

func (e *Ellipse) initAttributes() {
	if e.Attributes == nil {
		e.Attributes = make(map[string]string)
	}
}

func (e *Ellipse) GetAttr(key string) (string, bool) {
	e.initAttributes()
	val, inAttrs := e.Attributes[key]
	return val, inAttrs
}

func (e *Ellipse) SetAttr(key string, value string) {
	e.initAttributes()
	e.Attributes[key] = value
}

func (e *Ellipse) DelAttr(key string) {
	e.initAttributes()
	delete(e.Attributes, key)
}

