package leetcode

import (
	"testing"
)

func TestLeetcode(t *testing.T) {
	tree := Deserialize("71,62,84,14,null,null,88,null,null,null,null")
	t.Log(minDiffInBST(tree))
}
