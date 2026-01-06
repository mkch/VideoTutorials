package main

import "os"

func checkError(f func() error) {
	if err := f(); err != nil {
		panic(err)
	}
}

func main() {
	f, err := os.OpenFile(
		"a.txt",
		os.O_CREATE|os.O_WRONLY|os.O_TRUNC,
		0644)

	if err != nil {
		panic(err)
	}

	// defer f.Close()
	defer checkError(f.Close)

	_, err = f.WriteString("hello, world\n")

	if err != nil {
		panic(err)
	}
}
