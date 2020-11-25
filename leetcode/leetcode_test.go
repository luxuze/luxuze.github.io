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
	t.Log(merge([][]int{{1, 4}, {2, 3}}))
	// t.Log(merge([][]int{{1, 3}, {2, 4}}))
	// t.Log(merge([][]int{{1, 2}, {2, 3}, {5, 6}}))
	// t.Log(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
}
