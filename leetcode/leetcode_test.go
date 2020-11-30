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
			Val: 6,
		},
		Right: &TreeNode{
			Val: 7,
		},
	},
}

func TestLeetcode(t *testing.T) {
	t.Log(Serialize(tree))
	tn := Deserialize(Serialize(tree))
	t.Log(Serialize(tn))
}
