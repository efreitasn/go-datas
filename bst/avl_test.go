package bst

import (
	"slices"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestAVL_int_Insert(t *testing.T) {
	type testCase struct {
		name              string
		values            []int
		expectedNLRValues []*Node[int]
	}

	testCases := []testCase{
		{
			name:   "no_rotation",
			values: []int{100, 50, 110, 40, 105, 111},
			expectedNLRValues: []*Node[int]{
				{
					value:         100,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 100},
					value:         50,
					balanceFactor: -1,
				},
				{
					parent:        &Node[int]{value: 50},
					value:         40,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 100},
					value:         110,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 110},
					value:         105,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 110},
					value:         111,
					balanceFactor: 0,
				},
			},
		},
		{
			name:   "rotate_left",
			values: []int{2, 1, 3, 4, 5},
			expectedNLRValues: []*Node[int]{
				{
					value:         2,
					balanceFactor: 1,
				},
				{
					parent:        &Node[int]{value: 2},
					value:         1,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 2},
					value:         4,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 4},
					value:         3,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 4},
					value:         5,
					balanceFactor: 0,
				},
			},
		},
		{
			name:   "rotate_right_left_bf_0",
			values: []int{2, 1, 3, 5, 4},
			expectedNLRValues: []*Node[int]{
				{
					value:         2,
					balanceFactor: 1,
				},
				{
					parent:        &Node[int]{value: 2},
					value:         1,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 2},
					value:         4,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 4},
					value:         3,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 4},
					value:         5,
					balanceFactor: 0,
				},
			},
		},
		{
			name:   "rotate_right_left_bf_-1",
			values: []int{2, 1, 9, 7, 10, 6},
			expectedNLRValues: []*Node[int]{
				{
					value:         7,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 7},
					value:         2,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 2},
					value:         1,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 2},
					value:         6,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 7},
					value:         9,
					balanceFactor: 1,
				},
				{
					parent:        &Node[int]{value: 9},
					value:         10,
					balanceFactor: 0,
				},
			},
		},
		{
			name:   "rotate_right_left_bf_1",
			values: []int{2, 1, 9, 7, 10, 8},
			expectedNLRValues: []*Node[int]{
				{
					value:         7,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 7},
					value:         2,
					balanceFactor: -1,
				},
				{
					parent:        &Node[int]{value: 2},
					value:         1,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 7},
					value:         9,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 9},
					value:         8,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 9},
					value:         10,
					balanceFactor: 0,
				},
			},
		},
		{
			name:   "rotate_right",
			values: []int{4, 5, 3, 2, 1},
			expectedNLRValues: []*Node[int]{
				{
					value:         4,
					balanceFactor: -1,
				},
				{
					parent:        &Node[int]{value: 4},
					value:         2,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 2},
					value:         1,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 2},
					value:         3,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 4},
					value:         5,
					balanceFactor: 0,
				},
			},
		},
		{
			name:   "rotate_left_right_bf_0",
			values: []int{4, 5, 3, 1, 2},
			expectedNLRValues: []*Node[int]{
				{
					value:         4,
					balanceFactor: -1,
				},
				{
					parent:        &Node[int]{value: 4},
					value:         2,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 2},
					value:         1,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 2},
					value:         3,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 4},
					value:         5,
					balanceFactor: 0,
				},
			},
		},
		{
			name:   "rotate_left_right_bf_1",
			values: []int{500, 600, 400, 300, 450, 460},
			expectedNLRValues: []*Node[int]{
				{
					value:         450,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 450},
					value:         400,
					balanceFactor: -1,
				},
				{
					parent:        &Node[int]{value: 400},
					value:         300,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 450},
					value:         500,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 500},
					value:         460,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 500},
					value:         600,
					balanceFactor: 0,
				},
			},
		},
		{
			name:   "rotate_left_right_bf_-1",
			values: []int{500, 600, 400, 300, 450, 440},
			expectedNLRValues: []*Node[int]{
				{
					value:         450,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 450},
					value:         400,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 400},
					value:         300,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 400},
					value:         440,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 450},
					value:         500,
					balanceFactor: 1,
				},
				{
					parent:        &Node[int]{value: 500},
					value:         600,
					balanceFactor: 0,
				},
			},
		},
		{
			name:   "no_height_increase",
			values: []int{500, 400, 600, 300, 700, 200, 800, 590, 900},
			expectedNLRValues: []*Node[int]{
				{
					value:         500,
					balanceFactor: 1,
				},
				{
					parent: &Node[int]{value: 500},
					value:  300,
				},
				{
					parent: &Node[int]{value: 300},
					value:  200,
				},
				{
					parent: &Node[int]{value: 300},
					value:  400,
				},
				{
					parent: &Node[int]{value: 500},
					value:  700,
				},
				{
					parent:        &Node[int]{value: 700},
					value:         600,
					balanceFactor: -1,
				},
				{
					parent: &Node[int]{value: 600},
					value:  590,
				},
				{
					parent:        &Node[int]{value: 700},
					value:         800,
					balanceFactor: 1,
				},
				{
					parent: &Node[int]{value: 800},
					value:  900,
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			avl := NewAVL[int]()

			for _, v := range tc.values {
				avl.Insert(v)
			}

			vals := slices.Collect(avl.Traverse(NLR))

			if diff := cmp.Diff(len(tc.values), avl.Size()); diff != "" {
				t.Errorf("invalid size:\n%v", diff)
			}

			if diff := cmp.Diff(tc.expectedNLRValues, vals, nodeComparer); diff != "" {
				t.Errorf("invalid values:\n%v", diff)
			}
		})
	}
}

