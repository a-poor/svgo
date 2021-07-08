package svgo

import (
	"fmt"
	"strings"
)

type Element struct {
	TagName string
	Attributes map[string]string
	Children []interface{}
}

type Stringer interface {
	String() string
}

func (e Element) String() string {
	return fmt.Sprintf(
		"<%s %s>%s</%s>",
		e.TagName,
		formatAttributes(e.Attributes),
		formatChildren(e.Children),
		e.TagName,
	)
}

type Point2D struct {
	X, Y float64
}

func (p Point2D) String() string {
	return fmt.Sprintf("%f,%f", p.X, p.Y)
}

func formatChildren(children []interface{}) string {
	res := ""
	for _, c := range children {
		e, ok := c.(Element)
		if ok { // `c` is an `Element`
			res += e.String()
		} else { // Treat `c` as a `string`
			s, _ := c.(string)
			res += s
		}
	}
	return res
}

func formatAttributes(attrs map[string]string) string {
	var sattrs string
	for k, v := range attrs {
		sattrs += fmt.Sprintf("%s=\"%s\" ", k, v)
	}
	return strings.TrimSpace(sattrs)
}
