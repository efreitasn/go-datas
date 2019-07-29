/*
Package tree provides functions to create/work with trees of ints.

The tree is implemented using an adjacency list, since it uses the graph package under the hood.
*/
package tree

import (
	"github.com/efreitasn/go-datas/graph"
	"github.com/efreitasn/go-datas/linkedlist"
)

// Tree is a tree of ints.
type Tree struct {
	g    *graph.Graph
	root int
}

// New create a tree of ints.
func New(root int) *Tree {
	g := graph.New(true)

	g.AddVertex(root)

	return &Tree{
		g,
		root,
	}
}

// AddNode adds a node to the tree.
func (tr *Tree) AddNode(parent int, v int) (ok bool) {
	if !tr.g.HasVertex(parent) || tr.g.HasVertex(v) {
		return false
	}

	tr.g.AddVertex(v)
	tr.g.AddEdge(parent, v)

	return true
}

// NodeChildren returns the children of a node.
func (tr *Tree) NodeChildren(v int) (children *linkedlist.LinkedList, found bool) {
	adjVertices, found := tr.g.AdjacentVertices(v)

	if !found {
		return nil, false
	}

	return adjVertices, true
}

// NodeHeight returns the number of edges from a node to its deepest descendent (the height of a node) using DFS.
func (tr *Tree) NodeHeight(v int) (height int, found bool) {
	if !tr.g.HasVertex(v) {
		return 0, false
	}

	return tr.nodeHeightRecursive(v), true
}

func (tr *Tree) nodeHeightRecursive(v int) int {
	adjVertices, _ := tr.g.AdjacentVertices(v)

	if adjVertices.Size() == 0 {
		return 0
	}

	var heights []int

	adjVertices.Traverse(true, func(adjV int) {
		heights = append(heights, tr.nodeHeightRecursive(adjV))
	})

	var maxHeight int

	for _, h := range heights {
		if h > maxHeight {
			maxHeight = h
		}
	}

	return maxHeight + 1
}

// NodeDepth returns the number of edges from the root to a node (the depth of a node) using DFS.
func (tr *Tree) NodeDepth(v int) (depth int, found bool) {
	if !tr.g.HasVertex(v) {
		return 0, false
	}

	depth, _ = tr.nodeDepthRecursive(tr.Root(), v)

	return depth, true
}

type nodeDepthItem struct {
	depth  int
	vFound bool
}

func (tr *Tree) nodeDepthRecursive(v, valueToFind int) (depth int, vFound bool) {
	adjVertices, _ := tr.g.AdjacentVertices(v)

	if adjVertices.Size() == 0 || v == valueToFind {
		return 0, v == valueToFind
	}

	var depths []nodeDepthItem

	adjVertices.Traverse(true, func(adjV int) {
		depth, vFound := tr.nodeDepthRecursive(adjV, valueToFind)

		depths = append(
			depths,
			nodeDepthItem{
				depth,
				vFound,
			},
		)
	})

	for _, d := range depths {
		if d.vFound {
			return d.depth + 1, true
		}
	}

	return 0, false
}

// Root returns the root node of the tree.
func (tr *Tree) Root() int {
	return tr.root
}

// Height returns the number of edges between the root of the tree and its deepest descendent, i.e. the heigh of the root.
func (tr *Tree) Height() int {
	height, _ := tr.NodeHeight(tr.Root())

	return height
}

// Size returns the number of nodes in the tree.
func (tr *Tree) Size() int {
	return tr.g.NumVertices()
}
