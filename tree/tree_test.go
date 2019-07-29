package tree

import (
	"reflect"
	"testing"

	"github.com/efreitasn/go-datas/linkedlist"
)

func TestAddNode(t *testing.T) {
	tr := New(1)

	tr.AddNode(1, 2)
	tr.AddNode(2, 4)
	tr.AddNode(2, 5)

	parentFound := tr.AddNode(1, 3)
	expectedParentFound := true

	if parentFound != expectedParentFound {
		t.Errorf("got %v, want %v", parentFound, expectedParentFound)
	}

	root := tr.Root()
	expectedRoot := 1

	if root != expectedRoot {
		t.Errorf("got %v, want %v", root, expectedRoot)
	}

	size := tr.Size()
	expectedSize := 5

	if size != expectedSize {
		t.Errorf("got %v, want %v", size, expectedSize)
	}

	// When the parent node doesn't exist
	parentFound = tr.AddNode(1000, 50)
	expectedParentFound = false

	if parentFound != expectedParentFound {
		t.Errorf("got %v, want %v", parentFound, expectedParentFound)
	}
}

func TestNodeChildren(t *testing.T) {
	tr := New(1)

	tr.AddNode(1, 2)
	tr.AddNode(1, 3)
	tr.AddNode(1, 4)

	children, found := tr.NodeChildren(1)
	expectedChildren := linkedlist.New(2, 3, 4)
	expectedFound := true

	if !reflect.DeepEqual(children, expectedChildren) {
		t.Errorf("got %v, want %v", children, expectedChildren)
	}

	if found != expectedFound {
		t.Errorf("got %v, want %v", found, expectedFound)
	}

	// When the node doesn't exist
	children, found = tr.NodeChildren(1000)
	expectedChildren = nil
	expectedFound = false

	if children != expectedChildren {
		t.Errorf("got %v, want %v", children, expectedChildren)
	}

	if found != expectedFound {
		t.Errorf("got %v, want %v", found, expectedFound)
	}
}

func TestNodeHeight(t *testing.T) {
	trRoot := 1
	tr := New(trRoot)

	tr.AddNode(1, 2)
	tr.AddNode(1, 3)

	tr.AddNode(2, 4)
	tr.AddNode(2, 5)

	tr.AddNode(3, 6)

	tr.AddNode(6, 7)

	height, found := tr.NodeHeight(2)
	expectedHeight := 1
	expectedFound := true

	if height != expectedHeight {
		t.Errorf("got %v, want %v", height, expectedHeight)
	}

	if found != expectedFound {
		t.Errorf("got %v, want %v", found, expectedFound)
	}

	// When the node doesn't exist
	height, found = tr.NodeHeight(1000)
	expectedHeight = 0
	expectedFound = false

	if height != expectedHeight {
		t.Errorf("got %v, want %v", height, expectedHeight)
	}

	if found != expectedFound {
		t.Errorf("got %v, want %v", found, expectedFound)
	}

	// When the node is the root
	height, found = tr.NodeHeight(trRoot)
	expectedHeight = 3
	expectedFound = true

	if height != expectedHeight {
		t.Errorf("got %v, want %v", height, expectedHeight)
	}

	if found != expectedFound {
		t.Errorf("got %v, want %v", found, expectedFound)
	}

	// When the node is a leaf
	height, found = tr.NodeHeight(7)
	expectedHeight = 0
	expectedFound = true

	if height != expectedHeight {
		t.Errorf("got %v, want %v", height, expectedHeight)
	}

	if found != expectedFound {
		t.Errorf("got %v, want %v", found, expectedFound)
	}

}

func TestNodeDepth(t *testing.T) {
	trRoot := 1
	tr := New(trRoot)

	tr.AddNode(1, 2)
	tr.AddNode(1, 3)

	tr.AddNode(2, 4)
	tr.AddNode(2, 5)

	tr.AddNode(3, 6)

	tr.AddNode(6, 7)

	depth, found := tr.NodeDepth(6)
	expectedDepth := 2
	expectedFound := true

	if depth != expectedDepth {
		t.Errorf("got %v, want %v", depth, expectedDepth)
	}

	if found != expectedFound {
		t.Errorf("got %v, want %v", found, expectedFound)
	}

	// When the node doesn't exist
	depth, found = tr.NodeDepth(1000)
	expectedDepth = 0
	expectedFound = false

	if depth != expectedDepth {
		t.Errorf("got %v, want %v", depth, expectedDepth)
	}

	if found != expectedFound {
		t.Errorf("got %v, want %v", found, expectedFound)
	}

	// When the node is the root
	depth, found = tr.NodeDepth(trRoot)
	expectedDepth = 0
	expectedFound = true

	if depth != expectedDepth {
		t.Errorf("got %v, want %v", depth, expectedDepth)
	}

	if found != expectedFound {
		t.Errorf("got %v, want %v", found, expectedFound)
	}
}

func TestHeight(t *testing.T) {
	tr := New(1)

	tr.AddNode(1, 2)
	tr.AddNode(1, 5)

	tr.AddNode(2, 9)

	height := tr.Height()
	expectedHeight := 2

	if height != expectedHeight {
		t.Errorf("got %v, want %v", height, expectedHeight)
	}
}

func TestHasNode(t *testing.T) {
	tr := New(1)

	tr.AddNode(1, 10)
	tr.AddNode(10, 20)

	hasNode := tr.HasNode(10)
	expectedHasNode := true

	if hasNode != expectedHasNode {
		t.Errorf("got %v, want %v", hasNode, expectedHasNode)
	}

	// When the node doesn't exist
	hasNode = tr.HasNode(100)
	expectedHasNode = false

	if hasNode != expectedHasNode {
		t.Errorf("got %v, want %v", hasNode, expectedHasNode)
	}

}
