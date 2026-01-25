package a

import "lib/internal/c"

type Node struct {
	parent *Node
	Value  int
}

func (n *Node) Parent() *Node {
	return n.parent
}

// func FindNode(
// 	root *Node,
// 	eq func(*Node) bool,
// ) *Node {
// 	if root == nil {
// 		return nil
// 	}
// 	if eq(root) {
// 		return root
// 	}
// 	return FindNode(root.Parent(), eq)
// }

func API1() {
	var root *Node // from somewhere
	_ = c.FindNode(root, func(n *Node) bool {
		return n.Value == 1
	})
}
