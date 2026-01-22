package main

import "fmt"

type Closure1 struct {
	n *int
}

func (c *Closure1) Call() (id int) {
	id = *c.n
	(*c.n)++
	return
}

type Closure10 struct {
	n *int
}

func (c *Closure10) Call() (id int) {
	id = *c.n
	(*c.n) += 10
	return
}

func NewIDGenerators2(start int) (
	*Closure1,
	*Closure10,
) {
	n := start
	return &Closure1{&n}, &Closure10{&n}
}

func main2() {
	idGen1, idGen10 :=
		NewIDGenerators2(1)
	fmt.Println(idGen1.Call()) // 1
	fmt.Println(idGen1.Call()) // 2

	fmt.Println(idGen10.Call()) // 3?
	fmt.Println(idGen10.Call()) // 13?
}
