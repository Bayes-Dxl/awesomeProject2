package main

import (
	"flag"
	"fmt"
	"time"
)

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

func main() {
	var c ByteCounter
	_, err := c.Write([]byte("hello"))
	if err != nil {
		return
	}
	fmt.Println(c)

	c = 0
	var name = "Dolly"
	fprintf, err := fmt.Fprintf(&c, "hello, %s", name)
	if err != nil {
		return
	}
	fmt.Println(fprintf)
	fmt.Println(c)

	var period = flag.Duration("period", 1*time.Second, "sleep period")
	flag.Parse()
	fmt.Printf("Sleep for %v...", *period)
	time.Sleep(*period)
	fmt.Println()
}
