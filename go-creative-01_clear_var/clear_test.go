package helper

import (
	"testing"
)

func TestClear(t *testing.T) {
	var n int = 1
	Clear(&n)
	if n != 0 {
		t.Errorf("expected 0, got %d", n)
	}
	var s = struct {
		I int
		S string
	}{I: 1, S: "test"}
	Clear(&s)
	if s.I != 0 || s.S != "" {
		t.Errorf("expected zero value, got %+v", s)
	}
}

const SMALL = 1024
const LARGE = 1024 * 1024

func BenchmarkClearSmall(b *testing.B) {
	var v [SMALL]byte
	for b.Loop() {
		Clear(&v)
	}
}

func BenchmarkClearLarge(b *testing.B) {
	var v [LARGE]byte
	for b.Loop() {
		Clear(&v)
	}
}

func TestClearSafe(t *testing.T) {
	var n int = 1
	clearSafe(&n)
	if n != 0 {
		t.Errorf("expected 0, got %d", n)
	}
	var s = struct {
		I int
		S string
	}{I: 1, S: "test"}
	clearSafe(&s)
	if s.I != 0 || s.S != "" {
		t.Errorf("expected zero value, got %+v", s)
	}
}

func BenchmarkClearSafeSmall(b *testing.B) {
	var v [SMALL]byte
	for b.Loop() {
		clearSafe(&v)
	}
}

func BenchmarkClearSafeLarge(b *testing.B) {
	var v [LARGE]byte
	for b.Loop() {
		clearSafe(&v)
	}
}
