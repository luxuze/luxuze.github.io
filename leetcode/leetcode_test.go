package leetcode

import (
	"testing"
)

var (
	tree = treeNode.Deserialize("3,1,2")
	ln   = listNode.Deserialize("1,2,3,4")
)

func TestT(t *testing.T) {
	t.Log(combinationSum([]int{2, 3, 6, 7}, 7))
}
