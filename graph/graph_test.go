package graph

import (
	"reflect"
	"testing"

	"github.com/efreitasn/go-datas/linkedlist"
)

func TestAddVertex(t *testing.T) {
	g := New(true)
	v := 10

	g.AddVertex(v)

	numVertices := g.NumVertices()
	expectedNumVertices := 1

	if numVertices != expectedNumVertices {
		t.Errorf("got %v, want %v", numVertices, expectedNumVertices)
	}

	hasVertex := g.HasVertex(v)
	expectedHasVertex := true

	if hasVertex != expectedHasVertex {
		t.Errorf("got %v, want %v", hasVertex, expectedHasVertex)
	}
}

func TestHasEdge(t *testing.T) {
	g := New(true)
	v1 := 10
	v2 := 20

	g.AddVertex(v1)
	g.AddVertex(v2)

	g.AddEdge(v1, v2)

	hasEdge := g.HasEdge(v1, v2)
	expectedHasEdge := true

	if hasEdge != expectedHasEdge {
		t.Errorf("got %v, want %v", hasEdge, expectedHasEdge)
	}

	hasEdge = g.HasEdge(v2, v1)
	expectedHasEdge = false

	if hasEdge != expectedHasEdge {
		t.Errorf("got %v, want %v", hasEdge, expectedHasEdge)
	}
}

func TestAddEdge(t *testing.T) {
	// Directed graph
	gDirected := New(true)
	v1 := 10
	v2 := 20

	gDirected.AddVertex(v1)
	gDirected.AddVertex(v2)

	gDirected.AddEdge(v1, v2)

	v1Degree, _ := gDirected.VertexDegree(v1)
	expectedV1Degree := 1

	if v1Degree != expectedV1Degree {
		t.Errorf("got %v, want %v", v1Degree, expectedV1Degree)
	}

	v2Degree, _ := gDirected.VertexDegree(v2)
	expectedV2Degree := 0

	if v2Degree != expectedV2Degree {
		t.Errorf("got %v, want %v", v2Degree, expectedV2Degree)
	}

	// Undirected graph
	gUndirected := New(false)

	gUndirected.AddVertex(v1)
	gUndirected.AddVertex(v2)

	gUndirected.AddEdge(v1, v2)

	v1Degree, _ = gUndirected.VertexDegree(v1)
	expectedV1Degree = 1

	if v1Degree != expectedV1Degree {
		t.Errorf("got %v, want %v", v1Degree, expectedV1Degree)
	}

	v2Degree, _ = gUndirected.VertexDegree(v2)
	expectedV2Degree = 1

	if v2Degree != expectedV2Degree {
		t.Errorf("got %v, want %v", v2Degree, expectedV2Degree)
	}
}

func TestNumEdges(t *testing.T) {
	// Directed graph
	gDirected := New(true)
	v1 := 10
	v2 := 20
	v3 := 30

	gDirected.AddVertex(v1)
	gDirected.AddVertex(v2)
	gDirected.AddVertex(v3)

	gDirected.AddEdge(v1, v2)
	gDirected.AddEdge(v2, v3)
	gDirected.AddEdge(v3, v2)

	numEdges := gDirected.NumEdges()
	expectedNumEdges := 3

	if numEdges != expectedNumEdges {
		t.Errorf("got %v, want %v", numEdges, expectedNumEdges)
	}

	// Undirected graphs
	gUndirected := New(false)

	gUndirected.AddVertex(v1)
	gUndirected.AddVertex(v2)
	gUndirected.AddVertex(v3)

	gUndirected.AddEdge(v1, v2)
	gUndirected.AddEdge(v2, v3)

	numEdges = gUndirected.NumEdges()
	expectedNumEdges = 2

	if numEdges != expectedNumEdges {
		t.Errorf("got %v, want %v", numEdges, expectedNumEdges)
	}
}

func TestAdjacentVertices(t *testing.T) {
	g := New(true)
	v1 := 30
	v2 := 100
	v3 := 300

	g.AddVertex(v1)
	g.AddVertex(v2)
	g.AddVertex(v3)

	g.AddEdge(v1, v2)
	g.AddEdge(v1, v3)

	adjVertices, _ := g.AdjacentVertices(v1)
	expectedAdjVertices := linkedlist.New(v2, v3)

	if !reflect.DeepEqual(adjVertices, expectedAdjVertices) {
		t.Errorf("got %v, want %v", adjVertices, expectedAdjVertices)
	}
}
