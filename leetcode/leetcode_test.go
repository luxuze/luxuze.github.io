package leetcode

import (
	"testing"
)

var (
	tree = treeNode.Deserialize("-2,null,-3")
	ln   = listNode.Deserialize("1,2,2,1")
)

func TestT(t *testing.T) {
	l := []int{-10, -3, 0, 5, 9}
	t.Log(sortedArrayToBST(l))
}
