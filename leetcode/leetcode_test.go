package leetcode

import (
	"testing"
)

var (
	tree = Deserialize("71,62,84,14,null,null,88,null,null,null,null")
)

func TestLeetcode(t *testing.T) {
	answer := kthSmallest([][]int{{1, 3, 5}, {6, 7, 12}, {11, 14, 14}}, 3)

	t.Log(answer)
}
