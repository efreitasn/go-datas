package linkedlist

import (
	"reflect"
	"testing"
)

func TestPeekBeginning(t *testing.T) {
	ll := New[int]()

	ll.InsertBeginning(1)
	ll.InsertEnd(2)

	head, hasHead := ll.PeekBeginning()
	expectedHead := 1
	expectedHasHead := true

	if head != expectedHead {
		t.Errorf("got %v, want %v", head, expectedHead)
	}

	if hasHead != expectedHasHead {
		t.Errorf("got %v, want %v", hasHead, expectedHasHead)
	}

	// When the linked list is empty
	ll2 := New[int]()

	head, hasHead = ll2.PeekBeginning()
	expectedHead = 0
	expectedHasHead = false

	if head != expectedHead {
		t.Errorf("got %v, want %v", head, expectedHead)
	}

	if hasHead != expectedHasHead {
		t.Errorf("got %v, want %v", hasHead, expectedHasHead)
	}
}

func TestInsertBeginning(t *testing.T) {
	ll := New[int]()

	ll.InsertBeginning(0)
	ll.InsertBeginning(3)
	ll.InsertBeginning(10)

	head, _ := ll.PeekBeginning()
	expectedHead := 10

	if head != expectedHead {
		t.Errorf("got %v, want %v", head, expectedHead)
	}

	size := ll.Size()
	expectedSize := 3

	if size != expectedSize {
		t.Errorf("got %v, want %v", size, expectedSize)
	}
}

func TestDeleteBeginning(t *testing.T) {
	ll := New[int]()

	ll.InsertBeginning(100)
	ll.InsertBeginning(10)
	ll.InsertBeginning(20)
	ll.DeleteBeginning()

	head, _ := ll.PeekBeginning()
	expectedHead := 10

	if head != expectedHead {
		t.Errorf("got %v, want %v", head, expectedHead)
	}

	// Delete when there's only one element in the list
	ll = New[int]()

	ll.InsertBeginning(1000)
	ll.DeleteBeginning()

	_, hasTail := ll.PeekEnd()
	expectedHasTail := false
	_, hasHead := ll.PeekBeginning()
	expectedHasHead := false

	if hasTail != expectedHasTail {
		t.Errorf("got %v, want %v", hasTail, expectedHasTail)
	}

	if hasHead != expectedHasHead {
		t.Errorf("got %v, want %v", hasHead, expectedHasHead)
	}
}

func TestPeekEnd(t *testing.T) {
	ll := New[int]()

	ll.InsertEnd(10)
	ll.InsertBeginning(599)

	tail, hasTail := ll.PeekEnd()
	expectedTail := 10
	expectedHasTail := true

	if tail != expectedTail {
		t.Errorf("got %v, want %v", tail, expectedTail)
	}

	if hasTail != expectedHasTail {
		t.Errorf("got %v, want %v", hasTail, expectedHasTail)
	}

	// When the linked list is empty
	ll2 := New[int]()

	tail, hasTail = ll2.PeekBeginning()
	expectedTail = 0
	expectedHasTail = false

	if tail != expectedTail {
		t.Errorf("got %v, want %v", tail, expectedTail)
	}

	if hasTail != expectedHasTail {
		t.Errorf("got %v, want %v", hasTail, expectedHasTail)
	}

}

func TestInsertEnd(t *testing.T) {
	ll := New(300, 900)

	ll.InsertEnd(1203)

	tail, _ := ll.PeekEnd()
	expectedTail := 1203

	if tail != expectedTail {
		t.Errorf("got %v, want %v", tail, expectedTail)
	}

	size := ll.Size()
	expectedSize := 3

	if size != expectedSize {
		t.Errorf("got %v, want %v", size, expectedSize)
	}
}

func TestDeleteEnd(t *testing.T) {
	ll := New[int]()

	ll.InsertEnd(399)
	ll.InsertEnd(1020)
	ll.InsertEnd(3000)
	ll.DeleteEnd()

	tail, _ := ll.PeekEnd()
	expectedTail := 1020

	if tail != expectedTail {
		t.Errorf("got %v, want %v", tail, expectedTail)
	}

	// Delete when there's only one element in the list
	ll = New[int]()

	ll.InsertEnd(120)
	ll.DeleteEnd()

	_, hasTail := ll.PeekEnd()
	expectedHasTail := false
	_, hasHead := ll.PeekBeginning()
	expectedHasHead := false

	if hasTail != expectedHasTail {
		t.Errorf("got %v, want %v", hasTail, expectedHasTail)
	}

	if hasHead != expectedHasHead {
		t.Errorf("got %v, want %v", hasHead, expectedHasHead)
	}
}

func TestTraverse(t *testing.T) {
	ll := New[int]()

	ll.InsertBeginning(1)
	ll.InsertEnd(2)
	ll.InsertEnd(3)
	ll.InsertEnd(4)
	ll.InsertEnd(5)
	ll.InsertEnd(6)

	vals := []int{}
	cb := func(v int) bool {
		vals = append(vals, v)

		return true
	}

	// From beginning
	ll.Traverse(true, cb)

	expectedVals := []int{1, 2, 3, 4, 5, 6}

	if !reflect.DeepEqual(vals, expectedVals) {
		t.Errorf("got %v, want %v", vals, expectedVals)
	}

	// From end
	vals = []int{}
	ll.Traverse(false, cb)

	expectedVals = []int{6, 5, 4, 3, 2, 1}

	if !reflect.DeepEqual(vals, expectedVals) {
		t.Errorf("got %v, want %v", vals, expectedVals)
	}

	// Partial
	vals = []int{}
	ll.Traverse(true, func(v int) bool {
		if v >= 5 {
			return false
		}

		vals = append(vals, v)

		return true
	})

	expectedVals = []int{1, 2, 3, 4}

	if !reflect.DeepEqual(vals, expectedVals) {
		t.Errorf("got %v, want %v", vals, expectedVals)
	}
}

func TestString(t *testing.T) {
	ll := New[int]()

	ll.InsertEnd(10)
	ll.InsertEnd(30)
	ll.InsertEnd(90)

	str := ll.String()
	expectedStr := "LinkedList{10, 30, 90}"

	if str != expectedStr {
		t.Errorf("got %v, want %v", str, expectedStr)
	}
}
