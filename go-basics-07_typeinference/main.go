package main

import "fmt"

func Add[E int | string](a E, b E) E {
	return a + b
}

func main2() {
	r1 := Add[int](1, 2)
	r2 := Add[string]("a", "b")
	fmt.Println(r1, r2)
	Add(1, 2)
}
