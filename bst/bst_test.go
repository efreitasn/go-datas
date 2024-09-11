package bst

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

var nodeComparer = cmp.Comparer(func(a, b *Node[int]) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	parentsEqual := (a.parent == nil && b.parent == nil) ||
		(a.parent != nil && b.parent != nil && a.parent.value == b.parent.value)

	return (parentsEqual && a.value == b.value && a.balanceFactor == b.balanceFactor)
})

func TestBST_int_Insert(t *testing.T) {
	type testCase struct {
		name              string
		values            []int
		expectedNLRValues []*Node[int]
	}

	testCases := []testCase{
		{
			name:   "valid",
			values: []int{5, 3, 6, 4, 7, 2, 8},
			expectedNLRValues: []*Node[int]{
				{
					value: 5,
				},
				{
					parent: &Node[int]{value: 5},
					value:  3,
				},
				{
					parent: &Node[int]{value: 3},
					value:  2,
				},
				{
					parent: &Node[int]{value: 3},
					value:  4,
				},
				{
					parent: &Node[int]{value: 5},
					value:  6,
				},
				{
					parent: &Node[int]{value: 6},
					value:  7,
				},
				{
					parent: &Node[int]{value: 7},
					value:  8,
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			bst := NewBST[int]()

			for _, v := range tc.values {
				bst.Insert(v)
			}

			vals := make([]*Node[int], 0, bst.Size())
			bst.Traverse(NLR, func(n *Node[int]) bool {
				vals = append(vals, n)

				return true
			})

			if diff := cmp.Diff(len(tc.values), bst.Size()); diff != "" {
				t.Errorf("invalid size:\n%v", diff)
			}

			if diff := cmp.Diff(tc.expectedNLRValues, vals, nodeComparer); diff != "" {
				t.Errorf("invalid values:\n%v", diff)
			}
		})
	}
}

func TestBST_int_Find(t *testing.T) {
	type testCase struct {
		name           string
		valuesToInsert []int
		valueToFind    int
		expectedNode   *Node[int]
		expectedFound  bool
	}

	testCases := []testCase{
		{
			name:           "node_found",
			valuesToInsert: []int{3, 2, 1, 9, 5, 10},
			valueToFind:    9,
			expectedNode: &Node[int]{
				parent: &Node[int]{value: 3},
				value:  9,
			},
			expectedFound: true,
		},
		{
			name:           "node_not_found",
			valuesToInsert: []int{1, 2, 3, 10, 9, 5},
			valueToFind:    15,
			expectedNode:   nil,
			expectedFound:  false,
		},
		{
			name:          "empty_tree",
			valueToFind:   15,
			expectedNode:  nil,
			expectedFound: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			bst := NewBST[int]()

			for _, v := range tc.valuesToInsert {
				bst.Insert(v)
			}

			n, found := bst.Find(tc.valueToFind)

			if diff := cmp.Diff(tc.expectedNode, n, nodeComparer); diff != "" {
				t.Errorf("invalid node:\n%v", diff)
			}

			if diff := cmp.Diff(tc.expectedFound, found); diff != "" {
				t.Errorf("invalid found:\n%v", diff)
			}
		})
	}
}

