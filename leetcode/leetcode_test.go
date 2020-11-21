package leetcode

import (
	"testing"
)

func TestLeetcode(t *testing.T) {
	ln := &ListNode{
		Val: 4,
	}
	cursor := ln
	cursor.Next = &ListNode{
		Val: 2,
	}
	cursor = cursor.Next
	cursor.Next = &ListNode{
		Val: 1,
	}
	cursor = cursor.Next
	cursor.Next = &ListNode{
		Val: 3,
	}
	cursor = cursor.Next
	res := sortList(ln)
	t.Log(res)
	t.Log(res.Next)
	t.Log(res.Next.Next)
	t.Log(res.Next.Next.Next)
}
