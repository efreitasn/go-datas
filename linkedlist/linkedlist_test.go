package linkedlist

import "testing"

func TestPeekBeginning(t *testing.T) {
	ll := New()

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
	ll2 := New()

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
	ll := New()

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
	ll := New()

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
	ll = New()

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
	ll := New()

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
	ll2 := New()

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
	ll := New()

	ll.InsertEnd(300)
	ll.InsertEnd(900)
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
	ll := New()

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
	ll = New()

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

func TestContains(t *testing.T) {
	ll := New()

	ll.InsertEnd(2939)
	ll.InsertBeginning(3948)
	ll.InsertEnd(9192)
	ll.InsertBeginning(12)

	result := ll.Contains(9192)
	expectedResult := true

	if result != expectedResult {
		t.Errorf("got %v, want %v", result, expectedResult)
	}

	result = ll.Contains(100000)
	expectedResult = false

	if result != expectedResult {
		t.Errorf("got %v, want %v", result, expectedResult)
	}

}
