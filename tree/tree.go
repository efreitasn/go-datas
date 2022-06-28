/*
Package tree provides functions to create/work with generic trees

The tree is implemented using an adjacency list, since it uses the graph package under the hood.
*/
package tree

import (
	"github.com/efreitasn/go-datas/graph"
	"github.com/efreitasn/go-datas/linkedlist"
)

// Tree is a tree of a comparable type T.
type Tree[T comparable] struct {
	g    *graph.Graph[T]
	root T
}

// New creates a tree of a comparable type T.
func New[T comparable](root T) *Tree[T] {
	g := graph.New[T](true)

	g.AddVertex(root)

	return &Tree[T]{
		g,
		root,
	}
}

// HasNode checks whether a node exists in the tree
func (tr *Tree[T]) HasNode(v T) bool {
	return tr.g.HasVertex(v)
}

// AddNode adds a node to the tree.
func (tr *Tree[T]) AddNode(parent, v T) (ok bool) {
	if !tr.HasNode(parent) || tr.g.HasVertex(v) {
		return false
	}

	tr.g.AddVertex(v)
	tr.g.AddEdge(parent, v)

	return true
}

// NodeChildren returns the children of a node.
func (tr *Tree[T]) NodeChildren(v T) (children *linkedlist.LinkedList[T], found bool) {
	adjVertices, found := tr.g.AdjacentVertices(v)

	if !found {
		return nil, false
	}

	return adjVertices, true
}

// NodeHeight returns the number of edges from a node to its deepest descendent (the height of a node) using DFS.
func (tr *Tree[T]) NodeHeight(v T) (height int, found bool) {
	if !tr.HasNode(v) {
		return 0, false
	}

	return tr.nodeHeightRecursive(v), true
}

func (tr *Tree[T]) nodeHeightRecursive(v T) int {
	adjVertices, _ := tr.g.AdjacentVertices(v)

	if adjVertices.Size() == 0 {
		return 0
	}

	var heights []int

	adjVertices.Traverse(true, func(adjV T) bool {
		heights = append(heights, tr.nodeHeightRecursive(adjV))

		return true
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
func (tr *Tree[T]) NodeDepth(v T) (depth int, found bool) {
	if !tr.HasNode(v) {
		return 0, false
	}

	depth, _ = tr.nodeDepthRecursive(tr.Root(), v)

	return depth, true
}

type nodeDepthItem struct {
	depth  int
	vFound bool
}

func (tr *Tree[T]) nodeDepthRecursive(v, valueToFind T) (depth int, vFound bool) {
	adjVertices, _ := tr.g.AdjacentVertices(v)

	if adjVertices.Size() == 0 || v == valueToFind {
		return 0, v == valueToFind
	}

	var depths []nodeDepthItem

	adjVertices.Traverse(true, func(adjV T) bool {
		depth, vFound := tr.nodeDepthRecursive(adjV, valueToFind)

		depths = append(
			depths,
			nodeDepthItem{
				depth,
				vFound,
			},
		)

		return true
	})

	for _, d := range depths {
		if d.vFound {
			return d.depth + 1, true
		}
	}

	return 0, false
}

// Root returns the root node of the tree.
func (tr *Tree[T]) Root() T {
	return tr.root
}

// Height returns the number of edges between the root of the tree and its deepest descendent, i.e. the heigh of the root.
func (tr *Tree[T]) Height() int {
	height, _ := tr.NodeHeight(tr.Root())

	return height
}

// Size returns the number of nodes in the tree.
func (tr *Tree[T]) Size() int {
	return tr.g.NumVertices()
}
