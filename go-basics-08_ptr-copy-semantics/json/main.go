package main

import (
	"encoding/json"
	"fmt"
)

const data = `{
    "F1":1
}`

type S struct {
	F1 *int
}

func main() {
	var p *S = &S{} // p.F1 == nil
	err := json.Unmarshal([]byte(data), p)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(p, *p.F1)
}
