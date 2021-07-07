package svgo

import (
	"fmt"
	"strings"
)

type SVGElement interface {
	toString() string
}

type SVGParent interface {
	addChild(c SVGElement)
}

//type Element struct {
//	attributes map[string]string
//	children []Element
//}

//type Point2D struct {
//	x, y float64
//}


type SVG struct {
	width float64
	height float64
	children []SVGElement
}



type Group struct {
	attributes map[string]string
	children []SVGElement
}

type Circle struct {
	cx float64
	cy float64
	r  float64
	attributes map[string]string
}

type Ellipse struct {
	cx float64
	cy float64
	rx float64
	ry float64
	attributes map[string]string
}

type Rect struct {
	attributes map[string]string
}

type Text struct {
	attributes map[string]string
	text string
}

type Line struct {
	attributes map[string]string
}

func (g Group) toString() {
	sattrs := formatAttributes(g)
	start_tag := fmt.Sprintf("<%s %s>", e.tagname, sattrs)
	middle_tags := ""
	for _, c := range e.children {

	}
	end_tag := fmt.Sprintf("</%s>", e.tagname)
	return start_tag + middle_tags + end_tag
}



func (e *Element) toString() string {
	sattrs := formatAttributes(e)
	start_tag := fmt.Sprintf("<%s %s>", e.tagname, sattrs)
	middle_tags := ""
	for _, c := range e.children {

	}
	end_tag := fmt.Sprintf("</%s>", e.tagname)
	return start_tag + middle_tags + end_tag
}

func formatAttributes(e *Element) string {
	var attrs string
	for k, v := range(e.attributes) {
		attrs += fmt.Sprintf("%s=\"%s\" ", k, v)
	}
	return strings.TrimSpace(attrs)
}
