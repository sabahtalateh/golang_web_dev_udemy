package main

import (
	"fmt"
	"time"
)

func worker(ch chan bool) {
	fmt.Print("start")
	fmt.Print(".")
	time.Sleep(200 * time.Millisecond)
	fmt.Print(".")
	time.Sleep(200 * time.Millisecond)
	fmt.Print(".")
	time.Sleep(200 * time.Millisecond)
	fmt.Print("done\n")
	ch <- true
}

func main() {
	ch := make(chan bool)
	go worker(ch)
	<-ch
}
