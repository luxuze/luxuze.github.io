package leetcode

import (
	"testing"
)

var tree = &TreeNode{
	Val: 10,
	Left: &TreeNode{
		Val: 5,
	},
	Right: &TreeNode{
		Val: 15,
		Left: &TreeNode{
			Val: 6,
		},
		Right: &TreeNode{
			Val: 20,
		},
	},
}

func TestLeetcode(t *testing.T) {
	t.Log(
		isValidBST(tree),
	)
}
