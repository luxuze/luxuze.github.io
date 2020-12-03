package leetcode

// 133/133 cases passed (20 ms)
// Your runtime beats 40.54 % of golang submissions
// Your memory usage beats 14.89 % of golang submissions (6.2 MB)
import "sort"

/*
 * @lc app=leetcode.cn id=23 lang=golang
 *
 * [23] 合并K个排序链表
 *
 * https://leetcode-cn.com/problems/merge-k-sorted-lists/description/
 *
 * algorithms
 * Hard (48.21%)
 * Likes:    748
 * Dislikes: 0
 * Total Accepted:    135.2K
 * Total Submissions: 259.8K
 * Testcase Example:  '[[1,4,5],[1,3,4],[2,6]]'
 *
 * 合并 k 个排序链表，返回合并后的排序链表。请分析和描述算法的复杂度。
 *
 * 示例:
 *
 * 输入:
 * [
 * 1->4->5,
 * 1->3->4,
 * 2->6
 * ]
 * 输出: 1->1->2->3->4->4->5->6
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
func mergeKLists(lists []*ListNode) *ListNode {
	var head = &ListNode{}
	cursor := head
	var arr []*ListNode
	for i := range lists {
		for lists[i] != nil {
			arr = append(arr, lists[i])
			lists[i] = lists[i].Next
		}
	}
	sort.SliceStable(arr, func(a, b int) bool { return arr[a].Val < arr[b].Val })
	for i := range arr {
		cursor.Next = arr[i]
		cursor = cursor.Next
	}
	cursor.Next = nil
	return head.Next
}

// @lc code=end
