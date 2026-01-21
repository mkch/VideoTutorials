package closure

import "fmt"

type Closure struct {
	n *int
}

func (c *Closure) Call() int {
	var base = 100
	(*c.n)++
	return base + *c.n
}

func NewIDGenerator2() *Closure {
	n := 1
	return &Closure{&n}
}

func f2() {
	idGen := NewIDGenerator2()
	fmt.Println(idGen.Call()) // 102
	fmt.Println(idGen.Call()) // 103
	fmt.Println(idGen.Call()) // 104
}
