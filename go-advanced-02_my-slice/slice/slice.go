package slice

import (
	"bytes"
	"fmt"
)

//go:generate go run genalloc/genalloc.go

type Slice[T any] struct {
	data *T
	len  int
	cap  int
}

// Make[T](len, -1) -> make([]T, len)
// Make[T](len, cap) -> make([]T, len, cap)
func Make[T any](len, cap int) Slice[T] {
	if cap == -1 {
		cap = len
	}
	if len < 0 {
		panic("wrong len")
	}
	if cap < len {
		panic("wrong cap")
	}

	return Slice[T]{
		data: alloc[T](cap),
		len:  len,
		cap:  cap,
	}
}

// Len(s) -> len(s)
func Len[T any](s Slice[T]) int {
	return s.len
}

// Cap(s) -> cap(s)
func Cap[T any](s Slice[T]) int {
	return s.cap
}

// Append(s, elems...) -> append(s, elems...)
func Append[T any](s Slice[T], elems ...T) Slice[T] {
	var result = s
	result.len = len(elems) + s.len
	if result.len <= s.cap { // enough cap
		for i, v := range elems {
			*ptrAdd[T](s.data, s.len+i) = v
		}
	} else { // need more cap
		result.cap = (result.len<<1 + 7) &^ 7
		result.data = alloc[T](result.cap)
		// copy old data
		for i := 0; i < s.len; i++ {
			*ptrAdd[T](result.data, i) =
				*ptrAdd[T](s.data, i)
		}
		// copy new data
		for i, v := range elems {
			*ptrAdd[T](result.data, s.len+i) = v
		}
	}
	return result
}

// s.Get(i) -> s[i]
func (s Slice[T]) Get(i int) T {
	if i < 0 || i > s.len {
		panic("index out of range")
	}
	return *ptrAdd(s.data, i)
}

// s.Set(i, v) -> s[i] = v
func (s Slice[T]) Set(i int, v T) {
	if i < 0 || i > s.len {
		panic("index out of range")
	}
	*ptrAdd(s.data, i) = v
}

// s.Slice(-1, -1, -1) -> s[:]
// s.Slice(-1, high, -1) -> s[:high]
// s.Slice(-1, high, max) -> s[:high:max]
// s.Slice(low, high, max) -> s[low:high:max]
func (s Slice[T]) Slice(low, high, max int) Slice[T] {
	if low == -1 {
		low = 0
	}
	if high == -1 {
		high = s.len
	}
	if max == -1 {
		max = s.cap
	}
	if low < 0 || high < low || max < high || s.cap < max {
		panic("index out of range")
	}
	return Slice[T]{
		data: ptrAdd[T](s.data, low),
		len:  high - low,
		cap:  max - low,
	}
}

func (s Slice[T]) String() string {
	buf := bytes.NewBufferString("[")
	for i := 0; i < s.len; i++ {
		fmt.Fprint(buf, s.Get(i))
		if i < s.len-1 {
			buf.WriteRune(' ')
		}
	}
	buf.WriteRune(']')
	return buf.String()
}
