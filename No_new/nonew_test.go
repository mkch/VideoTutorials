package nonew

import "testing"

func BenchmarkNew(b *testing.B) {
	var v S
	for b.Loop() {
		F(new(v))
	}
}

func BenchmarkAddr(b *testing.B) {
	var v S
	for b.Loop() {
		F(&v)
	}
}

func TestF2New(t *testing.T) {
	var v = S{N1: 100}
	F2(new(v))
	if v.N1 != 200 {
		t.Errorf("Expected %v, but got %v",
			200, v.N1)
	}
}
