package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
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
			log.Fatalln(err.Error())
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	// read request
	request(conn)

	// write response
	response(conn)
}

func request(conn net.Conn) {
	scaner := bufio.NewScanner(conn)
	scaner.Split(bufio.ScanLines)
	i := 0
	for scaner.Scan() {
		line := scaner.Text()
		// fmt.Println(line)
		if i == 0 {
			m := strings.Fields(line)[0]
			url := strings.Fields(line)[1]
			fmt.Println("***METHOD", m)
			fmt.Println("***URL", url)
		}
		if line == "" {
			break
		}
		i++
	}
}

func response(conn net.Conn) {
	body := "<html><body><h1>HELLO!!</h1></body></html>"

	fmt.Fprint(conn, "HTTP/1.1 200 OK \r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
