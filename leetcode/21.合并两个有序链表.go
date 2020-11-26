package leetcode

// 208/208 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 85.55 % of golang submissions (2.5 MB)
/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 *
 * https://leetcode-cn.com/problems/merge-two-sorted-lists/description/
 *
 * algorithms
 * Easy (64.77%)
 * Likes:    1399
 * Dislikes: 0
 * Total Accepted:    417.9K
 * Total Submissions: 645K
 * Testcase Example:  '[1,2,4]\n[1,3,4]'
 *
 * 将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
 *
 *
 *
 * 示例：
 *
 * 输入：1->2->4, 1->3->4
 * 输出：1->1->2->3->4->4
 *
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
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	ans := &ListNode{}
	cursor := ans
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			cursor.Next = l1
			l1 = l1.Next
		} else {
			cursor.Next = l2
			l2 = l2.Next
		}
		cursor = cursor.Next
	}
	if l1 != nil {
		cursor.Next = l1
	} else if l2 != nil {
		cursor.Next = l2
	}
	return ans.Next
}

// @lc code=end
