package leetcode

// 44/44 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 96.01 % of golang submissions (2.1 MB)
/*
 * @lc app=leetcode.cn id=92 lang=golang
 *
 * [92] 反转链表 II
 *
 * https://leetcode-cn.com/problems/reverse-linked-list-ii/description/
 *
 * algorithms
 * Medium (48.28%)
 * Likes:    434
 * Dislikes: 0
 * Total Accepted:    62.6K
 * Total Submissions: 123.2K
 * Testcase Example:  '[1,2,3,4,5]\n2\n4'
 *
 * 反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。
 *
 * 说明:
 * 1 ≤ m ≤ n ≤ 链表长度。
 *
 * 示例:
 *
 * 输入: 1->2->3->4->5->NULL, m = 2, n = 4
 * 输出: 1->4->3->2->5->NULL
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
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil {
		return head
	}
	var truePrev *ListNode
	middleTail := head
	for m > 1 {
		truePrev = middleTail
		middleTail = middleTail.Next
		m--
		n--
	}
	tmpPrev := truePrev
	current := middleTail
	for ; n > 0; n-- {
		next := current.Next
		current.Next = tmpPrev
		tmpPrev = current
		current = next
	}
	middleTail.Next = current
	if truePrev == nil {
		head = tmpPrev
	} else {
		truePrev.Next = tmpPrev
	}
	return head
}

// @lc code=end
