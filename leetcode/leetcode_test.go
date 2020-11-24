package leetcode

import (
	"testing"
)

var tree = &TreeNode{
	Val: 1,
	Left: &TreeNode{
		Val: 2,
	},
	Right: &TreeNode{
		Val: 3,
	},
}

func TestLeetcode(t *testing.T) {
	t.Log(checkPossibility([]int{1, 4, 1, 2}))
}
