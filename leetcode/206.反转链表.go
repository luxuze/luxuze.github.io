package leetcode

// 27/27 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 100 % of golang submissions (2.5 MB)
/*
 * @lc app=leetcode.cn id=206 lang=golang
 *
 * [206] 反转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// func reverseList(head *ListNode) *ListNode {
// 	var newHead *ListNode
// 	for head != nil {
// 		next := head.Next
// 		head.Next = newHead
// 		newHead = head
// 		head = next
// 	}
// 	return newHead
// }

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

// @lc code=end
