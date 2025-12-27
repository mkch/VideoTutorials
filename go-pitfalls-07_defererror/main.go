package main

import (
	"bufio"
	"os"

	"github.com/mkch/gg"
)

func GenerateFile() (err error) {
	// Open the file for writing
	f, err := os.OpenFile("a.txt", os.O_CREATE|os.O_TRUNC, 0644) // lacks os.O_WRONLY!!
	if err != nil {
		return
	}
	// Ensure the file is closed when the function exits
	defer gg.CollectError(f.Close, &err)

	// Wrap w in a buffered writer
	w := bufio.NewWriter(f)
	// Ensure the buffer is flushed when the function exits
	defer gg.CollectError(w.Flush, &err)

	_, err = w.Write([]byte("Hello, World!"))
	if err != nil {
		return
	}
	return
}

func main() {
	if err := GenerateFile(); err != nil {
		panic(err) // will panic with `bad file descriptor` error
	}
}
