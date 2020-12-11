package leetcode

// 13/13 cases passed (8 ms)
// Your runtime beats 98.41 % of golang submissions
// Your memory usage beats 30.49 % of golang submissions (5.9 MB)
/*
 * @lc app=leetcode.cn id=143 lang=golang
 *
 * [143] 重排链表
 *
 * https://leetcode-cn.com/problems/reorder-list/description/
 *
 * algorithms
 * Medium (59.66%)
 * Likes:    473
 * Dislikes: 0
 * Total Accepted:    73.5K
 * Total Submissions: 123.2K
 * Testcase Example:  '[1,2,3,4]'
 *
 * 给定一个单链表 L：L0→L1→…→Ln-1→Ln ，
 * 将其重新排列后变为： L0→Ln→L1→Ln-1→L2→Ln-2→…
 *
 * 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
 *
 * 示例 1:
 *
 * 给定链表 1->2->3->4, 重新排列为 1->4->2->3.
 *
 * 示例 2:
 *
 * 给定链表 1->2->3->4->5, 重新排列为 1->5->2->4->3.
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
func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	var list []*ListNode
	for ; head != nil; head = head.Next {
		list = append(list, head)
	}
	i, j := 0, len(list)-1
	for i < j {
		list[i].Next = list[j]
		i++
		if i == j {
			break
		}
		list[j].Next = list[i]
		j--
	}
	list[i].Next = nil
}

// @lc code=end
