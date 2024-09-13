package bst

import (
	"fmt"
	"iter"

	"golang.org/x/exp/constraints"
)

// TraverseOrder is the order in which the nodes of a binary search tree are traversed.
type TraverseOrder int

// List of valid traverse orders.
const (
	NLR TraverseOrder = iota + 1
	LNR
	LRN

	PreOrder  = NLR
	InOrder   = LNR
	PostOrder = LRN
)

func traverse[T constraints.Ordered](order TraverseOrder, root *Node[T]) iter.Seq[*Node[T]] {
	var recursiveFunc func(*Node[T], func(*Node[T]) bool) bool

	switch order {
	case NLR:
		recursiveFunc = traverseNLRRecursive

	case LNR:
		recursiveFunc = traverseLNRRecursive

	case LRN:
		recursiveFunc = traverseLRNRecursive

	default:
		panic(fmt.Sprintf("invalid traverse order: %v", order))
	}

	return func(yield func(*Node[T]) bool) {
		if root == nil {
			return
		}

		recursiveFunc(root, yield)
	}
}

func traverseNLRRecursive[T constraints.Ordered](n *Node[T], yield func(*Node[T]) bool) bool {
	if n == nil {
		return true
	}

	if !yield(n) {
		return false
	}

	if n.left != nil {
		if !traverseNLRRecursive(n.left, yield) {
			return false
		}
	}

	if n.right != nil {
		if !traverseNLRRecursive(n.right, yield) {
			return false
		}
	}

	return true
}

func traverseLNRRecursive[T constraints.Ordered](n *Node[T], yield func(*Node[T]) bool) bool {
	if n == nil {
		return true
	}

	if n.left != nil {
		if !traverseLNRRecursive(n.left, yield) {
			return false
		}
	}

	if !yield(n) {
		return false
	}

	if n.right != nil {
		if !traverseLNRRecursive(n.right, yield) {
			return false
		}
	}

	return true
}

func traverseLRNRecursive[T constraints.Ordered](n *Node[T], yield func(*Node[T]) bool) bool {
	if n == nil {
		return true
	}

	if n.left != nil {
		if !traverseLRNRecursive(n.left, yield) {
			return false
		}
	}

	if n.right != nil {
		if !traverseLRNRecursive(n.right, yield) {
			return false
		}
	}

	return yield(n)
}
