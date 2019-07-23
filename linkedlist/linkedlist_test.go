package linkedlist

import "testing"

func TestPeekBeginning(t *testing.T) {
	ll := New()

	ll.InsertBeginning("bar")
	ll.InsertEnd("foo")

	head, hasHead := ll.PeekBeginning()
	expectedHead := "bar"
	expectedHasHead := true

	if head != expectedHead {
		t.Errorf("got %v, want %v", head, expectedHead)
	}

	if hasHead != expectedHasHead {
		t.Errorf("got %v, want %v", hasHead, expectedHasHead)
	}
}

func TestInsertBeginning(t *testing.T) {
	ll := New()

	ll.InsertBeginning("foo")
	ll.InsertBeginning("bar")
	ll.InsertBeginning("foobar")

	head, _ := ll.PeekBeginning()
	expectedHead := "foobar"

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

	ll.InsertBeginning("foo")
	ll.InsertBeginning("bar")
	ll.InsertBeginning("foobar")
	ll.DeleteBeginning()

	head, _ := ll.PeekBeginning()
	expectedHead := "bar"

	if head != expectedHead {
		t.Errorf("got %v, want %v", head, expectedHead)
	}

	// Delete when there's only one element in the list
	ll = New()

	ll.InsertBeginning("bar")
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

	ll.InsertEnd("thing1")
	ll.InsertBeginning("thing2")

	tail, hasTail := ll.PeekEnd()
	expectedTail := "thing1"
	expectedHasTail := true

	if tail != expectedTail {
		t.Errorf("got %v, want %v", tail, expectedTail)
	}

	if hasTail != expectedHasTail {
		t.Errorf("got %v, want %v", hasTail, expectedHasTail)
	}
}

func TestInsertEnd(t *testing.T) {
	ll := New()

	ll.InsertEnd("thing1")
	ll.InsertEnd("thing2")
	ll.InsertEnd("thing3")

	tail, _ := ll.PeekEnd()
	expectedTail := "thing3"

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

	ll.InsertEnd("foo")
	ll.InsertEnd("bar")
	ll.InsertEnd("foobar")
	ll.DeleteEnd()

	tail, _ := ll.PeekEnd()
	expectedTail := "bar"

	if tail != expectedTail {
		t.Errorf("got %v, want %v", tail, expectedTail)
	}

	// Delete when there's only one element in the list
	ll = New()

	ll.InsertEnd("bar")
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

	ll.InsertEnd("foo")
	ll.InsertBeginning("bar")
	ll.InsertEnd("thing1")
	ll.InsertBeginning("thing2")

	result := ll.Contains("bar")
	expectedResult := true

	if result != expectedResult {
		t.Errorf("got %v, want %v", result, expectedResult)
	}

	result = ll.Contains("nothing")
	expectedResult = false

	if result != expectedResult {
		t.Errorf("got %v, want %v", result, expectedResult)
	}

}
