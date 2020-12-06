package leetcode

import (
	"testing"
)

var (
	tree = Deserialize("3,1,2")
	ln   = &ListNode{Val: 4, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: &ListNode{Val: 3}}}}
)

func TestT(t *testing.T) {
	answer := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(answer, 3)
	t.Log(answer)
}
