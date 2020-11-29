package leetcode

// 195/195 cases passed (4 ms)
// Your runtime beats 75.53 % of golang submissions
// Your memory usage beats 63.54 % of golang submissions (2.9 MB)
/*
 * @lc app=leetcode.cn id=101 lang=golang
 *
 * [101] 对称二叉树
 *
 * https://leetcode-cn.com/problems/symmetric-tree/description/
 *
 * algorithms
 * Easy (53.18%)
 * Likes:    1133
 * Dislikes: 0
 * Total Accepted:    234.7K
 * Total Submissions: 441.2K
 * Testcase Example:  '[1,2,2,3,4,4,3]'
 *
 * 给定一个二叉树，检查它是否是镜像对称的。
 *
 *
 *
 * 例如，二叉树 [1,2,2,3,4,4,3] 是对称的。
 *
 * ⁠   1
 * ⁠  / \
 * ⁠ 2   2
 * ⁠/ \ / \
 * 3  4 4  3
 *
 *
 *
 *
 * 但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:
 *
 * ⁠   1
 * ⁠  / \
 * ⁠ 2   2
 * ⁠  \   \
 * ⁠  3    3
 *
 *
 *
 *
 * 进阶：
 *
 * 你可以运用递归和迭代两种方法解决这个问题吗？
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
func isSymmetric(root *TreeNode) bool {
	var dfs func(*TreeNode, *TreeNode) bool
	dfs = func(tn1, tn2 *TreeNode) bool {
		if tn1 == nil && tn2 == nil {
			return true
		}
		if tn1 == nil || tn2 == nil {
			return false
		}
		return tn1.Val == tn2.Val && dfs(tn1.Left, tn2.Right) && dfs(tn1.Right, tn2.Left)
	}
	return dfs(root, root)
}

// @lc code=end
