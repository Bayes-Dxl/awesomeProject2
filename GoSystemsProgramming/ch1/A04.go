package main

import "fmt"

func main() {
	myArray := []int{1, 2, 3, -4}
	twoD := [3][3]int{{1, 2, 3}, {3, 4, 5}, {7, 8, 9}}
	threeD := [2][2][2]int{{{1, 2}, {3, 4}}, {{5, 6}, {7, 8}}}
	fmt.Println(myArray)
	fmt.Println(twoD)
	fmt.Println(threeD)
	myArray[0] = 8
	twoD[1][2] = 15
	threeD[1][0][1] = -1
	fmt.Println(myArray)
	fmt.Println(twoD)
	fmt.Println(threeD)

	fmt.Println("myArray[10]:", myArray[10])

}
