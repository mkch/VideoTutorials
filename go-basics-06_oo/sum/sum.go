package sum

import "fmt"

var Sum func(from, to int)int 

func Print(from, to int) {
	n := Sum(from, to)
	fmt.Printf("The sum is %v\n", n)
}

