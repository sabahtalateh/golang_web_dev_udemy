package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

func main() {
	s := "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book."

	encodeStd := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	s64 := base64.NewEncoding(encodeStd).EncodeToString([]byte(s))

	fmt.Println(len(s))
	fmt.Println(len(s64))
	fmt.Println(len(encodeStd))
	fmt.Println(s)
	fmt.Println(s64)

	s64 = base64.StdEncoding.EncodeToString([]byte(s))
	fmt.Println(s64)

	bs, err := base64.StdEncoding.DecodeString(s64)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bs))
}
