package leetcode

// 228/228 cases passed (4 ms)
// Your runtime beats 99.46 % of golang submissions
// Your memory usage beats 45.71 % of golang submissions (6 MB)
/*
 * @lc app=leetcode.cn id=110 lang=golang
 *
 * [110] 平衡二叉树
 *
 * https://leetcode-cn.com/problems/balanced-binary-tree/description/
 *
 * algorithms
 * Easy (54.89%)
 * Likes:    527
 * Dislikes: 0
 * Total Accepted:    152.3K
 * Total Submissions: 277.4K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给定一个二叉树，判断它是否是高度平衡的二叉树。
 *
 * 本题中，一棵高度平衡二叉树定义为：
 *
 *
 * 一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1 。
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：root = [3,9,20,null,null,15,7]
 * 输出：true
 *
 *
 * 示例 2：
 *
 *
 * 输入：root = [1,2,2,3,3,null,null,4,4]
 * 输出：false
 *
 *
 * 示例 3：
 *
 *
 * 输入：root = []
 * 输出：true
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树中的节点数在范围 [0, 5000] 内
 * -10^4
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
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if isBalanced_Abs(isBalanced_MaxDepth(root.Left)-isBalanced_MaxDepth(root.Right)) > 1 ||
		!isBalanced(root.Left) ||
		!isBalanced(root.Right) {
		return false
	}

	return true
}

func isBalanced_MaxDepth(tn *TreeNode) int {
	if tn == nil {
		return 0
	}
	return isBalanced_Max(isBalanced_MaxDepth(tn.Left), isBalanced_MaxDepth(tn.Right)) + 1
}

func isBalanced_Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func isBalanced_Abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

// @lc code=end
