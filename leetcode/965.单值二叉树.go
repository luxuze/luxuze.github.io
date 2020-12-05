package leetcode

// 72/72 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 45.77 % of golang submissions (2.3 MB)
/*
 * @lc app=leetcode.cn id=965 lang=golang
 *
 * [965] 单值二叉树
 *
 * https://leetcode-cn.com/problems/univalued-binary-tree/description/
 *
 * algorithms
 * Easy (68.42%)
 * Likes:    67
 * Dislikes: 0
 * Total Accepted:    22.5K
 * Total Submissions: 32.9K
 * Testcase Example:  '[1,1,1,1,1,null,1]'
 *
 * 如果二叉树每个节点都具有相同的值，那么该二叉树就是单值二叉树。
 *
 * 只有给定的树是单值二叉树时，才返回 true；否则返回 false。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 * 输入：[1,1,1,1,1,null,1]
 * 输出：true
 *
 *
 * 示例 2：
 *
 *
 *
 * 输入：[2,2,2,5,2]
 * 输出：false
 *
 *
 *
 *
 * 提示：
 *
 *
 * 给定树的节点数范围是 [1, 100]。
 * 每个节点的值都是整数，范围为 [0, 99] 。
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

func isUnivalTree(root *TreeNode) bool {
	var dfs func(tn *TreeNode) bool
	dfs = func(tn *TreeNode) bool {
		if tn == nil {
			return true
		}
		if tn.Left != nil && tn.Val != tn.Left.Val {
			return false
		}
		if tn.Right != nil && tn.Val != tn.Right.Val {
			return false
		}
		return dfs(tn.Left) && dfs(tn.Right)
	}
	return dfs(root)
}

// @lc code=end
