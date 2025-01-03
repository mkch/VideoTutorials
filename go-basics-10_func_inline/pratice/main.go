package main

import (
	"fmt"
	"slices"
)

func main() {
	tag1 := AddListener(func() { fmt.Println("1") })
	AddListener(func() { fmt.Println("2") })
	CallListeners()
	fmt.Println()
	RemoveListener(tag1)
	CallListeners()
}

type Listener func()

type Tag struct{ l *Listener }

var listeners []*Listener

func CallListeners() {
	for _, l := range listeners {
		(*l)()
	}
}

func AddListener(l Listener) Tag {
	listeners = append(listeners, &l)
	return Tag{&l}
}

func RemoveListener(tag Tag) {
	listeners = slices.DeleteFunc(listeners,
		func(d *Listener) bool { return d == tag.l })
}
