package main

import (
	"encoding/json"
	"fmt"
)

type object struct {
	Code    int    `json:"code"`
	Descrip string `json:"descrip"`
}

func main() {
	var objs []object

	rcvd := `[{"code":11, "descrip":"desc 1"}, {"code":65, "descrip":"desc 2"}]`

	json.Unmarshal([]byte(rcvd), &objs)
	fmt.Println(objs)

	for _, v := range objs {
		fmt.Println(v.Code)
		fmt.Println(v.Descrip)
	}
}
