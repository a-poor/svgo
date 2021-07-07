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
	text string
	x float64
	y float64
	attributes map[string]string
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
	sattrs := formatAttributes(c.attributes)
	return fmt.Sprintf("<circle cx=\"%f\" cy=\"%f\" r=\"%f\" %s />", c.cx, c.cy, c.r, sattrs)
}

func (e Ellipse) toString() string {
	sattrs := formatAttributes(e.attributes)
	return fmt.Sprintf("<ellipse cx=\"%f\" cy=\"%f\" rx=\"%f\" rx=\"%f\" %s />",
		e.cx, e.cy, e.rx, e.ry, sattrs)
}

func (r Rect) toString() string {
	sattrs := formatAttributes(r.attributes)
	return fmt.Sprintf("<rect x=\"%f\" y=\"%f\" width=\"%f\" height=\"%f\" %s />",
		r.x, r.y, r.width, r.height, sattrs)
}

func (l Line) toString() string {
	sattrs := formatAttributes(l.attributes)
	return fmt.Sprintf("<line x1=\"%f\" y1=\"%f\" x2=\"%f\" y2=\"%f\" %s />",
		l.x1, l.y1, l.x2, l.y2, sattrs)
}

func (t Text) toString() string {
	sattrs := formatAttributes(t.attributes)
	return fmt.Sprintf("<text x=\"%s\" y=\"%\" %s>%s</text>",
		t.x, t.y, sattrs, t.text)
}

func (p Polygon) toString() string {
	sattrs := formatAttributes(p.attributes)
	points := make([]string, 0)
	for _, p := range p.points {
		points = append(points, fmt.Sprintf("%f,%f", p.x, p.y))
	}
	spoints := strings.Join(points, " ")
	return fmt.Sprintf("<polygon points=\"%s\" %s/>", spoints, sattrs)
}



func (svg SVG) addChild(child SVGElement) {
	svg.children = append(svg.children, child)
}

func (g Group) addChild(child SVGElement) {
	g.children = append(g.children, child)
}


func (pg Polygon) addPoint2D(p Point2D) {
	pg.points = append(pg.points, p)
}

func (pg Polygon) addPoint(x, y float64) {
	pg.addPoint2D(Point2D{x, y})
}

func formatAttributes(attrs map[string]string) string {
	var sattrs string
	for k, v := range(attrs) {
		sattrs += fmt.Sprintf("%s=\"%s\" ", k, v)
	}
	return strings.TrimSpace(sattrs)
}
