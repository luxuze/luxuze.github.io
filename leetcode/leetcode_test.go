package leetcode

import (
	"testing"
)

var (
	tree = treeNode.Deserialize("-2,null,-3")
	ln   = listNode.Deserialize("1,2,3,4")
)

func TestT(t *testing.T) {
	t.Log(hasPathSum(tree, -5))
}
