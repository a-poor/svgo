package svgo

import "fmt"

type Circle struct {
	Cx         float64
	Cy         float64
	R          float64
	Attributes map[string]string
}

func (c Circle) ToString() string {
	sattrs := formatAttributes(c.Attributes)
	return fmt.Sprintf("<circle cx=\"%f\" cy=\"%f\" r=\"%f\" %s />", c.Cx, c.Cy, c.R, sattrs)
}

func (c *Circle) initAttributes() {
	if c.Attributes == nil {
		c.Attributes = make(map[string]string)
	}
}

func (c *Circle) GetAttr(key string) (string, bool) {
	c.initAttributes()
	val, inAttrs := c.Attributes[key]
	return val, inAttrs
}

func (c *Circle) SetAttr(key string, value string) {
	c.initAttributes()
	c.Attributes[key] = value
}

func (c *Circle) DelAttr(key string) {
	c.initAttributes()
	delete(c.Attributes, key)
}
