// Package graph provides functions to create/work with generic graphs.
package graph

import (
	"github.com/efreitasn/go-datas/linkedlist"
)

// Graph is a graph of a comparable type T implemented using an adjacency list.
type Graph[T comparable] struct {
	adjList  map[T]*linkedlist.LinkedList[T]
	directed bool
}

// New creates a graph of a comparable type T.
func New[T comparable](directed bool) *Graph[T] {
	return &Graph[T]{
		adjList:  map[T]*linkedlist.LinkedList[T]{},
		directed: directed,
	}
}

// IsDirected checks whether the graph is directed.
func (g *Graph[T]) IsDirected() bool {
	return g.directed
}

// HasVertex checks whether the graph contains a vertex.
func (g *Graph[T]) HasVertex(v T) bool {
	_, found := g.adjList[v]

	return found
}

// AddVertex adds a vertex to the graph.
func (g *Graph[T]) AddVertex(v T) {
	g.adjList[v] = linkedlist.New[T]()
}

// VertexDegree returns the degree of a vertex in the graph.
func (g *Graph[T]) VertexDegree(v T) (degree int, found bool) {
	adjVertices, found := g.AdjacentVertices(v)

	if !found {
		return 0, false
	}

	return adjVertices.Size(), true
}

// AdjacentVertices returns the adjacent vertices of a vertex.
func (g *Graph[T]) AdjacentVertices(v T) (list *linkedlist.LinkedList[T], found bool) {
	if !g.HasVertex(v) {
		return nil, false
	}

	return g.adjList[v], true
}

// NumVertices returns the number of vertices in the graph.
func (g *Graph[T]) NumVertices() int {
	return len(g.adjList)
}

// HasEdge checks whether the graph contains an edge
func (g *Graph[T]) HasEdge(v1, v2 T) bool {
	if !g.HasVertex(v1) || !g.HasVertex(v2) {
		return false
	}

	var contains bool

	g.adjList[v1].Traverse(true, func(v T) bool {
		if v == v2 {
			contains = true

			return false
		}

		return true
	})

	return contains
}

// AddEdge adds an edge to the graph.
func (g *Graph[T]) AddEdge(v T, v2 T) bool {
	if !g.HasVertex(v) || !g.HasVertex(v2) {
		return false
	}

	g.adjList[v].InsertEnd(v2)

	if !g.IsDirected() {
		g.adjList[v2].InsertEnd(v)
	}

	return true
}

// NumEdges returns the number of edges in the graph.
func (g *Graph[T]) NumEdges() int {
	var n int

	for _, edges := range g.adjList {
		n += edges.Size()
	}

	if !g.IsDirected() {
		return n / 2
	}

	return n
}
