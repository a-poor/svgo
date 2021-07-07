package svgo

import (
	"fmt"
	"strings"
)

type SVGElement interface {
	ToString() string
	//GetAttr(key string) string
	//SetAttr(key string, value string) string
	//DelAttr(key string) string
}

type SVGParent interface {
	AddChild(c SVGElement)
}

type Point2D struct {
	X, Y float64
}

func (p Point2D) ToString() string {
	return fmt.Sprintf("%f,%f", p.X, p.Y)
}

func formatAttributes(attrs map[string]string) string {
	var sattrs string
	for k, v := range attrs {
		sattrs += fmt.Sprintf("%s=\"%s\" ", k, v)
	}
	return strings.TrimSpace(sattrs)
}
