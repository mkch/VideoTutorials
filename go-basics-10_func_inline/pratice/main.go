package main

import (
	"fmt"
	"slices"
)

// func main() {
// 	AddListener(func() { fmt.Println("1") })
// 	AddListener(func() { fmt.Println("2") })
// 	CallListeners()
// }

// type Listener func()

// var listeners []Listener

// func CallListeners() {
// 	for _, l := range listeners {
// 		l()
// 	}
// }

// func AddListener(l Listener) {
// 	listeners = append(listeners, l)
// }

// func RemoveListener(l Listener) {
// 	slices.DeleteFunc(listeners,
// 		func(dl Listener) bool { return dl == l })
// }

func main() {
	tag1 := AddListener(func() { fmt.Println("1") })
	AddListener(func() { fmt.Println("2") })
	CallListeners()
	fmt.Println()
	RemoveListener(tag1)
	CallListeners()
}

type Listener func()

type Tag struct{ *int }

type taggedListener struct {
	tag      Tag
	listener Listener
}

var listeners []taggedListener

func CallListeners() {
	for _, l := range listeners {
		l.listener()
	}
}

func AddListener(l Listener) Tag {
	tag := Tag{new(int)}
	listeners = append(listeners,
		taggedListener{tag, l})
	return tag
}

func RemoveListener(tag Tag) {
	listeners = slices.DeleteFunc(listeners,
		func(d taggedListener) bool { return d.tag == tag })
}
