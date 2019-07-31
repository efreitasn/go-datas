// Package binarysearchtree provides functions to create/work with binary search trees of ints.
package binarysearchtree

// BinarySearchTree is a binary search tree of ints.
type BinarySearchTree struct {
	root *Node
	size int
}

// New creates a binary search tree of ints.
func New() *BinarySearchTree {
	return &BinarySearchTree{}
}

// Insert adds a value to the binary search tree.
func (bts *BinarySearchTree) Insert(v int) {
	newRoot, inserted := insertRecursive(nil, bts.root, v)

	bts.root = newRoot

	if inserted {
		bts.size++
	}
}

func insertRecursive(previous *Node, current *Node, v int) (node *Node, ok bool) {
	if current == nil {
		return &Node{
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
func (bts *BinarySearchTree) Remove(n *Node) {
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

		bts.Remove(successor)

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
func (bts *BinarySearchTree) NodeHeight(n *Node) int {
	return nodeHeightRecursive(n)
}

func nodeHeightRecursive(n *Node) int {
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
func (bts *BinarySearchTree) NodeDepth(n *Node) int {
	return nodeDepthRecursive(n)
}

func nodeDepthRecursive(n *Node) int {
	if n.Parent() == nil {
		return 0
	}

	return nodeDepthRecursive(n.Parent()) + 1
}

// Find returns the node of the binary search value that has a specific value.
func (bts *BinarySearchTree) Find(v int) (n *Node, found bool) {
	n = bts.Root()

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
func (bts *BinarySearchTree) Root() *Node {
	return bts.root
}

// Height returns the number of edges between the root of the tree and its deepest descendent, i.e. the height of the root.
func (bts *BinarySearchTree) Height() int {
	return bts.NodeHeight(bts.Root())
}

// Size returns the number of nodes in the binary search tree.
func (bts *BinarySearchTree) Size() int {
	return bts.size
}
