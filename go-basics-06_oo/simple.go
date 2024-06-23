package main

import "os"

var a int

func SetA(v int) {
	a = v
}

func A() int {
	return a
}

type S struct {
	a int
}

func (s *S) SetA(v int) {
	s.a = v
}

func (s *S) A() int {
	return s.a
}

type F struct {
	*os.File
}

func Test() {
	var f F
	// f.File = ...
	f.Stat()
}
