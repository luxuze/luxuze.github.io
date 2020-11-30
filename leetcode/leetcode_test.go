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
		Left: &TreeNode{
			Val: 4,
		},
		Right: &TreeNode{
			Val: 5,
		},
	},
}

func TestLeetcode(t *testing.T) {
	t.Log(
		Serialize(tree),
	)
	t.Log(
		Deserialize("1,2,null,null,3,4,null,null,5,null,null"),
	)
}
