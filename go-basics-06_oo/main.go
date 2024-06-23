package main

import "oo/sum"

func PrimarySum(from, to int) int {
	var n int
	for i := from; i <= to; i++ {
		n += i
	}
	return n
}

func FormulaSum(from, to int) int {
	return (from + to) * (to - from + 1) / 2
}

func main() {
	sum.Sum = PrimarySum
	sum.Print(1, 100)

	sum.Sum = FormulaSum
	sum.Print(1, 100)
}
