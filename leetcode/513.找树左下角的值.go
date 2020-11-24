package leetcode

// 74/74 cases passed (8 ms)
// Your runtime beats 74.4 % of golang submissions
// Your memory usage beats 77.97 % of golang submissions (5.5 MB)
/*
 * @lc app=leetcode.cn id=513 lang=golang
 *
 * [513] 找树左下角的值
 *
 * https://leetcode-cn.com/problems/find-bottom-left-tree-value/description/
 *
 * algorithms
 * Medium (71.85%)
 * Likes:    132
 * Dislikes: 0
 * Total Accepted:    24.3K
 * Total Submissions: 33.8K
 * Testcase Example:  '[2,1,3]'
 *
 * 给定一个二叉树，在树的最后一行找到最左边的值。
 *
 * 示例 1:
 *
 *
 * 输入:
 *
 * ⁠   2
 * ⁠  / \
 * ⁠ 1   3
 *
 * 输出:
 * 1
 *
 *
 *
 *
 * 示例 2:
 *
 *
 * 输入:
 *
 * ⁠       1
 * ⁠      / \
 * ⁠     2   3
 * ⁠    /   / \
 * ⁠   4   5   6
 * ⁠      /
 * ⁠     7
 *
 * 输出:
 * 7
 *
 *
 *
 *
 * 注意: 您可以假设树（即给定的根节点）不为 NULL。
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
func findBottomLeftValue(root *TreeNode) int {
	if root.Left == nil && root.Right == nil {
		return root.Val
	}
	var max, val int
	var dfs func(root *TreeNode, level int)

	dfs = func(root *TreeNode, level int) {
		if root.Left == nil && max < level {
			max = level
			val = root.Val
		}
		if root.Left != nil {
			dfs(root.Left, level+1)
		}
		if root.Right != nil {
			dfs(root.Right, level+1)
		}
	}

	dfs(root, 0)
	return val
}

// @lc code=end
