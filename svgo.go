package svgo

import (
	"fmt"
	"strings"
)

func (svg SVG) ToString() string {
	sattrs := formatAttributes(svg.attributes)
	schildren := ""
	for _, c := range svg.children {
		schildren += c.ToString()
	}
	return fmt.Sprintf("<svg viewbox=\"0 0 %d %d\" xmlns=\"http://www.w3.org/2000/svg\" %s>%s</svg>",
		svg.width, svg.height, sattrs, schildren)
}

func (g Group) ToString() string {
	sattrs := formatAttributes(g.attributes)
	schildren := ""
	for _, c := range g.children {
		schildren += c.ToString()
	}
	return fmt.Sprintf("<g %s>%s</g>", sattrs, schildren)
}

func (c Circle) ToString() string {
	sattrs := formatAttributes(c.attributes)
	return fmt.Sprintf("<circle cx=\"%f\" cy=\"%f\" r=\"%f\" %s />", c.cx, c.cy, c.r, sattrs)
}

func (e Ellipse) ToString() string {
	sattrs := formatAttributes(e.attributes)
	return fmt.Sprintf("<ellipse cx=\"%f\" cy=\"%f\" rx=\"%f\" rx=\"%f\" %s />",
		e.cx, e.cy, e.rx, e.ry, sattrs)
}

func (r Rect) ToString() string {
	sattrs := formatAttributes(r.attributes)
	return fmt.Sprintf("<rect x=\"%f\" y=\"%f\" width=\"%f\" height=\"%f\" %s />",
		r.x, r.y, r.width, r.height, sattrs)
}

func (l Line) ToString() string {
	sattrs := formatAttributes(l.attributes)
	return fmt.Sprintf("<line x1=\"%f\" y1=\"%f\" x2=\"%f\" y2=\"%f\" %s />",
		l.x1, l.y1, l.x2, l.y2, sattrs)
}

func (t Text) ToString() string {
	sattrs := formatAttributes(t.attributes)
	return fmt.Sprintf("<text x=\"%f\" y=\"%f\" %s>%s</text>",
		t.x, t.y, sattrs, t.text)
}

func (p Polygon) ToString() string {
	sattrs := formatAttributes(p.attributes)
	points := make([]string, 0)
	for _, pt := range p.points {
		points = append(points, pt.ToString())
	}
	spoints := strings.Join(points, " ")
	return fmt.Sprintf("<polygon points=\"%s\" %s/>", spoints, sattrs)
}

func (p Point2D) ToString() string {
	return fmt.Sprintf("%f,%f", p.x, p.y)
}

func (svg SVG) AddChild(child SVGElement) {
	svg.children = append(svg.children, child)
}

func (g Group) AddChild(child SVGElement) {
	g.children = append(g.children, child)
}

func (pg Polygon) AddPoint2D(p Point2D) {
	pg.points = append(pg.points, p)
}

func (pg Polygon) AddPoint(x, y float64) {
	pg.AddPoint2D(Point2D{x, y})
}
