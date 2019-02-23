package leetcode

import "testing"

func TestIsCousinsCase1(t *testing.T) {
	tree := &TreeNode{
		1,
		&TreeNode{
			2, &TreeNode{4, nil, nil}, nil,
		},
		&TreeNode{3, nil, nil},
	}
	runTest(t, tree, 4, 3, false)
}

func TestIsCousinsCase2(t *testing.T) {
	tree := &TreeNode{
		1,
		&TreeNode{
			2, nil, &TreeNode{4, nil, nil},
		},
		&TreeNode{
			3, nil, &TreeNode{5, nil, nil},
		},
	}
	runTest(t, tree, 5, 4, true)
}

func TestIsCousinsCase3(t *testing.T) {
	tree := &TreeNode{
		1,
		&TreeNode{
			2, nil, &TreeNode{4, nil, nil},
		},
		&TreeNode{3, nil, nil},
	}
	runTest(t, tree, 2, 3, false)
}

func runTest(t *testing.T, tree *TreeNode, x int, y int, expected bool) {
	isCousin := IsCousins(tree, x, y)
	if isCousin != expected {
		t.Errorf("Expected %t but instead was %t", expected, isCousin)
	}
}
