package main

import (
	"fmt"
)

func main() {
	test_path := "./test.svg"
	svg := SVG{Width: 400, Height: 600}
	fmt.Println(svg.ToString())

	svg.SetAttr("width", "300")

	circle := Circle{Cx: 100, Cy: 300, R: 50}
	fmt.Println(circle.ToString())

	circle.SetAttr("fill", "green")

	svg.AddChild(circle)

	fmt.Println(svg.ToString())

	err := svg.ToFile(test_path)
	if err != nil {
		panic(err)
	}
}
