package main

import (
	"fmt"
	"time"
)

func main() {
	for n := range gen() {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
	time.Sleep(100 * time.Second)
}

func gen() <-chan int {
	ch := make(chan int)

	go func() {
		n := 0
		for {
			ch <- n
			n++
		}
	}()

	return ch
}
