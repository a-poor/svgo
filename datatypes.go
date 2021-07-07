package svgo

type SVGElement interface {
	ToString() string
}

type SVGParent interface {
	AddChild(c SVGElement)
}

type Point2D struct {
	x, y float64
}

type SVG struct {
	width      int32
	height     int32
	children   []SVGElement
	attributes map[string]string
}

type Group struct {
	attributes map[string]string
	children   []SVGElement
}

type Circle struct {
	cx         float64
	cy         float64
	r          float64
	attributes map[string]string
}

type Ellipse struct {
	cx         float64
	cy         float64
	rx         float64
	ry         float64
	attributes map[string]string
}

type Rect struct {
	x          float64
	y          float64
	width      float64
	height     float64
	attributes map[string]string
}

type Line struct {
	x1         float64
	y1         float64
	x2         float64
	y2         float64
	attributes map[string]string
}

type Text struct {
	text       string
	x          float64
	y          float64
	attributes map[string]string
}

type Polygon struct {
	attributes map[string]string
	points     []Point2D
}
