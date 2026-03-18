package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(ulpDistance(1.0, 1.1))
}

func ulpDistanceNaive(a, b float64) (n int) {
	for i := a; i != b; i = math.Nextafter(i, b) {
		n++
	}
	return
}

// ulpDistance calculates the ULP distance between two float64 values.
func ulpDistance(f1, f2 float64) uint64 {
	if math.IsNaN(f1) || math.IsNaN(f2) || math.IsInf(f1, 0) || math.IsInf(f2, 0) {
		panic("can't calculate ULP distance for NaN or Inf")
	}
	u1 := math.Float64bits(f1)
	u2 := math.Float64bits(f2)
	// Same sign, distance is the difference of the bit patterns.
	if int64(u1^u2) >= 0 {
		if u1 > u2 {
			return u1 - u2
		}
		return u2 - u1
	}
	// Different signs, distance is the sum of the distances to zero.
	// +0x8000000000000000 zeros out the sign bit.
	return u1 + u2 + 0x8000000000000000
}
