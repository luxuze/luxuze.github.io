package leetcode

import (
	"testing"
)

var (
	tree = Deserialize("3,1,2")
)

func TestT(t *testing.T) {
	answer := Serialize(buildTree([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3}))
	t.Log(answer)
}
