package leetcode

// 1563/1563 cases passed (16 ms)
// Your runtime beats 44.1 % of golang submissions
// Your memory usage beats 17.24 % of golang submissions (5 MB)

/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addNode(l1 *ListNode, l2 *ListNode, jw *int) *ListNode {
	val := l1.Val + l2.Val + *jw
	if val >= 10 {
		*jw = 1
		return &ListNode{Val: val - 10}
	}
	*jw = 0
	return &ListNode{Val: val}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var jw int
	result := new(ListNode)
	tmp := result

	for {
		tmp.Next = addNode(l1, l2, &jw)
		tmp = tmp.Next
		if l1.Next == nil && l2.Next != nil {
			l1.Next = &ListNode{Val: 0}
		} else if l1.Next != nil && l2.Next == nil {
			l2.Next = &ListNode{Val: 0}
		}
		if l1.Next == nil {
			if jw != 0 {
				tmp.Next = &ListNode{Val: 1}
			}
			break
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return result.Next
}

// @lc code=end
