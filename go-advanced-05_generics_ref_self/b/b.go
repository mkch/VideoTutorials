package b

import (
	"lib/a"
	"lib/internal/c"
)

// func findNode(root *a.Node) *a.Node {
// 	if root == nil {
// 		return nil
// 	}
// 	if root.Value == 2 {
// 		return root
// 	}
// 	return findNode(root.Parent())
// }

func API2() {
	var root *a.Node // from somewhere
	_ = c.FindNode(root, func(n *a.Node) bool {
		return n.Value == 2
	})
	// more code here
}
