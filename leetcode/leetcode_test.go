package leetcode

import (
	"testing"
)

var (
	tree = Deserialize("3,1,2")
)

func TestT(t *testing.T) {
	answer := Serialize(constructFromPrePost([]int{1, 2, 4, 5, 3, 6, 7}, []int{4, 5, 2, 6, 7, 3, 1}))
	t.Log(answer)
}
