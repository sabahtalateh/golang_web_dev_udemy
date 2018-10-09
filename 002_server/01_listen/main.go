package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	li, err := net.Listen("tcp", ":8999")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	err := conn.SetDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		log.Println("CONNECTION TIMEOUT")
	}
	scaner := bufio.NewScanner(conn)
	scaner.Split(bufio.ScanLines)
	for scaner.Scan() {
		ln := scaner.Text()
		fmt.Println(ln)
		fmt.Fprintf(conn, "I heard that [%s]\n", ln)
	}
	defer conn.Close()
	fmt.Println("***CODE GOT HERE***")
}