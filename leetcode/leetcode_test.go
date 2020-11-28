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
	t.Log(addBinary("1111", "1111"))
}