func TestAVL_int_Find(t *testing.T) {
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
			valuesToInsert: []int{1, 2, 3, 10, 9, 5},
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
			avl := NewAVL[int]()

			for _, v := range tc.valuesToInsert {
				avl.Insert(v)
			}

			n, found := avl.Find(tc.valueToFind)

			if diff := cmp.Diff(tc.expectedNode, n, nodeComparer); diff != "" {
				t.Errorf("invalid node:\n%v", diff)
			}

			if diff := cmp.Diff(tc.expectedFound, found); diff != "" {
				t.Errorf("invalid found:\n%v", diff)
			}
		})
	}
}

func TestAVL_int_Remove(t *testing.T) {
	type testCase struct {
		name              string
		valueToRemove     int
		valuesToInsert    []int
		expectedNLRValues []*Node[int]
	}

	testCases := []testCase{
		{
			name:           "leaf_node_left",
			valuesToInsert: []int{500, 400, 600, 300, 700, 200, 800, 590, 900},
			valueToRemove:  590,
			expectedNLRValues: []*Node[int]{
				{
					value:         500,
					balanceFactor: 1,
				},
				{
					parent:        &Node[int]{value: 500},
					value:         300,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 300},
					value:         200,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 300},
					value:         400,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 500},
					value:         700,
					balanceFactor: 1,
				},
				{
					parent:        &Node[int]{value: 700},
					value:         600,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 700},
					value:         800,
					balanceFactor: 1,
				},
				{
					parent:        &Node[int]{value: 800},
					value:         900,
					balanceFactor: 0,
				},
			},
		},
		{
			name:           "leaf_node_right",
			valuesToInsert: []int{500, 400, 600, 300, 700, 200, 800, 590, 900},
			valueToRemove:  900,
			expectedNLRValues: []*Node[int]{
				{
					value:         500,
					balanceFactor: 1,
				},
				{
					parent:        &Node[int]{value: 500},
					value:         300,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 300},
					value:         200,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 300},
					value:         400,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 500},
					value:         700,
					balanceFactor: -1,
				},
				{
					parent:        &Node[int]{value: 700},
					value:         600,
					balanceFactor: -1,
				},
				{
					parent:        &Node[int]{value: 600},
					value:         590,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 700},
					value:         800,
					balanceFactor: 0,
				},
			},
		},
		{
			name:           "leaf_node_right_left",
			valuesToInsert: []int{500, 400, 600, 300, 700, 200, 800, 650, 900},
			valueToRemove:  650,
			expectedNLRValues: []*Node[int]{
				{
					value:         500,
					balanceFactor: 1,
				},
				{
					parent:        &Node[int]{value: 500},
					value:         300,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 300},
					value:         200,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 300},
					value:         400,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 500},
					value:         700,
					balanceFactor: 1,
				},
				{
					parent:        &Node[int]{value: 700},
					value:         600,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 700},
					value:         800,
					balanceFactor: 1,
				},
				{
					parent:        &Node[int]{value: 800},
					value:         900,
					balanceFactor: 0,
				},
			},
		},
		{
			name:           "leaf_node_left_right",
			valuesToInsert: []int{500, 400, 600, 300, 700, 200, 800, 590, 750},
			valueToRemove:  750,
			expectedNLRValues: []*Node[int]{
				{
					value:         500,
					balanceFactor: 1,
				},
				{
					parent:        &Node[int]{value: 500},
					value:         300,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 300},
					value:         200,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 300},
					value:         400,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 500},
					value:         700,
					balanceFactor: -1,
				},
				{
					parent:        &Node[int]{value: 700},
					value:         600,
					balanceFactor: -1,
				},
				{
					parent:        &Node[int]{value: 600},
					value:         590,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 700},
					value:         800,
					balanceFactor: 0,
				},
			},
		},
		{
			name:           "with_left_child_only",
			valuesToInsert: []int{500, 400, 600, 300, 700, 200, 800, 650, 590, 900, 100, 580},
			valueToRemove:  590,
			expectedNLRValues: []*Node[int]{
				{
					value:         500,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 500},
					value:         300,
					balanceFactor: -1,
				},
				{
					parent:        &Node[int]{value: 300},
					value:         200,
					balanceFactor: -1,
				},
				{
					parent:        &Node[int]{value: 200},
					value:         100,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 300},
					value:         400,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 500},
					value:         700,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 700},
					value:         600,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 600},
					value:         580,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 600},
					value:         650,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 700},
					value:         800,
					balanceFactor: 1,
				},
				{
					parent:        &Node[int]{value: 800},
					value:         900,
					balanceFactor: 0,
				},
			},
		},
		{
			name:           "with_right_child_only",
			valuesToInsert: []int{500, 400, 600, 300, 700, 200, 800, 650, 590, 900, 100, 595},
			valueToRemove:  590,
			expectedNLRValues: []*Node[int]{
				{
					value:         500,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 500},
					value:         300,
					balanceFactor: -1,
				},
				{
					parent:        &Node[int]{value: 300},
					value:         200,
					balanceFactor: -1,
				},
				{
					parent:        &Node[int]{value: 200},
					value:         100,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 300},
					value:         400,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 500},
					value:         700,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 700},
					value:         600,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 600},
					value:         595,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 600},
					value:         650,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 700},
					value:         800,
					balanceFactor: 1,
				},
				{
					parent:        &Node[int]{value: 800},
					value:         900,
					balanceFactor: 0,
				},
			},
		},
		{
			name:           "with_both_children",
			valuesToInsert: []int{500, 400, 600, 300, 700, 200, 800, 650, 590, 900, 100, 595, 580},
			valueToRemove:  590,
			expectedNLRValues: []*Node[int]{
				{
					value:         500,
					balanceFactor: 1,
				},
				{
					parent:        &Node[int]{value: 500},
					value:         300,
					balanceFactor: -1,
				},
				{
					parent:        &Node[int]{value: 300},
					value:         200,
					balanceFactor: -1,
				},
				{
					parent:        &Node[int]{value: 200},
					value:         100,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 300},
					value:         400,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 500},
					value:         700,
					balanceFactor: -1,
				},
				{
					parent:        &Node[int]{value: 700},
					value:         600,
					balanceFactor: -1,
				},
				{
					parent:        &Node[int]{value: 600},
					value:         580,
					balanceFactor: 1,
				},
				{
					parent:        &Node[int]{value: 580},
					value:         595,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 600},
					value:         650,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 700},
					value:         800,
					balanceFactor: 1,
				},
				{
					parent:        &Node[int]{value: 800},
					value:         900,
					balanceFactor: 0,
				},
			},
		},
		{
			name:           "with_both_children_and_left_right_child",
			valuesToInsert: []int{500, 400, 600, 300, 700, 200, 800, 650, 590, 900, 100, 595, 580},
			valueToRemove:  600,
			expectedNLRValues: []*Node[int]{
				{
					value:         500,
					balanceFactor: 1,
				},
				{
					parent:        &Node[int]{value: 500},
					value:         300,
					balanceFactor: -1,
				},
				{
					parent:        &Node[int]{value: 300},
					value:         200,
					balanceFactor: -1,
				},
				{
					parent:        &Node[int]{value: 200},
					value:         100,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 300},
					value:         400,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 500},
					value:         700,
					balanceFactor: -1,
				},
				{
					parent:        &Node[int]{value: 700},
					value:         595,
					balanceFactor: -1,
				},
				{
					parent:        &Node[int]{value: 595},
					value:         590,
					balanceFactor: -1,
				},
				{
					parent:        &Node[int]{value: 590},
					value:         580,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 595},
					value:         650,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 700},
					value:         800,
					balanceFactor: 1,
				},
				{
					parent:        &Node[int]{value: 800},
					value:         900,
					balanceFactor: 0,
				},
			},
		},
		{
			name: "with_both_children_and_left_right_child_2",
			valuesToInsert: []int{500, 400, 600, 300, 700, 200, 800, 650, 590, 900, 100, 592, 580, 750, 901, 651,
				401, 101, 201, 591, 594, 99, 199, 399, 402, 649, 652, 749, 899, 902, 581, 202, 593},
			valueToRemove: 600,
			expectedNLRValues: []*Node[int]{
				{
					value: 500,
				},
				{
					parent:        &Node[int]{value: 500},
					value:         300,
					balanceFactor: -1,
				},
				{
					parent:        &Node[int]{value: 300},
					value:         101,
					balanceFactor: 1,
				},
				{
					parent:        &Node[int]{value: 101},
					value:         100,
					balanceFactor: -1,
				},
				{
					parent:        &Node[int]{value: 100},
					value:         99,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 101},
					value:         200,
					balanceFactor: 1,
				},
				{
					parent:        &Node[int]{value: 200},
					value:         199,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 200},
					value:         201,
					balanceFactor: 1,
				},
				{
					parent:        &Node[int]{value: 201},
					value:         202,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 300},
					value:         400,
					balanceFactor: 1,
				},
				{
					parent:        &Node[int]{value: 400},
					value:         399,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 400},
					value:         401,
					balanceFactor: 1,
				},
				{
					parent:        &Node[int]{value: 401},
					value:         402,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 500},
					value:         700,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 700},
					value:         594,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 594},
					value:         590,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 590},
					value:         580,
					balanceFactor: 1,
				},
				{
					parent:        &Node[int]{value: 580},
					value:         581,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 590},
					value:         592,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 592},
					value:         591,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 592},
					value:         593,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 594},
					value:         650,
					balanceFactor: 1,
				},
				{
					parent:        &Node[int]{value: 650},
					value:         649,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 650},
					value:         651,
					balanceFactor: 1,
				},
				{
					parent:        &Node[int]{value: 651},
					value:         652,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 700},
					value:         800,
					balanceFactor: 1,
				},
				{
					parent:        &Node[int]{value: 800},
					value:         750,
					balanceFactor: -1,
				},
				{
					parent:        &Node[int]{value: 750},
					value:         749,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 800},
					value:         900,
					balanceFactor: 1,
				},
				{
					parent:        &Node[int]{value: 900},
					value:         899,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 900},
					value:         901,
					balanceFactor: 1,
				},
				{
					parent:        &Node[int]{value: 901},
					value:         902,
					balanceFactor: 0,
				},
			},
		},
		{
			name:              "root_only_node",
			valueToRemove:     400,
			valuesToInsert:    []int{400},
			expectedNLRValues: nil,
		},
		{
			name:           "root_many_nodes",
			valueToRemove:  400,
			valuesToInsert: []int{400, 300, 500, 200, 350, 510, 100, 340, 360},
			expectedNLRValues: []*Node[int]{
				{
					value:         360,
					balanceFactor: -1,
				},
				{
					parent:        &Node[int]{value: 360},
					value:         300,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 300},
					value:         200,
					balanceFactor: -1,
				},
				{
					parent:        &Node[int]{value: 200},
					value:         100,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 300},
					value:         350,
					balanceFactor: -1,
				},
				{
					parent:        &Node[int]{value: 350},
					value:         340,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 360},
					value:         500,
					balanceFactor: 1,
				},
				{
					parent:        &Node[int]{value: 500},
					value:         510,
					balanceFactor: 0,
				},
			},
		},
		{
			name:           "rotate_left",
			valuesToInsert: []int{300, 200, 400, 500},
			valueToRemove:  200,
			expectedNLRValues: []*Node[int]{
				{
					value:         400,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 400},
					value:         300,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 400},
					value:         500,
					balanceFactor: 0,
				},
			},
		},
		{
			name:           "rotate_left_right_child_zero_bf",
			valuesToInsert: []int{500, 400, 600, 550, 700},
			valueToRemove:  400,
			expectedNLRValues: []*Node[int]{
				{
					value:         600,
					balanceFactor: -1,
				},
				{
					parent:        &Node[int]{value: 600},
					value:         500,
					balanceFactor: 1,
				},
				{
					parent:        &Node[int]{value: 500},
					value:         550,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 600},
					value:         700,
					balanceFactor: 0,
				},
			},
		},
		{
			name:           "rotate_right_left",
			valuesToInsert: []int{300, 200, 400, 350},
			valueToRemove:  200,
			expectedNLRValues: []*Node[int]{
				{
					value:         350,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 350},
					value:         300,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 350},
					value:         400,
					balanceFactor: 0,
				},
			},
		},
		{
			name:           "rotate_right",
			valuesToInsert: []int{300, 200, 400, 100},
			valueToRemove:  400,
			expectedNLRValues: []*Node[int]{
				{
					value:         200,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 200},
					value:         100,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 200},
					value:         300,
					balanceFactor: 0,
				},
			},
		},
		{
			name:           "rotate_right_left_child_zero_bf",
			valuesToInsert: []int{500, 400, 600, 300, 450},
			valueToRemove:  600,
			expectedNLRValues: []*Node[int]{
				{
					value:         400,
					balanceFactor: 1,
				},
				{
					parent:        &Node[int]{value: 400},
					value:         300,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 400},
					value:         500,
					balanceFactor: -1,
				},
				{
					parent:        &Node[int]{value: 500},
					value:         450,
					balanceFactor: 0,
				},
			},
		},
		{
			name:           "rotate_left_right",
			valuesToInsert: []int{300, 200, 400, 250},
			valueToRemove:  400,
			expectedNLRValues: []*Node[int]{
				{
					value:         250,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 250},
					value:         200,
					balanceFactor: 0,
				},
				{
					parent:        &Node[int]{value: 250},
					value:         300,
					balanceFactor: 0,
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			avl := NewAVL[int]()

			for _, v := range tc.valuesToInsert {
				avl.Insert(v)
			}

			n, ok := avl.Find(tc.valueToRemove)
			if !ok {
				t.Fatalf("node not found")
			}

			oldSize := avl.Size()
			avl.Remove(n)

			vals := slices.Collect(avl.Traverse(NLR))

			if diff := cmp.Diff(oldSize-1, avl.Size()); diff != "" {
				t.Errorf("invalid size:\n%v", diff)
			}

			if diff := cmp.Diff(tc.expectedNLRValues, vals, nodeComparer); diff != "" {
				t.Errorf("invalid values:\n%v", diff)
			}
		})
	}
}
