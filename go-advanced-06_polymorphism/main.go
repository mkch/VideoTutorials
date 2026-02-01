package main

import (
	"fmt"
	"main/gui"
	"slices"
)

func main() {
	w := &gui.ColorWindow{Background: "red"}
	b := &gui.Button{}

	gui.AddChild(w, b)

	fmt.Printf("Button's parent: %T\n",
		b.Parent())
	fmt.Printf("Parent's background: %s\n",
		b.Parent().(*gui.ColorWindow).Background)
	fmt.Printf("Windows's child: %T\n",
		slices.Collect(w.Children())[0])

}
