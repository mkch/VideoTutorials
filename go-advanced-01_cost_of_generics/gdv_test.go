package gdv

import (
	"bytes"
	"io"
	"testing"
)

func Write(w io.Writer) {
	w.Write([]byte{1})
}

func WriteG[T io.Writer](w T) {
	w.Write([]byte{1})
}

func BenchmarkWrite(b *testing.B) {
	var w bytes.Buffer
	for range b.N {
		Write(&w)
	}
}

func BenchmarkWriteG(b *testing.B) {
	var w bytes.Buffer
	for range b.N {
		WriteG(&w)
	}
}

type Number interface {
	int | uint | float64
}

func Add(n float64) float64 {
	return n + 1
}

func AddG[T Number](n T) T {
	return n + 1
}

func BenchmarkAdd(b *testing.B) {
	var n float64 = 1
	for range b.N {
		n = Add(n)
	}
}

func BenchmarkAddg(b *testing.B) {
	var n float64 = 1
	for range b.N {
		n = AddG(n)
	}
}
