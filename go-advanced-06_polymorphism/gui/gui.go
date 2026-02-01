package gui

import (
	"iter"
	"slices"
)

// Control represents a UI control element.
type Control interface {
	Parent() ControlParent
	setParent(p ControlParent)
}

// ControlParent represents a parent that
// contains controls.
type ControlParent interface {
	Children() iter.Seq[Control]
	appendChild(c Control)
}

// Button is a concrete [Control].
type Button struct {
	parent ControlParent
}

func (b *Button) Parent() ControlParent {
	return b.parent
}

func (b *Button) setParent(p ControlParent) {
	b.parent = p
}

// Window is a concrete [ControlParent].
type Window struct {
	children []Control
}

func (w *Window) Children() iter.Seq[Control] {
	return slices.Values(w.children)
}

func (w *Window) appendChild(c Control) {
	w.children = append(w.children, c)
}

func AddChild(p ControlParent, c Control) {
	// validation ...
	p.appendChild(c)
	c.setParent(p)
}

type ColorWindow struct {
	Window
	Background string
}
