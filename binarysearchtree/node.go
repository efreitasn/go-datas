package binarysearchtree

// Node is a node of a binary search tree.
type Node struct {
	value  int
	left   *Node
	right  *Node
	parent *Node
}

// Value returns the value of the node.
func (n *Node) Value() int {
	return n.value
}

// Left returns the left child of the node.
func (n *Node) Left() *Node {
	return n.left
}

// Right returns the right child of the node.
func (n *Node) Right() *Node {
	return n.right
}

// Parent returns the parent of the node.
func (n *Node) Parent() *Node {
	return n.parent
}
