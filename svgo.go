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

type Point2D struct {
	x, y float64
}


type SVG struct {
	width  int32
	height int32
	children []SVGElement
	attributes map[string]string
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
	x float64
	y float64
	width float64
	height float64
	attributes map[string]string
}

type Line struct {
	x1 float64
	y1 float64
	x2 float64
	y2 float64
	attributes map[string]string
}

type Text struct {
	attributes map[string]string
	text string
}

type Polygon struct {
	attributes map[string]string
	points []Point2D
}



func (svg SVG) toString() string {
	sattrs := formatAttributes(svg.attributes)
	schildren := ""
	for _, c := range svg.children {
		schildren += c.toString()
	}
	return fmt.Sprintf("<svg viewbox=\"0 0 %d %d\" xmlns=\"http://www.w3.org/2000/svg\" %s>%s</svg>",
		svg.width, svg.height, sattrs, schildren)
}

func (g Group) toString() string {
	sattrs := formatAttributes(g.attributes)
	schildren := ""
	for _, c := range g.children {
		schildren += c.toString()
	}
	return fmt.Sprintf("<g %s>%s</g>", sattrs, schildren)
}

func (c Circle) toString() string {
	return fmt.Sprintf("<circle cx=\"%f\" cy=\"%f\" r=\"%f\" />", c.cx, c.cy, c.r)
}

func (e Ellipse) toString() string {
	return fmt.Sprintf("<ellipse cx=\"%f\" cy=\"%f\" rx=\"%f\" rx=\"%f\" />", e.cx, e.cy, e.rx, e.ry)
}

func (r Rect) toString() string {
	return fmt.Sprintf("<rect x=\"%f\" y=\"%f\" width=\"%f\" height=\"%f\" />", r.x, r.y, r.width, r.height)
}

func (l Line) toString() string {
	return fmt.Sprintf("<line x1=\"%f\" y1=\"%f\" x2=\"%f\" y2=\"%f\" />", l.x1, l.y1, l.x2, l.y2)
}




func formatAttributes(attrs map[string]string) string {
	var sattrs string
	for k, v := range(attrs) {
		sattrs += fmt.Sprintf("%s=\"%s\" ", k, v)
	}
	return strings.TrimSpace(sattrs)
}
