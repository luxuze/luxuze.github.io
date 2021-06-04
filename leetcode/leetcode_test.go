package leetcode

import (
	"testing"
)

var (
	tree = treeNode.Deserialize("3,1,2")
	ln   = listNode.Deserialize("1,2,3,4")
)

func TestT(t *testing.T) {
	arr := []int{1, 2, 2, 3, 4, 4, 5}
	n := removeDuplicates(arr)
	t.Log(n, arr[:n])
}
