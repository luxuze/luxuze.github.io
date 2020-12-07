package leetcode

import (
	"testing"
)

var (
	tree = Deserialize("3,1,2")
	ln   = &ListNode{Val: 4, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: &ListNode{Val: 3}}}}
)

func TestT(t *testing.T) {
	answer := uniquePaths(3, 3)
	t.Log(answer)
}
