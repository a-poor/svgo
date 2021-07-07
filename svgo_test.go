package svgo

import (
	"testing"
)

func TestSimple(t *testing.T) {
	test_path := "./test.svg"
	svg := SVG{Width: 400, Height: 600}
	svgStr := svg.ToString()
	if len(svgStr) == 0 {
		t.Error("Error creating SVG string")
	}

	svg.SetAttr("width", "300")
	circle := Circle{Cx: 100, Cy: 300, R: 50}
	svgStr = svg.ToString()
	if len(svgStr) == 0 {
		t.Error("Error creating SVG string")
	}

	circle.SetAttr("fill", "green")
	svg.AddChild(circle)
	svgStr = svg.ToString()
	if len(svgStr) == 0 {
		t.Error("Error creating SVG string")
	}

	err := svg.ToFile(test_path)
	if err != nil {
		t.Error("Failed to write file")
	}
}
