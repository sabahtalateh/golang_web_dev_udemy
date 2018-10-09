package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func chanOps(ch chan string) {
	n := 0
	for {
		select {
		case ch <- strconv.Itoa(n):
			sleepMs := rand.Intn(100)
			time.Sleep(time.Duration(sleepMs) * time.Millisecond)
			n++

		case msg := <-ch:
			if msg == "exit" {
				return
			}
		}
	}
}

func main() {
	ch := make(chan string)
	go chanOps(ch)
	for {
		msg := <-ch
		fmt.Println(msg)
		if msg == "10" {
			ch <- "exit"
			return
		}
	}
}
