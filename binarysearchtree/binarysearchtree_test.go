package binarysearchtree_test

import (
	"reflect"
	"testing"

	"github.com/efreitasn/go-algs/dfs"
	. "github.com/efreitasn/go-datas/binarysearchtree"
)

func TestInsert(t *testing.T) {
	bts := New()

	bts.Insert(10)
	bts.Insert(30)
	bts.Insert(50)

	size := bts.Size()
	expectedSize := 3

	if size != expectedSize {
		t.Errorf("got %v, want %v", size, expectedSize)
	}

	rootValue := bts.Root().Value()
	expectedRootValue := 10

	if rootValue != expectedRootValue {
		t.Errorf("got %v, want %v", rootValue, expectedRootValue)
	}
}

func TestNodeHeight(t *testing.T) {
	bts := New()

	bts.Insert(100)
	bts.Insert(90)
	bts.Insert(80)
	bts.Insert(95)
	bts.Insert(120)
	bts.Insert(110)
	bts.Insert(130)

	height := bts.NodeHeight(bts.Root().Left())
	expectedHeight := 1

	if height != expectedHeight {
		t.Errorf("got %v, want %v", height, expectedHeight)
	}
}

func TestNodeDepth(t *testing.T) {
	bts := New()

	bts.Insert(100)
	bts.Insert(90)
	bts.Insert(80)
	bts.Insert(95)
	bts.Insert(120)
	bts.Insert(110)
	bts.Insert(130)

	depth := bts.NodeDepth(bts.Root().Left().Left())
	expectedDepth := 2

	if depth != expectedDepth {
		t.Errorf("got %v, want %v", depth, expectedDepth)
	}
}

func TestHeight(t *testing.T) {
	bts := New()

	bts.Insert(100)
	bts.Insert(90)
	bts.Insert(80)
	bts.Insert(95)
	bts.Insert(120)
	bts.Insert(110)
	bts.Insert(130)

	height := bts.Height()
	expectedHeight := 2

	if height != expectedHeight {
		t.Errorf("got %v, want %v", height, expectedHeight)
	}
}

func TestRemove(t *testing.T) {
	bts := New()

	bts.Insert(100)
	bts.Insert(90)
	bts.Insert(80)
	bts.Insert(95)
	bts.Insert(92)
	bts.Insert(96)
	bts.Insert(93)
	bts.Insert(120)
	bts.Insert(110)
	bts.Insert(130)

	bts.Remove(bts.Root().Left().Right())

	vals := make([]int, 0, bts.Size())

	dfs.BinarySearchTreeNLR(bts, func(v int) bool {
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
