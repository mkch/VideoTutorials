package main

import (
	"fmt"
	"myslice/slice"
)

func main() {
	// last := []int{2, 0, 2, 4}
	last := slice.Append(
		slice.Make[int](0, 3),
		2, 0, 2, 4)

	// this := append(last[:3], 5)
	this := slice.Append(
		last.Slice(0, 3, -1),
		5)

	fmt.Println(last, this)
}
