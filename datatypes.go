package svgo

type SVGElement interface {
	ToString() string
}

type SVGParent interface {
	AddChild(c SVGElement)
}

type Point2D struct {
	X, Y float64
}

type SVG struct {
	Width      int32
	Height     int32
	Children   []SVGElement
	Attributes map[string]string
}

type Group struct {
	Attributes map[string]string
	Children   []SVGElement
}

type Circle struct {
	Cx         float64
	Cy         float64
	R          float64
	Attributes map[string]string
}

type Ellipse struct {
	Cx         float64
	Cy         float64
	Rx         float64
	Ry         float64
	Attributes map[string]string
}

type Rect struct {
	X          float64
	Y          float64
	Width      float64
	Height     float64
	Attributes map[string]string
}

type Line struct {
	X1         float64
	Y1         float64
	X2         float64
	Y2         float64
	Attributes map[string]string
}

type Text struct {
	Text       string
	X          float64
	Y          float64
	Attributes map[string]string
}

type Polygon struct {
	Attributes map[string]string
	Points     []Point2D
}
