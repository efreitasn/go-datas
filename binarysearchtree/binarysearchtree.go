// Package binarysearchtree provides functions to create/work with generic binary search trees.
package binarysearchtree

import (
	"golang.org/x/exp/constraints"
)

// BinarySearchTree is a binary search tree of an ordered type T.
type BinarySearchTree[T constraints.Ordered] struct {
	root *Node[T]
	size int
}

// New creates a binary search tree of a constraints.Ordered type T.
func New[T constraints.Ordered]() *BinarySearchTree[T] {
	return &BinarySearchTree[T]{}
}

// Insert adds a value to the binary search tree.
func (bst *BinarySearchTree[T]) Insert(v T) {
	newRoot, inserted := insertRecursive(nil, bst.root, v)

	bst.root = newRoot

	if inserted {
		bst.size++
	}
}

func insertRecursive[T constraints.Ordered](previous *Node[T], current *Node[T], v T) (node *Node[T], ok bool) {
	if current == nil {
		return &Node[T]{
			value:  v,
			parent: previous,
		}, true
	}

	if v > current.Value() {
		current.right, ok = insertRecursive(current, current.right, v)
	} else if v < current.Value() {
		current.left, ok = insertRecursive(current, current.left, v)
	}

	return current, ok
}

// Remove removes a node from the binary search tree.
func (bst *BinarySearchTree[T]) Remove(n *Node[T]) {
	if n.Left() == nil && n.Right() == nil {
		p := n.Parent()

		if p.Left() == n {
			p.left = nil
		} else {
			p.right = nil
		}
	} else if n.Left() != nil && n.Right() != nil {
		successor := n.Right()

		for successor.Left() != nil {
			successor = successor.Left()
		}

		n.value = successor.Value()

		bst.Remove(successor)

		return
	} else if n.Left() != nil {
		n.value = n.Left().Value()
		n.left = nil
	} else {
		n.value = n.Right().Value()
		n.right = nil
	}
}

// NodeHeight returns the number of edges from a node to its deepest descendent (the height of a node).
func (bst *BinarySearchTree[T]) NodeHeight(n *Node[T]) int {
	return nodeHeightRecursive(n)
}

func nodeHeightRecursive[T constraints.Ordered](n *Node[T]) int {
	if n.Left() == nil && n.Right() == nil {
		return 0
	}

	leftHeight := nodeHeightRecursive(n.Left())
	rightHeight := nodeHeightRecursive(n.Right())

	if leftHeight > rightHeight {
		return leftHeight + 1
	}

	return rightHeight + 1
}

// NodeDepth returns the number of edges from the root of the tree to a node (the depth of a node).
func (bst *BinarySearchTree[T]) NodeDepth(n *Node[T]) int {
	return nodeDepthRecursive(n)
}

func nodeDepthRecursive[T constraints.Ordered](n *Node[T]) int {
	if n.Parent() == nil {
		return 0
	}

	return nodeDepthRecursive(n.Parent()) + 1
}

// Find returns the node of the binary search value that has a specific value.
func (bst *BinarySearchTree[T]) Find(v T) (n *Node[T], found bool) {
	n = bst.Root()

	for n != nil {
		nV := n.Value()

		if v < nV {
			n = n.Left()
		} else if v > nV {
			n = n.Right()
		} else {
			return n, true
		}
	}

	return n, n == nil
}

// Root returns the root of the binary search tree.
func (bst *BinarySearchTree[T]) Root() *Node[T] {
	return bst.root
}

// Height returns the number of edges between the root of the tree and its deepest descendent, i.e. the height of the root.
func (bst *BinarySearchTree[T]) Height() int {
	return bst.NodeHeight(bst.Root())
}

// Size returns the number of nodes in the binary search tree.
func (bst *BinarySearchTree[T]) Size() int {
	return bst.size
}
