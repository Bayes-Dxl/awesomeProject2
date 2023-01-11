//package coloredpoint
package main

import (
	geo "awesomeProject2/ch5/geometry"
	"fmt"
	"image/color"
)

//type Point struct {
//	X, Y float64
//}
type ColorPoint struct {
	geo.Point
	Color color.RGBA
}

func main() {
	var cp ColorPoint
	cp.X = 1
	fmt.Println(cp.Point.X)
	cp.Point.Y = 2
	fmt.Println(cp.Y)

	red := color.RGBA{
		R: 255,
		G: 0,
		B: 0,
		A: 255,
	}

	blue := color.RGBA{B: 255, A: 255}
	var p = ColorPoint{Point: geo.Point{X: 1, Y: 1}, Color: red}
	var q = ColorPoint{Point: geo.Point{X: 5, Y: 4}, Color: blue}
	fmt.Println(p.Distance(q.Point))
}
