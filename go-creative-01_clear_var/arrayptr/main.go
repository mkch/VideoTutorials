package main

import "fmt"

func main() {
	var ary [1]int
	pa := &ary
	p0 := &ary[0]
	fmt.Printf("pa=%p\np0=%p\n", pa, p0)
}
