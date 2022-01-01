package leetcode

import (
	"testing"
)

var (
	tree = treeNode.Deserialize("1,2,3,null,5")
	ln   = listNode.Deserialize("1,2,2,1")
)

func TestT(t *testing.T) {
	t.Log(firstUniqChar("leetcode"))
	t.Log(firstUniqChar("lloleetcode"))
}
