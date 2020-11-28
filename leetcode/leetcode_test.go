package leetcode

import (
	"testing"
)

var tree = &TreeNode{
	Val: 3,
	Left: &TreeNode{
		Val: 9,
		Left: &TreeNode{
			Val: 21,
			Left: &TreeNode{
				Val: 21,
			},
			// Right: &TreeNode{
			// 	Val: 8,
			// },
		},
		Right: &TreeNode{
			Val: 8,
		},
	},
	Right: &TreeNode{
		Val: 20,
		Left: &TreeNode{
			Val: 15,
		},
	},
}

func TestLeetcode(t *testing.T) {
	t.Log(search([]int{-1, 0, 3, 5, 9, 12}, 9))
}
