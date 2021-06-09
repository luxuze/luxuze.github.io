package leetcode

import (
	"testing"
)

var (
	tree = treeNode.Deserialize("-2,null,-3")
	ln   = listNode.Deserialize("1,2,2,1")
)

func TestT(t *testing.T) {
	l := removeElements(ln, 2)
	t.Log(ln.Serialize(l))
}
