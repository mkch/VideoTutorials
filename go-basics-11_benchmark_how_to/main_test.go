package main_test

import (
	"testing"
)

type Element [4]int

var a [100]Element

var s = a[:]

var v = Element{1}

func loop1() {
	for i, e := range s {
		if e == v {
			panic(i)
		}
	}
}

func loop2() {
	for i := range s {
		if s[i] == v {
			panic(i)
		}
	}
}

func BenchmarkLoop1(b *testing.B) {
	for range b.N {
		loop1()
	}
}

func BenchmarkLoop2(b *testing.B) {
	for range b.N {
		loop2()
	}
}
