package main

import "fmt"

type Group struct {
	Attributes map[string]string
	Children   []SVGElement
}

func (g Group) ToString() string {
	sattrs := formatAttributes(g.Attributes)
	schildren := ""
	for _, c := range g.Children {
		schildren += c.ToString()
	}
	return fmt.Sprintf("<g %s>%s</g>", sattrs, schildren)
}

func (g *Group) AddChild(child SVGElement) {
	g.Children = append(g.Children, child)
}

func (g *Group) initAttributes() {
	if g.Attributes == nil {
		g.Attributes = make(map[string]string)
	}
}

func (g *Group) GetAttr(key string) (string, bool) {
	g.initAttributes()
	val, inAttrs := g.Attributes[key]
	return val, inAttrs
}

func (g *Group) SetAttr(key string, value string) {
	g.initAttributes()
	g.Attributes[key] = value
}

func (g *Group) DelAttr(key string) {
	g.initAttributes()
	delete(g.Attributes, key)
}
