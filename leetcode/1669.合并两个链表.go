package leetcode

// 60/60 cases passed (124 ms)
// Your runtime beats 81.08 % of golang submissions
// Your memory usage beats 75.68 % of golang submissions (7.6 MB)
/*
 * @lc app=leetcode.cn id=1669 lang=golang
 *
 * [1669] 合并两个链表
 *
 * https://leetcode-cn.com/problems/merge-in-between-linked-lists/description/
 *
 * algorithms
 * Medium (80.93%)
 * Likes:    1
 * Dislikes: 0
 * Total Accepted:    2.1K
 * Total Submissions: 2.6K
 * Testcase Example:  '[0,1,2,3,4,5]\n3\n4\n[1000000,1000001,1000002]'
 *
 * 给你两个链表 list1 和 list2 ，它们包含的元素分别为 n 个和 m 个。
 *
 * 请你将 list1 中第 a 个节点到第 b 个节点删除，并将list2 接在被删除节点的位置。
 *
 * 下图中蓝色边和节点展示了操作后的结果：
 *
 * 请你返回结果链表的头指针。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 *
 * 输入：list1 = [0,1,2,3,4,5], a = 3, b = 4, list2 = [1000000,1000001,1000002]
 * 输出：[0,1,2,1000000,1000001,1000002,5]
 * 解释：我们删除 list1 中第三和第四个节点，并将 list2 接在该位置。上图中蓝色的边和节点为答案链表。
 *
 *
 * 示例 2：
 *
 *
 * 输入：list1 = [0,1,2,3,4,5,6], a = 2, b = 5, list2 =
 * [1000000,1000001,1000002,1000003,1000004]
 * 输出：[0,1,1000000,1000001,1000002,1000003,1000004,6]
 * 解释：上图中蓝色的边和节点为答案链表。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 3
 * 1
 * 1
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
func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	cursor := list1
	for i := 1; i < a; i++ {
		cursor = cursor.Next
	}
	tmpHead := cursor.Next
	cursor.Next = list2
	for list2.Next != nil {
		list2 = list2.Next
	}
	for i := 0; i < b-a+1; i++ {
		tmpHead = tmpHead.Next
	}
	list2.Next = tmpHead
	return list1
}

// @lc code=end
