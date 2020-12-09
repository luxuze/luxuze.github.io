package leetcode

import (
	"testing"
)

var (
	tree = Deserialize("3,1,2")
	ln   = &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 3}}}}}
)

func TestT(t *testing.T) {
	answer := deleteDuplicates(ln)
	t.Log(answer.Val)
}

func TestTmp(t *testing.T) {
	var answer int
	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	var s string = "abbcd"
	for i := range s {
		mp := make(map[byte]int, len(s))
		for j := i; j < len(s); j++ {
			if _, ok := mp[s[j]]; !ok {
				mp[s[j]] = 1
				answer = max(answer, j-i+1)
				continue
			}
			break
		}
	}
	t.Log(answer)
}
