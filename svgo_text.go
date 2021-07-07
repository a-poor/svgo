package main

import "fmt"

type Text struct {
	Text       string
	X          float64
	Y          float64
	Attributes map[string]string
}

func (t Text) ToString() string {
	sattrs := formatAttributes(t.Attributes)
	return fmt.Sprintf("<text x=\"%f\" y=\"%f\" %s>%s</text>",
		t.X, t.Y, sattrs, t.Text)
}

func (t *Text) initAttributes() {
	if t.Attributes == nil {
		t.Attributes = make(map[string]string)
	}
}

func (t *Text) GetAttr(key string) (string, bool) {
	t.initAttributes()
	val, inAttrs := t.Attributes[key]
	return val, inAttrs
}

func (t *Text) SetAttr(key string, value string) {
	t.initAttributes()
	t.Attributes[key] = value
}

func (t *Text) DelAttr(key string) {
	t.initAttributes()
	delete(t.Attributes, key)
}
