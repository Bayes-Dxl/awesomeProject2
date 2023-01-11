package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func WaitForServert(url string) error {
	const timeout = 100 * time.Second
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			fmt.Println("成功！")
			return nil //成功
		}

		log.SetPrefix("Wait: ")
		log.SetFlags(0)
		log.Printf("server not responding (%s); retrying...", err)
		t := (time.Second << uint(tries))
		fmt.Println(uint8(tries))
		p := t / 1000000000
		log.Printf("wait: %d ms", p)
		time.Sleep(t)
	}
	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}

func main() {
	for _, url := range os.Args[1:] {
		if err := WaitForServert(url); err != nil {
			fmt.Fprintf(os.Stderr, " Site is down: %v\n", err)
		}
	}
}
