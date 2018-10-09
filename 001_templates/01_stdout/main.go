package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {

	name := os.Args[1]
	str := `
		<html>
		<body>
		<h1>Hello ` + name + `</h1>
		</body>
		</html>
	`
	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatal("error creating file", err)
	}
	defer nf.Close()

	io.Copy(nf, strings.NewReader(str))
}
