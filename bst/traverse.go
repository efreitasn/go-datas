package bst

import (
	"fmt"

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

// TraverseFunc is the function called when traversing a tree.
type TraverseFunc[T constraints.Ordered] func(n *Node[T]) bool

func traverse[T constraints.Ordered](order TraverseOrder, root *Node[T], fn TraverseFunc[T]) {
	if root == nil {
		return
	}

	switch order {
	case NLR:
		traverseNLRRecursive(root, fn)

	case LNR:
		traverseLNRRecursive(root, fn)

	case LRN:
		traverseLRNRecursive(root, fn)

	default:
		panic(fmt.Sprintf("invalid traverse order: %v", order))
	}
}

func traverseNLRRecursive[T constraints.Ordered](n *Node[T], fn TraverseFunc[T]) bool {
	if n == nil {
		return true
	}

	if !fn(n) {
		return false
	}

	if n.left != nil {
		if !traverseNLRRecursive(n.left, fn) {
			return false
		}
	}

	if n.right != nil {
		if !traverseNLRRecursive(n.right, fn) {
			return false
		}
	}

	return true
}

func traverseLNRRecursive[T constraints.Ordered](n *Node[T], fn TraverseFunc[T]) bool {
	if n == nil {
		return true
	}

	if n.left != nil {
		if !traverseLNRRecursive(n.left, fn) {
			return false
		}
	}

	if !fn(n) {
		return false
	}

	if n.right != nil {
		if !traverseLNRRecursive(n.right, fn) {
			return false
		}
	}

	return true
}

func traverseLRNRecursive[T constraints.Ordered](n *Node[T], fn TraverseFunc[T]) bool {
	if n == nil {
		return true
	}

	if n.left != nil {
		if !traverseLRNRecursive(n.left, fn) {
			return false
		}
	}

	if n.right != nil {
		if !traverseLRNRecursive(n.right, fn) {
			return false
		}
	}

	return fn(n)
}
