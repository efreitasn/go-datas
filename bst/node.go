package bst

import "golang.org/x/exp/constraints"

// Node is a node of a binary search tree.
type Node[T constraints.Ordered] struct {
	value         T
	left          *Node[T]
	right         *Node[T]
	parent        *Node[T]
	balanceFactor int8
}

// Value returns the value of the node.
func (n *Node[T]) Value() T {
	return n.value
}

// Left returns the left child of the node.
func (n *Node[T]) Left() *Node[T] {
	return n.left
}

// Right returns the right child of the node.
func (n *Node[T]) Right() *Node[T] {
	return n.right
}

// Parent returns the parent of the node.
func (n *Node[T]) Parent() *Node[T] {
	return n.parent
}

// BalanceFactor returns the balance factor of the node.
func (n *Node[T]) BalanceFactor() int8 {
	return n.balanceFactor
}
