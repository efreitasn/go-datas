package binarysearchtree_test

import (
	"reflect"
	"testing"

	"github.com/efreitasn/go-algs/dfs"
	"github.com/efreitasn/go-datas/binarysearchtree"
)

func TestInsert(t *testing.T) {
	bst := binarysearchtree.New[int]()

	bst.Insert(10)
	bst.Insert(30)
	bst.Insert(50)

	size := bst.Size()
	expectedSize := 3

	if size != expectedSize {
		t.Errorf("got %v, want %v", size, expectedSize)
	}

	rootValue := bst.Root().Value()
	expectedRootValue := 10

	if rootValue != expectedRootValue {
		t.Errorf("got %v, want %v", rootValue, expectedRootValue)
	}
}

func TestNodeHeight(t *testing.T) {
	bst := binarysearchtree.New[int]()

	bst.Insert(100)
	bst.Insert(90)
	bst.Insert(80)
	bst.Insert(95)
	bst.Insert(120)
	bst.Insert(110)
	bst.Insert(130)

	height := bst.NodeHeight(bst.Root().Left())
	expectedHeight := 1

	if height != expectedHeight {
		t.Errorf("got %v, want %v", height, expectedHeight)
	}
}

func TestNodeDepth(t *testing.T) {
	bst := binarysearchtree.New[int]()

	bst.Insert(100)
	bst.Insert(90)
	bst.Insert(80)
	bst.Insert(95)
	bst.Insert(120)
	bst.Insert(110)
	bst.Insert(130)

	depth := bst.NodeDepth(bst.Root().Left().Left())
	expectedDepth := 2

	if depth != expectedDepth {
		t.Errorf("got %v, want %v", depth, expectedDepth)
	}
}

func TestHeight(t *testing.T) {
	bst := binarysearchtree.New[int]()

	bst.Insert(100)
	bst.Insert(90)
	bst.Insert(80)
	bst.Insert(95)
	bst.Insert(120)
	bst.Insert(110)
	bst.Insert(130)

	height := bst.Height()
	expectedHeight := 2

	if height != expectedHeight {
		t.Errorf("got %v, want %v", height, expectedHeight)
	}
}

func TestRemove(t *testing.T) {
	bst := binarysearchtree.New[int]()

	bst.Insert(100)
	bst.Insert(90)
	bst.Insert(80)
	bst.Insert(95)
	bst.Insert(92)
	bst.Insert(96)
	bst.Insert(93)
	bst.Insert(120)
	bst.Insert(110)
	bst.Insert(130)

	bst.Remove(bst.Root().Left().Right())

	vals := make([]int, 0, bst.Size())

	dfs.BinarySearchTreeNLR(bst, func(v int) bool {
		vals = append(vals, v)

		return true
	})

	expectedVals := []int{
		100,
		90,
		80,
		96,
		92,
		93,
		120,
		110,
		130,
	}

	if !reflect.DeepEqual(vals, expectedVals) {
		t.Errorf("got %v, want %v", vals, expectedVals)
	}
}

func TestFind(t *testing.T) {
	bst := binarysearchtree.New[int]()

	bst.Insert(100)
	bst.Insert(90)
	bst.Insert(80)
	bst.Insert(95)
	bst.Insert(92)
	bst.Insert(96)
	bst.Insert(93)
	bst.Insert(120)
	bst.Insert(110)
	bst.Insert(130)

	// When a node with the provided value exists
	n, found := bst.Find(92)
	expectedN := bst.Root().Left().Right().Left()
	expectedFound := true

	if !reflect.DeepEqual(n, expectedN) {
		t.Errorf("got %v, want %v", n, expectedN)
	}

	if found != expectedFound {
		t.Errorf("got %v, want %v", found, expectedFound)
	}
}
