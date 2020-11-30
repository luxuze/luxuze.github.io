package leetcode

// 75/75 cases passed (8 ms)
// Your runtime beats 74.86 % of golang submissions
// Your memory usage beats 26.65 % of golang submissions (5.5 MB)
import "math"

/*
 * @lc app=leetcode.cn id=98 lang=golang
 *
 * [98] 验证二叉搜索树
 *
 * https://leetcode-cn.com/problems/validate-binary-search-tree/description/
 *
 * algorithms
 * Medium (32.83%)
 * Likes:    844
 * Dislikes: 0
 * Total Accepted:    194.8K
 * Total Submissions: 593.3K
 * Testcase Example:  '[2,1,3]'
 *
 * 给定一个二叉树，判断其是否是一个有效的二叉搜索树。
 *
 * 假设一个二叉搜索树具有如下特征：
 *
 *
 * 节点的左子树只包含小于当前节点的数。
 * 节点的右子树只包含大于当前节点的数。
 * 所有左子树和右子树自身必须也是二叉搜索树。
 *
 *
 * 示例 1:
 *
 * 输入:
 * ⁠   2
 * ⁠  / \
 * ⁠ 1   3
 * 输出: true
 *
 *
 * 示例 2:
 *
 * 输入:
 * ⁠   5
 * ⁠  / \
 * ⁠ 1   4
 * / \
 * 3   6
 * 输出: false
 * 解释: 输入为: [5,1,4,null,null,3,6]。
 * 根节点的值为 5 ，但是其右子节点值为 4 。
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

func isValidBST(root *TreeNode) bool {
	var dfs func(tn *TreeNode, min, max int) bool
	dfs = func(tn *TreeNode, min, max int) bool {
		if tn == nil {
			return true
		}
		if tn.Val <= min || tn.Val >= max {
			return false
		}
		return dfs(tn.Left, min, tn.Val) && dfs(tn.Right, tn.Val, max)
	}
	return dfs(root, math.MinInt64, math.MaxInt64)
}

// @lc code=end
