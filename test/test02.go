//package main
//
//import "fmt"
//
//func main02() {
//	x := 200
//	var any interface{} = x
//	fmt.Println(any)
//
//	g := Gopher{"Go"}
//	var c coder = g
//	fmt.Println(c)
//}
//
//type coder interface {
//	code()
//	debug()
//}
//
//type Gopher struct {
//	language string
//}
//
//func (p Gopher) code() {
//	fmt.Printf("I am coding %s language\n", p.language)
//}
//
//func (p Gopher) debug() {
//	fmt.Printf("I am debuging %s language\n", p.language)
//}

package main

import (
	"fmt"
	"time"
)

func Count(ch chan int) {
	ch <- 1
	fmt.Println("Counting")
}
func main02() {
	chs := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go Count(chs[i])
	}
	for _, ch := range chs {
		<-ch
	}
	time.Sleep(100 * time.Millisecond)

	ch := make(chan int, 1)
	for {
		select {
		case ch <- 0:
		case ch <- 1:
		case ch <- 2:
		case ch <- 3:
		}
		i := <-ch
		fmt.Println("Value received:", i)
		time.Sleep(100 * time.Millisecond)
	}
}
