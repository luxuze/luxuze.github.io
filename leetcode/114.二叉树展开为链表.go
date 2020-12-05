package leetcode

// 225/225 cases passed (4 ms)
// Your runtime beats 74.19 % of golang submissions
// Your memory usage beats 8.64 % of golang submissions (3.4 MB)
/*
 * @lc app=leetcode.cn id=114 lang=golang
 *
 * [114] 二叉树展开为链表
 *
 * https://leetcode-cn.com/problems/flatten-binary-tree-to-linked-list/description/
 *
 * algorithms
 * Medium (71.48%)
 * Likes:    648
 * Dislikes: 0
 * Total Accepted:    96.2K
 * Total Submissions: 134.6K
 * Testcase Example:  '[1,2,5,3,4,null,6]'
 *
 * 给定一个二叉树，原地将它展开为一个单链表。
 *
 *
 *
 * 例如，给定二叉树
 *
 * ⁠   1
 * ⁠  / \
 * ⁠ 2   5
 * ⁠/ \   \
 * 3   4   6
 *
 * 将其展开为：
 *
 * 1
 * ⁠\
 * ⁠ 2
 * ⁠  \
 * ⁠   3
 * ⁠    \
 * ⁠     4
 * ⁠      \
 * ⁠       5
 * ⁠        \
 * ⁠         6
 *
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	var preorderTraversal func(tn *TreeNode) []*TreeNode
	preorderTraversal = func(tn *TreeNode) (list []*TreeNode) {
		if tn != nil {
			list = append(list, tn)
			list = append(list, preorderTraversal(tn.Left)...)
			list = append(list, preorderTraversal(tn.Right)...)
		}
		return
	}
	list := preorderTraversal(root)
	for i := 1; i < len(list); i++ {
		preTN, currentTN := list[i-1], list[i]
		preTN.Left, preTN.Right = nil, currentTN
	}
}

// @lc code=end
