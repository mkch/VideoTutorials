package slice_test

import (
	"myslice/slice"
	"testing"
)

func TestMake(t *testing.T) {
	s := slice.Make[byte](1, -1)
	if slice.Len(s) != 1 {
		t.Fail()
	}
	if slice.Cap(s) != 1 {
		t.Fail()
	}
	s = slice.Make[byte](2, 3)
	if slice.Len(s) != 2 {
		t.Fatal()
	}
	if slice.Cap(s) != 3 {
		t.Fatal()
	}
}

func TestGetSet(t *testing.T) {
	var s slice.Slice[int]
	if slice.Len(s) != 0 || slice.Cap(s) != 0 {
		t.Fatal()
	}
	s = slice.Make[int](3, -1)
	s.Set(0, 1)
	s.Set(1, 2)
	s.Set(2, 3)
	if s.Get(0) != 1 {
		t.Fatal()
	}
	if s.Get(1) != 2 {
		t.Fatal()
	}
	if s.Get(2) != 3 {
		t.Fatal()
	}
}

func TestAppend(t *testing.T) {
	var s slice.Slice[int64]
	s = slice.Append(s, 10, 20, 30, 40)
	if slice.Len(s) != 4 {
		t.Fatal()
	}
	if slice.Cap(s) != 8 {
		t.Fatal()
	}
	if s.Get(0) != 10 || s.Get(1) != 20 ||
		s.Get(2) != 30 || s.Get(3) != 40 {
		t.Fatal()
	}

	s = slice.Make[int64](60, 63)
	s = slice.Append(s, 1, 2, 3, 4)
	if slice.Len(s) != 64 {
		t.Fatal()
	}
	if slice.Cap(s) != 128 {
		t.Fatal()
	}
	if s.Get(60) != 1 {
		t.Fatal()
	}
	if s.Get(61) != 2 {
		t.Fatal()
	}
	if s.Get(62) != 3 {
		t.Fatal()
	}
	if s.Get(63) != 4 {
		t.Fatal()
	}

	s = slice.Append(s, 5)
	if slice.Len(s) != 65 {
		t.Fatal()
	}
	if slice.Cap(s) != 128 {
		t.Fatal()
	}
	if s.Get(64) != 5 {
		t.Fatal()
	}
}

func TestString(t *testing.T) {
	var s slice.Slice[string]
	if str := s.String(); str != "[]" {
		t.Fatal(str)
	}
	s = slice.Append(slice.Slice[string]{}, "abc", "def", "ghi")
	if str := s.String(); str != "[abc def ghi]" {
		t.Fatal(str)
	}
}

func TestSlice(t *testing.T) {
	s := slice.Append(slice.Slice[byte]{}, 1, 2, 3, 4, 5)
	s2 := s.Slice(1, 3, -1)
	s2.Set(1, 30)
	if str := s.String(); str != "[1 2 30 4 5]" {
		t.Fatal(str)
	}
	if str := s2.String(); str != "[2 30]" {
		t.Fatal(str)
	}
	s2 = slice.Append(s2, 40, 50)
	if str := s2.String(); str != "[2 30 40 50]" {
		t.Fatal(str)
	}
	if str := s.String(); str != "[1 2 30 40 50]" {
		t.Fatal(str)
	}
	s2 = s.Slice(0,2,2)
	s2 = slice.Append(s2, 99)
	if str := s2.String(); str != "[1 2 99]" {
		t.Fatal(str)
	}
	if str := s.String(); str != "[1 2 30 40 50]" {
		t.Fatal(str)
	}
}
