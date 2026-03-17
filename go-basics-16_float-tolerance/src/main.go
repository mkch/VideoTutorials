package main

import (
	"fmt"
	"math"
)

func main() {
	bits()
	ulp()
}

func absIsClose(a, b, tol float64) bool {
	diff := math.Abs(a - b)
	return diff <= tol
}

func relIsClose(a, b, tol float64) bool {
	diff := math.Abs(a - b)
	return diff <=
		tol*max(math.Abs(a), math.Abs(b))
}

func isClose(a, b, absTol float64, relTol float64) bool {
	diff := math.Abs(a - b)
	return diff <= absTol ||
		diff <= relTol*max(math.Abs(a), math.Abs(b))
}

func bits() {
	b := math.Float64bits(-26.375)
	fmt.Printf("%0x\n%064b\n", b, b)

	f := math.Float64frombits(0xc03a600000000000)
	fmt.Println(f)
}

func ulp() {
	a := -26.375
	b := math.Nextafter(a, 0)
	diff := b - a
	fmt.Printf("a-b: %g\n", diff)

	ulp := math.Pow(2, 4-52)
	fmt.Printf("ULP: %g\n", ulp)
}
