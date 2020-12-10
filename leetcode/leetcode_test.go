package leetcode

import (
	"testing"
)

var (
	tree = Deserialize("3,1,2")
	// ln   = &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	ln = &ListNode{Val: 3, Next: &ListNode{Val: 4}}
)

func TestT(t *testing.T) {
	answer := reverseBetween(ln, 1, 2)
	t.Log(answer)
	t.Log(answer.Next)
	// t.Log(answer.Next.Next)
}
