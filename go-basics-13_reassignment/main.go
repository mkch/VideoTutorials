package main

import (
	"fmt"
	"io"
)

func F() (err error) {
	n, err := 1, io.EOF
	_ = n
	return
}

func main() {
	fmt.Println(F())
}
