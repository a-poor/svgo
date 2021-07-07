package main

import "fmt"

type Line struct {
	X1         float64
	Y1         float64
	X2         float64
	Y2         float64
	Attributes map[string]string
}

func (l Line) ToString() string {
	sattrs := formatAttributes(l.Attributes)
	return fmt.Sprintf("<line x1=\"%f\" y1=\"%f\" x2=\"%f\" y2=\"%f\" %s />",
		l.X1, l.Y1, l.X2, l.Y2, sattrs)
}

func (l *Line) initAttributes() {
	if l.Attributes == nil {
		l.Attributes = make(map[string]string)
	}
}

func (l *Line) GetAttr(key string) (string, bool) {
	l.initAttributes()
	val, inAttrs := l.Attributes[key]
	return val, inAttrs
}

func (l *Line) SetAttr(key string, value string) {
	l.initAttributes()
	l.Attributes[key] = value
}

func (l *Line) DelAttr(key string) {
	l.initAttributes()
	delete(l.Attributes, key)
}

