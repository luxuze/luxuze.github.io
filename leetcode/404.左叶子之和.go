package leetcode

// 102/102 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 61.74 % of golang submissions (2.8 MB)
/*
 * @lc app=leetcode.cn id=404 lang=golang
 *
 * [404] 左叶子之和
 *
 * https://leetcode-cn.com/problems/sum-of-left-leaves/description/
 *
 * algorithms
 * Easy (56.31%)
 * Likes:    263
 * Dislikes: 0
 * Total Accepted:    63.5K
 * Total Submissions: 112.7K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 计算给定二叉树的所有左叶子之和。
 *
 * 示例：
 *
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 *
 * 在这个二叉树中，有两个左叶子，分别是 9 和 15，所以返回 24
 *
 *
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
func sumOfLeftLeaves(root *TreeNode) (answer int) {
	var dfs func(tn *TreeNode)
	dfs = func(tn *TreeNode) {
		if tn == nil {
			return
		}
		if tn.Left != nil && tn.Left.Left == nil && tn.Left.Right == nil {
			answer += tn.Left.Val
		}
		dfs(tn.Left)
		dfs(tn.Right)
	}
	dfs(root)
	return
}

// @lc code=end
