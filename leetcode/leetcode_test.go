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
	trie := Trie_Constructor()
	trie.Insert("apple")
	t.Log(trie)
}
