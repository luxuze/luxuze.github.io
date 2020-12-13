package leetcode

// 55/55 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 58.61 % of golang submissions (2.1 MB)
/*
 * @lc app=leetcode.cn id=24 lang=golang
 *
 * [24] 两两交换链表中的节点
 *
 * https://leetcode-cn.com/problems/swap-nodes-in-pairs/description/
 *
 * algorithms
 * Medium (68.63%)
 * Likes:    759
 * Dislikes: 0
 * Total Accepted:    202.5K
 * Total Submissions: 294.9K
 * Testcase Example:  '[1,2,3,4]'
 *
 * 给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
 *
 * 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：head = [1,2,3,4]
 * 输出：[2,1,4,3]
 *
 *
 * 示例 2：
 *
 *
 * 输入：head = []
 * 输出：[]
 *
 *
 * 示例 3：
 *
 *
 * 输入：head = [1]
 * 输出：[1]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 链表中节点的数目在范围 [0, 100] 内
 * 0
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
// func swapPairs(head *ListNode) *ListNode {
// 	var list []*ListNode
// 	for head != nil {
// 		list = append(list, head)
// 		head = head.Next
// 	}
// 	switch len(list) {
// 	case 0:
// 		return nil
// 	case 1:
// 		return list[0]
// 	}
// 	answer := &ListNode{}
// 	cursor := answer
// 	for i, j := 0, 1; ; {
// 		cursor.Next = list[j]
// 		cursor = cursor.Next
// 		cursor.Next = list[i]
// 		cursor = cursor.Next
// 		i += 2
// 		j += 2
// 		if i >= len(list) {
// 			break
// 		}
// 		if j >= len(list) {
// 			cursor.Next = list[i]
// 			cursor = cursor.Next
// 			break
// 		}
// 	}
// 	cursor.Next = nil
// 	return answer.Next
// }
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := head.Next
	head.Next = swapPairs(newHead.Next)
	newHead.Next = head
	return newHead
}

// @lc code=end
