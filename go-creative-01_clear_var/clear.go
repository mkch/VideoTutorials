package helper

import "unsafe"

// Clear clears v to its zero value.
// Clearing a variable helps prevent
// unintended data retention, especially
// when reusing objects from pools.
func Clear[T any](ptr *T) {
	// Converts ptr to a pointer to a
	// one-element array,
	// then slices it to get a []T of length 1,
	// and finally calls clear on the slice.
	clear((*[1]T)(unsafe.Pointer(ptr))[:])
}

// clearSafe is a safe version of Clear without using unsafe.
//
// The benchmark shows that clearSafe is almost as fast as Clear
// for small sized types(< 500K), but significantly slower for
// large sized types.
// The benchmark also shows that clearSafe do heap allocation for
// large sized types.
//
// The benchmark results:
//
//	go test -run ^BenchmarkClear -bench . -benchmem
//	goos: darwin
//	goarch: amd64
//	pkg: github.com/mkch/gg
//	cpu: Intel(R) Core(TM) i7-7920HQ CPU @ 3.10GHz
//	BenchmarkClearSafeSmall-8       13269362                94.46 ns/op            0 B/op          0 allocs/op
//	BenchmarkClearSmall-8           12717471                87.57 ns/op            0 B/op          0 allocs/op
//	BenchmarkClearSafeLarge-8          18404             67189 ns/op          516105 B/op          1 allocs/op
//	BenchmarkClearLarge-8             117871              9510 ns/op               0 B/op          0 allocs/op
//	PASS
//	ok      github.com/mkch/gg      5.062s
func clearSafe[T any](ptr *T) {
	var zero T
	*ptr = zero
}
