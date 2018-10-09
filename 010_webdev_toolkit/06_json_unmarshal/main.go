package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type thumbnail struct {
	URL    string
	Width  int
	Height int
}

type img struct {
	Width     int
	Height    int
	Title     string
	Thumbnail thumbnail
	Animated  bool
	IDs       []int
}

func main() {
	var data img
	var data2 interface{}
	s := `{"Width":800,"Height":600,"Title":"Jopa","Thumbnail":{"URL":"https://www.bitchinseatstore.net/v/vspfiles/photos/09-15vrscf-18.jpg","Width":200,"Height":300},"Animated":false,"IDs":[1128939,9238492,7]}`

	err := json.Unmarshal([]byte(s), &data)
	if err != nil {
		log.Println(err)
	}
	err = json.Unmarshal([]byte(s), &data2)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("%+v\n", data)
	fmt.Printf("%+v\n", data2)
}
