package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}
	defer li.Close()
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}

		scaner := bufio.NewScanner(conn)
		for scaner.Scan() {
			fmt.Println(scaner.Text())
		}

		io.WriteString(conn, "\nHello from TCP server")
		conn.Close()
		fmt.Println("got here")
	}
}
