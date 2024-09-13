package bst

import (
	"iter"

	"golang.org/x/exp/constraints"
)

// BST is a binary search tree of an ordered type T.
type BST[T constraints.Ordered] struct {
	root *Node[T]
	size int
}

// NewBST creates a binary search tree of an ordered type T.
func NewBST[T constraints.Ordered]() *BST[T] {
	return &BST[T]{}
}

// Root returns the root of the BST.
func (bst *BST[T]) Root() *Node[T] {
	return bst.root
}

// Size returns the number of nodes in the BST.
func (bst *BST[T]) Size() int {
	return bst.size
}

// Insert adds a value to the BST.
func (bst *BST[T]) Insert(v T) {
	newRoot, inserted := bst.insertRecursive(nil, bst.root, v)

	bst.root = newRoot

	if inserted {
		bst.size++
	}
}

// Find returns the node of the binary search value that has a specific value.
func (bst *BST[T]) Find(v T) (n *Node[T], found bool) {
	n = bst.Root()

	for n != nil {
		switch {
		case v < n.value:
			n = n.left

		case v > n.value:
			n = n.right

		default:
			return n, true
		}
	}

	return n, false
}

// Traverse returns an iterator over all nodes in the BST in the given order.
func (bst *BST[T]) Traverse(order TraverseOrder) iter.Seq[*Node[T]] {
	return traverse(order, bst.root)
}

// Remove removes a node from the BST.
func (bst *BST[T]) Remove(n *Node[T]) {
	if n == bst.root && bst.size == 1 {
		bst.root = nil
		bst.size = 0
		return
	}

	var replacement *Node[T]

	switch {
	case n.left != nil && n.right != nil:
		r := n.left
		for r.right != nil {
			r = r.right
		}

		bst.Remove(r)

		r.left = n.left
		if r.left != nil {
			r.left.parent = r
		}

		r.right = n.right
		r.right.parent = r

		bst.replaceNode(n, r)
		return

	case n.right != nil:
		replacement = n.right

	case n.left != nil:
		replacement = n.left
	}

	bst.replaceNode(n, replacement)

	bst.size--
}

func (bst *BST[T]) insertRecursive(previous *Node[T], current *Node[T], v T) (node *Node[T], ok bool) {
	if current == nil {
		return &Node[T]{
			value:  v,
			parent: previous,
		}, true
	}

	if v > current.value {
		current.right, ok = bst.insertRecursive(current, current.right, v)
	} else if v < current.value {
		current.left, ok = bst.insertRecursive(current, current.left, v)
	}

	return current, ok
}

func (bst *BST[T]) replaceNode(n, replacement *Node[T]) {
	switch {
	case n.parent == nil:
		bst.root = replacement

	case n.parent.left == n:
		n.parent.left = replacement

	default:
		n.parent.right = replacement
	}

	if replacement != nil {
		replacement.parent = n.parent
	}
}
