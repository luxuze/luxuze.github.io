package leetcode

import (
	"testing"
)

var (
	tree = treeNode.Deserialize("3,1,2")
	ln   = listNode.Deserialize("1,2,3,4")
)

func TestT(t *testing.T) {
	t.Log(combinationSum2([]int{3, 1, 3, 5, 1, 1}, 8))
}
