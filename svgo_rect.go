package svgo

import "fmt"

type Rect struct {
	X          float64
	Y          float64
	Width      float64
	Height     float64
	Attributes map[string]string
}

func (r Rect) ToString() string {
	sattrs := formatAttributes(r.Attributes)
	return fmt.Sprintf("<rect x=\"%f\" y=\"%f\" width=\"%f\" height=\"%f\" %s />",
		r.X, r.Y, r.Width, r.Height, sattrs)
}

func (r *Rect) initAttributes() {
	if r.Attributes == nil {
		r.Attributes = make(map[string]string)
	}
}

func (r *Rect) GetAttr(key string) (string, bool) {
	r.initAttributes()
	val, inAttrs := r.Attributes[key]
	return val, inAttrs
}

func (r *Rect) SetAttr(key string, value string) {
	r.initAttributes()
	r.Attributes[key] = value
}

func (r *Rect) DelAttr(key string) {
	r.initAttributes()
	delete(r.Attributes, key)
}