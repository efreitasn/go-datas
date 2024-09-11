package bst

import (
	"golang.org/x/exp/constraints"
)

// AVL is an AVL tree, which is a self-balancing binary search tree, of an ordered type T.
type AVL[T constraints.Ordered] struct {
	root *Node[T]
	size int
}

// NewAVL creates an AVL tree of an ordered type T.
func NewAVL[T constraints.Ordered]() *AVL[T] {
	return &AVL[T]{}
}

// Root returns the root of the binary search tree.
func (avl *AVL[T]) Root() *Node[T] {
	return avl.root
}

// Size returns the size of the AVL tree.
func (avl *AVL[T]) Size() int {
	return avl.size
}

// Insert adds a value to the AVL tree.
func (avl *AVL[T]) Insert(v T) {
	newRoot, inserted, _ := avl.insertRecursive(nil, avl.root, v)
	avl.root = newRoot

	if inserted {
		avl.size++
	}
}

// Find returns the node of the AVL tree that has a value v.
func (avl *AVL[T]) Find(v T) (n *Node[T], found bool) {
	n = avl.Root()

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

// Traverse traverses the AVL tree in the given order.
func (avl *AVL[T]) Traverse(order TraverseOrder, fn TraverseFunc[T]) {
	traverse(order, avl.root, fn)
}

// Remove removes a node from the AVL tree.
func (avl *AVL[T]) Remove(n *Node[T]) {
	// Deletion
	if n == avl.root && avl.size == 1 {
		avl.root = nil
		avl.size = 0
		return
	}

	var replacement *Node[T]

	switch {
	case n.left != nil && n.right != nil:
		r := n.left
		for r.right != nil {
			r = r.right
		}

		avl.Remove(r)

		r.left = n.left
		if r.left != nil {
			r.left.parent = r
		}

		r.right = n.right
		r.right.parent = r
		r.balanceFactor = n.balanceFactor

		avl.replaceNode(n, r)

		return

	case n.right != nil:
		replacement = n.right

	case n.left != nil:
		replacement = n.left
	}

	if replacement != nil {
		replacement.parent = n.parent
	}

	avl.size--

	// Balancing
parentLoop:
	for x, p := n, n.parent; p != nil; {
		if p.left == x {
			p.balanceFactor++
		} else {
			p.balanceFactor--
		}

		var (
			newP        = p
			grandParent = p.parent
		)

		switch p.balanceFactor {
		case 1, -1:
			break parentLoop

		case 2:
			if p.right.balanceFactor >= 0 {
				newP = avl.rotateLeft(p, true)
			} else {
				newP = avl.rotateRightLeft(p)
			}

		case -2:
			if p.left.balanceFactor <= 0 {
				newP = avl.rotateRight(p, true)
			} else {
				newP = avl.rotateLeftRight(p)
			}
		}

		switch {
		case grandParent == nil:
			avl.root = newP
			break parentLoop

		case grandParent.left == p:
			grandParent.left = newP

		default:
			grandParent.right = newP
		}

		x, p = newP, grandParent
	}

	avl.replaceNode(n, replacement)
}

func (avl *AVL[T]) insertRecursive(previous *Node[T], current *Node[T], v T) (node *Node[T], ok, heightIncrease bool) {
	if current == nil {
		return &Node[T]{
			parent: previous,
			value:  v,
		}, true, true
	}

	// Insertion
	switch {
	case v > current.value:
		current.right, ok, heightIncrease = avl.insertRecursive(current, current.right, v)

		if heightIncrease {
			current.balanceFactor++
		}

	case v < current.value:
		current.left, ok, heightIncrease = avl.insertRecursive(current, current.left, v)

		if heightIncrease {
			current.balanceFactor--
		}

	default:
		return current, false, false
	}

	// Balancing
	switch current.balanceFactor {
	case 0:
		heightIncrease = false

	case 2:
		if current.right.balanceFactor >= 0 {
			current = avl.rotateLeft(current, true)
		} else {
			current = avl.rotateRightLeft(current)
		}

		heightIncrease = false

	case -2:
		if current.left.balanceFactor <= 0 {
			current = avl.rotateRight(current, true)
		} else {
			current = avl.rotateLeftRight(current)
		}

		heightIncrease = false
	}

	return current, ok, heightIncrease
}

func (avl *AVL[T]) rotateLeft(n *Node[T], updateBalanceFactor bool) *Node[T] {
	right := n.right

	n.right = right.left
	if n.right != nil {
		n.right.parent = n
	}

	right.parent = n.parent
	right.left = n
	right.left.parent = right

	if !updateBalanceFactor {
		return right
	}

	if right.balanceFactor == 0 {
		right.balanceFactor = -1
		right.left.balanceFactor = 1
	} else {
		right.balanceFactor = 0
		right.left.balanceFactor = 0
	}

	return right
}

func (avl *AVL[T]) rotateRight(n *Node[T], updateBalanceFactor bool) *Node[T] {
	left := n.left

	n.left = left.right
	if n.left != nil {
		n.left.parent = n
	}

	left.parent = n.parent
	left.right = n
	left.right.parent = left

	if !updateBalanceFactor {
		return left
	}

	if left.balanceFactor == 0 {
		left.balanceFactor = 1
		left.right.balanceFactor = -1
	} else {
		left.balanceFactor = 0
		left.right.balanceFactor = 0
	}

	return left
}

func (avl *AVL[T]) rotateRightLeft(n *Node[T]) *Node[T] {
	n.right = avl.rotateRight(n.right, false)
	n = avl.rotateLeft(n, false)

	switch {
	case n.balanceFactor == 0:
		n.left.balanceFactor = 0
		n.right.balanceFactor = 0

	case n.balanceFactor > 0:
		n.left.balanceFactor = -1
		n.right.balanceFactor = 0

	default:
		n.left.balanceFactor = 0
		n.right.balanceFactor = 1
	}

	n.balanceFactor = 0

	return n
}

func (avl *AVL[T]) rotateLeftRight(n *Node[T]) *Node[T] {
	n.left = avl.rotateLeft(n.left, false)
	n = avl.rotateRight(n, false)

	switch {
	case n.balanceFactor == 0:
		n.left.balanceFactor = 0
		n.right.balanceFactor = 0

	case n.balanceFactor > 0:
		n.left.balanceFactor = -1
		n.right.balanceFactor = 0

	default:
		n.left.balanceFactor = 0
		n.right.balanceFactor = 1
	}

	n.balanceFactor = 0

	return n
}

func (avl *AVL[T]) replaceNode(n, replacement *Node[T]) {
	switch {
	case n.parent == nil:
		avl.root = replacement

	case n.parent.left == n:
		n.parent.left = replacement

	default:
		n.parent.right = replacement
	}

	if replacement != nil {
		replacement.parent = n.parent
	}
}
