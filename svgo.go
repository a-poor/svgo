package svgo

import (
	"fmt"
	"io/ioutil"
	"strings"
)

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

func (g Group) ToString() string {
	sattrs := formatAttributes(g.Attributes)
	schildren := ""
	for _, c := range g.Children {
		schildren += c.ToString()
	}
	return fmt.Sprintf("<g %s>%s</g>", sattrs, schildren)
}

func (c Circle) ToString() string {
	sattrs := formatAttributes(c.Attributes)
	return fmt.Sprintf("<circle cx=\"%f\" cy=\"%f\" r=\"%f\" %s />", c.Cx, c.Cy, c.R, sattrs)
}

func (e Ellipse) ToString() string {
	sattrs := formatAttributes(e.Attributes)
	return fmt.Sprintf("<ellipse cx=\"%f\" cy=\"%f\" rx=\"%f\" rx=\"%f\" %s />",
		e.Cx, e.Cy, e.Rx, e.Ry, sattrs)
}

func (r Rect) ToString() string {
	sattrs := formatAttributes(r.Attributes)
	return fmt.Sprintf("<rect x=\"%f\" y=\"%f\" width=\"%f\" height=\"%f\" %s />",
		r.X, r.Y, r.Width, r.Height, sattrs)
}

func (l Line) ToString() string {
	sattrs := formatAttributes(l.Attributes)
	return fmt.Sprintf("<line x1=\"%f\" y1=\"%f\" x2=\"%f\" y2=\"%f\" %s />",
		l.X1, l.Y1, l.X2, l.Y2, sattrs)
}

func (t Text) ToString() string {
	sattrs := formatAttributes(t.Attributes)
	return fmt.Sprintf("<text x=\"%f\" y=\"%f\" %s>%s</text>",
		t.X, t.Y, sattrs, t.Text)
}

func (p Polygon) ToString() string {
	sattrs := formatAttributes(p.Attributes)
	points := make([]string, 0)
	for _, pt := range p.Points {
		points = append(points, pt.ToString())
	}
	spoints := strings.Join(points, " ")
	return fmt.Sprintf("<polygon points=\"%s\" %s/>", spoints, sattrs)
}

func (p Point2D) ToString() string {
	return fmt.Sprintf("%f,%f", p.X, p.Y)
}

func (svg SVG) AddChild(child SVGElement) {
	svg.Children = append(svg.Children, child)
}

func (g Group) AddChild(child SVGElement) {
	g.Children = append(g.Children, child)
}

func (pg Polygon) AddPoint2D(p Point2D) {
	pg.Points = append(pg.Points, p)
}

func (pg Polygon) AddPoint(x, y float64) {
	pg.AddPoint2D(Point2D{x, y})
}
