package leetcode

// 27/27 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 98.06 % of golang submissions (2.5 MB)
/*
 * @lc app=leetcode.cn id=206 lang=golang
 *
 * [206] 反转链表
 *
 * https://leetcode-cn.com/problems/reverse-linked-list/description/
 *
 * algorithms
 * Easy (66.04%)
 * Likes:    1113
 * Dislikes: 0
 * Total Accepted:    288.2K
 * Total Submissions: 412K
 * Testcase Example:  '[1,2,3,4,5]'
 *
 * 反转一个单链表。
 *
 * 示例:
 *
 * 输入: 1->2->3->4->5->NULL
 * 输出: 5->4->3->2->1->NULL
 *
 * 进阶:
 * 你可以迭代或递归地反转链表。你能否用两种方法解决这道题？
 *
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
// 	if head == nil || head.Next == nil {
// 		return head
// 	}
// 	newHead := reverseList(head)
// 	head.Next.Next = head.Next
// 	head.Next = nil
// 	return newHead
// }
func reverseList(head *ListNode) *ListNode {
	var current, next *ListNode = nil, head
	for next != nil {
		newNext := next.Next
		next.Next = current
		current = next
		next = newNext
	}
	return current
}

// @lc code=end
