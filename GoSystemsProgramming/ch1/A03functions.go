package main

import "fmt"

func main04() {
	//y := 4
	//square := func(s int) int {
	//	return s * s
	//}
	//fmt.Println("The square of", y, "is", square(y))
	//square = func(s int) int {
	//	return s + s
	//}
	//fmt.Println("The square of", y, "is", square(y))
	//fmt.Println(minMax(15, 6))
	//fmt.Println(namedMinMax(15, 6))
	//min, max := namedMinMax(12, -1)
	//fmt.Println(min, max)

	//a1()
	//fmt.Println()
	//a2()
	//fmt.Println()
	//a3()
	//fmt.Println()

	x := -2
	withPointer(&x)
	fmt.Println(x)
	w := newComplex(4, -5)
	fmt.Println(*w)
	fmt.Println(w)
}

func minMax(x, y int) (min, max int) {
	if x > y {
		min = y
		max = x
	} else {
		min = x
		max = y
	}
	return min, max
}

func namedMinMax(x, y int) (min, max int) {
	if x > y {
		min = y
		max = x
	} else {
		min = x
		max = y
	}
	return
}

func a1() {
	for i := 0; i < 3; i++ {
		defer fmt.Print(i, " ")
	}
}

func a2() {
	for i := 0; i < 3; i++ {
		defer func() { fmt.Print(i, " ") }()
	}
}

func a3() {
	for i := 0; i < 3; i++ {
		defer func(n int) { fmt.Print(n, " ") }(i)
	}
}

func withPointer(x *int) {
	*x = *x * *x
}

type complex struct {
	x, y int
}

func newComplex(x, y int) *complex {
	return &complex{x, y}
}
