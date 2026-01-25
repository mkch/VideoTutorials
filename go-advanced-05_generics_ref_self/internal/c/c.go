package c

type Node[P any] interface {
	comparable
	Parent() P
}

func FindNode[T Node[T]](
	root T,
	eq func(T) bool,
) T {
	var zero T
	if root == zero {
		return zero
	}
	if eq(root) {
		return root
	}
	return FindNode(root.Parent(), eq)
}
