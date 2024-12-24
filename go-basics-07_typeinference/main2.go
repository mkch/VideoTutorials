package main

import "io"

func F[T any](a T, b T)           {}
func F1(a io.Reader, b io.Reader) {}

func main() {
	var r io.Reader
	var rc io.ReadCloser
	F(r, rc) // T: io.Reader
	F[io.Reader](r, rc)
	F1(r, rc)
}
