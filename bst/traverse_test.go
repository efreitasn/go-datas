package bst

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestTraverse(t *testing.T) {
	type testCase struct {
		name           string
		order          TraverseOrder
		values         []int
		expectedValues []*Node[int]
	}

	testCases := []testCase{
		{
			name:   "pre_order",
			order:  PreOrder,
			values: []int{500, 400, 600, 300, 700, 450, 550},
			expectedValues: []*Node[int]{
				{
					value: 500,
				},
				{
					parent: &Node[int]{value: 500},
					value:  400,
				},
				{
					parent: &Node[int]{value: 400},
					value:  300,
				},
				{
					parent: &Node[int]{value: 400},
					value:  450,
				},
				{
					parent: &Node[int]{value: 500},
					value:  600,
				},
				{
					parent: &Node[int]{value: 600},
					value:  550,
				},
				{
					parent: &Node[int]{value: 600},
					value:  700,
				},
			},
		},
		{
			name:   "in_order",
			order:  InOrder,
			values: []int{500, 400, 600, 300, 700, 450, 550},
			expectedValues: []*Node[int]{
				{
					parent: &Node[int]{value: 400},
					value:  300,
				},
				{
					parent: &Node[int]{value: 500},
					value:  400,
				},
				{
					parent: &Node[int]{value: 400},
					value:  450,
				},
				{
					value: 500,
				},
				{
					parent: &Node[int]{value: 600},
					value:  550,
				},
				{
					parent: &Node[int]{value: 500},
					value:  600,
				},
				{
					parent: &Node[int]{value: 600},
					value:  700,
				},
			},
		},
		{
			name:   "post_order",
			order:  PostOrder,
			values: []int{500, 400, 600, 300, 700, 450, 550},
			expectedValues: []*Node[int]{
				{
					parent: &Node[int]{value: 400},
					value:  300,
				},
				{
					parent: &Node[int]{value: 400},
					value:  450,
				},
				{
					parent: &Node[int]{value: 500},
					value:  400,
				},
				{
					parent: &Node[int]{value: 600},
					value:  550,
				},
				{
					parent: &Node[int]{value: 600},
					value:  700,
				},
				{
					parent: &Node[int]{value: 500},
					value:  600,
				},
				{
					value: 500,
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

			traverse(tc.order, bst.root, func(n *Node[int]) bool {
				vals = append(vals, n)

				return true
			})

			if diff := cmp.Diff(tc.expectedValues, vals, nodeComparer); diff != "" {
				t.Errorf("invalid values:\n%v", diff)
			}
		})
	}
}
