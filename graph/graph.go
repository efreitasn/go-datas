// Package graph provides functions to create/work with graphs of integers
package graph

import (
	"github.com/efreitasn/go-datas/linkedlist"
)

// Graph is a graph of integers implemented using an adjacency list
type Graph struct {
	adjList  map[int]*linkedlist.LinkedList
	directed bool
}

// New creates a graph of integers
func New(directed bool) *Graph {
	return &Graph{
		adjList:  map[int]*linkedlist.LinkedList{},
		directed: directed,
	}
}

// IsDirected checks whether the graph is directed
func (g *Graph) IsDirected() bool {
	return g.directed
}

// AddVertex adds a vertex to the graph
func (g *Graph) AddVertex(v int) {
	g.adjList[v] = linkedlist.New()
}

// HasVertex checks whether the graph contains a vertex
func (g *Graph) HasVertex(v int) bool {
	_, found := g.adjList[v]

	return found
}

// AddEdge adds an edge to the graph
func (g *Graph) AddEdge(v int, v2 int) bool {
	if !g.HasVertex(v) || !g.HasVertex(v2) {
		return false
	}

	g.adjList[v].InsertEnd(v2)

	if !g.IsDirected() {
		g.adjList[v2].InsertEnd(v)
	}

	return true
}

// VertexDegree returns the degree of a vertex in the graph
func (g *Graph) VertexDegree(v int) (degree int, found bool) {
	if !g.HasVertex(v) {
		return 0, false
	}

	return g.adjList[v].Size(), true
}

// NumVertices returns the number of vertices in the graph
func (g *Graph) NumVertices() int {
	return len(g.adjList)
}

// NumEdges returns the number of edges in the graph
func (g *Graph) NumEdges() int {
	var n int

	for _, edges := range g.adjList {
		n += edges.Size()
	}

	if !g.IsDirected() {
		return n / 2
	}

	return n
}
