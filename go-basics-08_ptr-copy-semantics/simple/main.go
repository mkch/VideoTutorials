package main

import (
	"fmt"
)

const data = `{
	"F1":1
}`

type S struct {
	F1 *int
}

func main() {
	var p *S = &S{}
	F(p)
	fmt.Println(p, *p.F1)
}

func F(p1 *S) {
	// p1 = &S{}
	var n = 100
	p1.F1 = &n
}
