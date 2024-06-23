package main

import "oo/sum"

type PrimaryCalc struct{}

func (PrimaryCalc) Sum(from, to int) int {
	var n int
	for i := from; i <= to; i++ {
		n += i
	}
	return n
}

type FormulaCalc struct{}

func (FormulaCalc) Sum(from, to int) int {
	return (from + to) * (to - from + 1) / 2
}

func Main2() {
	sum.Sum2 = PrimaryCalc{}
	sum.Print2(1, 100)

	sum.Sum2 = FormulaCalc{}
	sum.Print2(1, 100)
}
