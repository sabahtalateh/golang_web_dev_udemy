package main

import (
	"fmt"
)

func ping(ping chan<- string, msg string) {
	ping <- msg
}

func pong(ping <-chan string, pong chan<- string) {
	msg := <-ping
	pong <- msg
}

func main() {
	pingCh := make(chan string)
	pongCh := make(chan string)
	go ping(pingCh, "ping")
	go pong(pingCh, pongCh)
	fmt.Println(<-pongCh)
}