func TestBST_int_Remove(t *testing.T) {
	type testCase struct {
		name              string
		valueToRemove     int
		valuesToInsert    []int
		expectedNLRValues []*Node[int]
	}

	testCases := []testCase{
		{
			name:           "leaf_node",
			valueToRemove:  415,
			valuesToInsert: []int{400, 200, 300, 210, 410, 405, 500, 415, 550},
			expectedNLRValues: []*Node[int]{
				{
					value: 400,
				},
				{
					parent: &Node[int]{value: 400},
					value:  200,
				},
				{
					parent: &Node[int]{value: 200},
					value:  300,
				},
				{
					parent: &Node[int]{value: 300},
					value:  210,
				},
				{
					parent: &Node[int]{value: 400},
					value:  410,
				},
				{
					parent: &Node[int]{value: 410},
					value:  405,
				},
				{
					parent: &Node[int]{value: 410},
					value:  500,
				},
				{
					parent: &Node[int]{value: 500},
					value:  550,
				},
			},
		},
		{
			name:           "with_right_child_only",
			valueToRemove:  415,
			valuesToInsert: []int{400, 200, 300, 210, 410, 405, 500, 415, 450, 550},
			expectedNLRValues: []*Node[int]{
				{
					value: 400,
				},
				{
					parent: &Node[int]{value: 400},
					value:  200,
				},
				{
					parent: &Node[int]{value: 200},
					value:  300,
				},
				{
					parent: &Node[int]{value: 300},
					value:  210,
				},
				{
					parent: &Node[int]{value: 400},
					value:  410,
				},
				{
					parent: &Node[int]{value: 410},
					value:  405,
				},
				{
					parent: &Node[int]{value: 410},
					value:  500,
				},
				{
					parent: &Node[int]{value: 500},
					value:  450,
				},
				{
					parent: &Node[int]{value: 500},
					value:  550,
				},
			},
		},
		{
			name:           "with_left_child_only",
			valueToRemove:  415,
			valuesToInsert: []int{400, 200, 300, 210, 410, 405, 500, 415, 412, 550},
			expectedNLRValues: []*Node[int]{
				{
					value: 400,
				},
				{
					parent: &Node[int]{value: 400},
					value:  200,
				},
				{
					parent: &Node[int]{value: 200},
					value:  300,
				},
				{
					parent: &Node[int]{value: 300},
					value:  210,
				},
				{
					parent: &Node[int]{value: 400},
					value:  410,
				},
				{
					parent: &Node[int]{value: 410},
					value:  405,
				},
				{
					parent: &Node[int]{value: 410},
					value:  500,
				},
				{
					parent: &Node[int]{value: 500},
					value:  412,
				},
				{
					parent: &Node[int]{value: 500},
					value:  550,
				},
			},
		},
		{
			name:           "with_both_children",
			valueToRemove:  415,
			valuesToInsert: []int{400, 200, 300, 210, 410, 405, 500, 415, 412, 411, 450, 445, 455, 550},
			expectedNLRValues: []*Node[int]{
				{
					value: 400,
				},
				{
					parent: &Node[int]{value: 400},
					value:  200,
				},
				{
					parent: &Node[int]{value: 200},
					value:  300,
				},
				{
					parent: &Node[int]{value: 300},
					value:  210,
				},
				{
					parent: &Node[int]{value: 400},
					value:  410,
				},
				{
					parent: &Node[int]{value: 410},
					value:  405,
				},
				{
					parent: &Node[int]{value: 410},
					value:  500,
				},
				{
					parent: &Node[int]{value: 500},
					value:  412,
				},
				{
					parent: &Node[int]{value: 412},
					value:  411,
				},
				{
					parent: &Node[int]{value: 412},
					value:  450,
				},
				{
					parent: &Node[int]{value: 450},
					value:  445,
				},
				{
					parent: &Node[int]{value: 450},
					value:  455,
				},
				{
					parent: &Node[int]{value: 500},
					value:  550,
				},
			},
		},
		{
			name:           "with_both_children_and_left_right_child",
			valueToRemove:  415,
			valuesToInsert: []int{400, 200, 300, 210, 410, 405, 500, 415, 412, 411, 414, 413, 450, 445, 455, 550},
			expectedNLRValues: []*Node[int]{
				{
					value: 400,
				},
				{
					parent: &Node[int]{value: 400},
					value:  200,
				},
				{
					parent: &Node[int]{value: 200},
					value:  300,
				},
				{
					parent: &Node[int]{value: 300},
					value:  210,
				},
				{
					parent: &Node[int]{value: 400},
					value:  410,
				},
				{
					parent: &Node[int]{value: 410},
					value:  405,
				},
				{
					parent: &Node[int]{value: 410},
					value:  500,
				},
				{
					parent: &Node[int]{value: 500},
					value:  414,
				},
				{
					parent: &Node[int]{value: 414},
					value:  412,
				},
				{
					parent: &Node[int]{value: 412},
					value:  411,
				},
				{
					parent: &Node[int]{value: 412},
					value:  413,
				},
				{
					parent: &Node[int]{value: 414},
					value:  450,
				},
				{
					parent: &Node[int]{value: 450},
					value:  445,
				},
				{
					parent: &Node[int]{value: 450},
					value:  455,
				},
				{
					parent: &Node[int]{value: 500},
					value:  550,
				},
			},
		},
		{
			name:           "with_both_children_and_left_right_child_2",
			valueToRemove:  440,
			valuesToInsert: []int{400, 200, 300, 210, 410, 405, 500, 440, 420, 411, 422, 421, 424, 423, 426, 425, 450, 445, 455, 550},
			expectedNLRValues: []*Node[int]{
				{
					value: 400,
				},
				{
					parent: &Node[int]{value: 400},
					value:  200,
				},
				{
					parent: &Node[int]{value: 200},
					value:  300,
				},
				{
					parent: &Node[int]{value: 300},
					value:  210,
				},
				{
					parent: &Node[int]{value: 400},
					value:  410,
				},
				{
					parent: &Node[int]{value: 410},
					value:  405,
				},
				{
					parent: &Node[int]{value: 410},
					value:  500,
				},
				{
					parent: &Node[int]{value: 500},
					value:  426,
				},
				{
					parent: &Node[int]{value: 426},
					value:  420,
				},
				{
					parent: &Node[int]{value: 420},
					value:  411,
				},
				{
					parent: &Node[int]{value: 420},
					value:  422,
				},
				{
					parent: &Node[int]{value: 422},
					value:  421,
				},
				{
					parent: &Node[int]{value: 422},
					value:  424,
				},
				{
					parent: &Node[int]{value: 424},
					value:  423,
				},
				{
					parent: &Node[int]{value: 424},
					value:  425,
				},
				{
					parent: &Node[int]{value: 426},
					value:  450,
				},
				{
					parent: &Node[int]{value: 450},
					value:  445,
				},
				{
					parent: &Node[int]{value: 450},
					value:  455,
				},
				{
					parent: &Node[int]{value: 500},
					value:  550,
				},
			},
		},
		{
			name:              "root_only_node",
			valueToRemove:     400,
			valuesToInsert:    []int{400},
			expectedNLRValues: []*Node[int]{},
		},
		{
			name:           "root_many_nodes",
			valueToRemove:  400,
			valuesToInsert: []int{400, 300, 500, 350, 340, 360},
			expectedNLRValues: []*Node[int]{
				{
					value: 360,
				},
				{
					parent: &Node[int]{value: 360},
					value:  300,
				},
				{
					parent: &Node[int]{value: 300},
					value:  350,
				},
				{
					parent: &Node[int]{value: 350},
					value:  340,
				},
				{
					parent: &Node[int]{value: 360},
					value:  500,
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			bst := NewBST[int]()

			for _, v := range tc.valuesToInsert {
				bst.Insert(v)
			}

			n, ok := bst.Find(tc.valueToRemove)
			if !ok {
				t.Fatalf("node not found")
			}

			oldSize := bst.Size()
			bst.Remove(n)

			vals := make([]*Node[int], 0, bst.Size())
			bst.Traverse(NLR, func(n *Node[int]) bool {
				vals = append(vals, n)

				return true
			})

			if diff := cmp.Diff(oldSize-1, bst.Size()); diff != "" {
				t.Errorf("invalid size:\n%v", diff)
			}

			if diff := cmp.Diff(tc.expectedNLRValues, vals, nodeComparer); diff != "" {
				t.Errorf("invalid values:\n%v", diff)
			}
		})
	}
}
