package leetcode

import (
	"testing"
)

var (
	tree = Deserialize("1,null,2,3,null")
)

func TestT(t *testing.T) {
	answer := preorderTraversal(tree)
	t.Log(answer)
}
