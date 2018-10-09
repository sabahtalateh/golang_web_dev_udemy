package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
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
			fmt.Fprintln(os.Stdout, err)
		}
		go serve(conn)
	}
}

func serve(conn net.Conn) {
	defer conn.Close()

	scaner := bufio.NewScanner(conn)
	i := 0
	for scaner.Scan() {
		ln := scaner.Text()
		fmt.Println(ln)
		if i == 0 {
			words := strings.Fields(ln)
			fmt.Println("***METHOD", words[0])
			fmt.Println("***URL", words[1])
		}
		if ln == "" {
			break
		}
		i++
	}

	body := "<html><h1>WOOOOW.. SO LOW LEVEL</h1></html>"
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	io.WriteString(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn, fmt.Sprintf("Content-Length: %d\r\n", len(body)))
	io.WriteString(conn, "\r\n\r\n")
	io.WriteString(conn, body)
	fmt.Println("Code got here")
}
