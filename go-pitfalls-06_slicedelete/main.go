package main

import (
	"fmt"
	"slices"
)

func main() {

	s := []int{2, 3, 4, 5, 6}

	s1 := slices.DeleteFunc(slices.Clone(s),
		func(n int) bool { return n > 2 && n%2 == 0 })

	s2 := slices.DeleteFunc(slices.Clone(s),
		func(n int) bool { return n > 2 && n%3 == 0 })

	fmt.Println(s1, s2, s)
}
