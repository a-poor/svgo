package main

import (
	"fmt"
	"strings"
)

type Polygon struct {
	Attributes map[string]string
	Points     []Point2D
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

func (p *Polygon) AddPoint2D(pt Point2D) {
	p.Points = append(p.Points, pt)
}

func (p *Polygon) AddPoint(x, y float64) {
	p.AddPoint2D(Point2D{x, y})
}

func (p *Polygon) initAttributes() {
	if p.Attributes == nil {
		p.Attributes = make(map[string]string)
	}
}

func (p *Polygon) GetAttr(key string) (string, bool) {
	p.initAttributes()
	val, inAttrs := p.Attributes[key]
	return val, inAttrs
}

func (p *Polygon) SetAttr(key string, value string) {
	p.initAttributes()
	p.Attributes[key] = value
}

func (p *Polygon) DelAttr(key string) {
	p.initAttributes()
	delete(p.Attributes, key)
}
