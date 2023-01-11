package main

import (
	geo "awesomeProject2/ch5/geometry"
	"fmt"
)

func main() {
	p := geo.Point{X: 1, Y: 2}
	q := geo.Point{X: 4, Y: 6}
	fmt.Println(geo.Distance(p, q))
	fmt.Println(p.Distance(q))

	perm := geo.Path{{1, 1}, {5, 1}, {5, 4}, {1, 1}}
	fmt.Println(perm.Distance())

	r := &geo.Point{X: 1, Y: 2}
	r.ScaleBy(5)
	fmt.Println(*r)

	p = geo.Point{
		X: 5,
		Y: 5,
	}

	p.ScaleBy(9)
	fmt.Println(p)
}
