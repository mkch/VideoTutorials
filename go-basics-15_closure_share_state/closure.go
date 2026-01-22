package main

import "fmt"

// NewIDGenerators returns two id generators.
func NewIDGenerators(start int) (
	x1 func() int,
	x10 func() int,
) {
	n := start
	x1 = func() (id int) {
		id = n
		n++
		return
	}
	x10 = func() (id int) {
		id = n
		n += 10
		return
	}
	return
}

func main() {
	idGen1, idGen10 :=
		NewIDGenerators(1)
	fmt.Println(idGen1()) // 1
	fmt.Println(idGen1()) // 2

	fmt.Println(idGen10()) // 1
	fmt.Println(idGen10()) // 11
}
