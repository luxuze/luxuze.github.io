package leetcode

import (
	"testing"
)

var (
	tree = Deserialize("3,9,20,null,null,15,7")
)

func TestLeetcode(t *testing.T) {
	answer := levelOrderBottom(nil)
	t.Log(answer)
}
