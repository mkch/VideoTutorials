package sum

import "fmt"

type Calculator interface {
	Sum(from, to int) int
}

var Sum2 Calculator

func Print2(from, to int) {
	n := Sum2.Sum(from, to)
	fmt.Printf("The sum is %v\n", n)
}
