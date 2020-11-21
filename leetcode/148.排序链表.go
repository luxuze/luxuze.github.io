package leetcode

// 28/28 cases passed (32 ms)
// Your runtime beats 50.17 % of golang submissions
// Your memory usage beats 37.08 % of golang submissions (7.2 MB)

/*
 * @lc app=leetcode.cn id=148 lang=golang
 *
 * [148] 排序链表
 *
 * https://leetcode-cn.com/problems/sort-list/description/
 *
 * algorithms
 * Medium (67.74%)
 * Likes:    880
 * Dislikes: 0
 * Total Accepted:    119.1K
 * Total Submissions: 175.7K
 * Testcase Example:  '[4,2,1,3]'
 *
 * 给你链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。
 *
 * 进阶：
 *
 *
 * 你可以在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序吗？
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：head = [4,2,1,3]
 * 输出：[1,2,3,4]
 *
 *
 * 示例 2：
 *
 *
 * 输入：head = [-1,5,3,4,0]
 * 输出：[-1,0,3,4,5]
 *
 *
 * 示例 3：
 *
 *
 * 输入：head = []
 * 输出：[]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 链表中节点的数目在范围 [0, 5 * 10^4] 内
 * -10^5
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
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	left, right := splitListNode(head)
	if left != nil && left.Next != nil {
		left = sortList(left)
	}
	if right != nil && right.Next != nil {
		right = sortList(right)
	}
	return mergeListNode(left, right)
}

func splitListNode(head *ListNode) (*ListNode, *ListNode) {
	if head == nil {
		return nil, nil
	}
	slow, fast := head, head.Next
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	middle := slow.Next
	slow.Next = nil
	return head, middle
}

func mergeListNode(ln1, ln2 *ListNode) *ListNode {
	if ln1 == nil {
		return ln2
	} else if ln2 == nil {
		return ln1
	}

	head := new(ListNode)
	cursor := head
	for ln1 != nil && ln2 != nil {
		if ln1.Val < ln2.Val {
			cursor.Next = ln1
			ln1 = ln1.Next
		} else {
			cursor.Next = ln2
			ln2 = ln2.Next
		}
		cursor = cursor.Next
	}
	if ln1 != nil {
		cursor.Next = ln1
	} else if ln2 != nil {
		cursor.Next = ln2
	}
	return head.Next
}

// @lc code=end
