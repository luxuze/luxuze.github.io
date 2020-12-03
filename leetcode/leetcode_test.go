package leetcode

import (
	"testing"
)

var (
	tree = Deserialize("71,62,84,14,null,null,88,null,null,null,null")
)

func TestLeetcode(t *testing.T) {
	merge(
		[]int{4, 0, 0, 0, 0, 0}, 1,
		[]int{1, 2, 3, 5, 6}, 5)

	// t.Log(answer)
}
