package main

import (
	"fmt"
	"os"
	"strconv"
)

func main01() {
	arguments := os.Args
	sum := 0
	fmt.Println(len(arguments))
	for i := 1; i < len(arguments); i++ {
		temp, _ := strconv.Atoi(arguments[i])
		sum = sum + temp
	}
	fmt.Println("Sum:", sum)
}
