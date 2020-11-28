package leetcode

// 52/52 cases passed (268 ms)
// Your runtime beats 12.44 % of golang submissions
// Your memory usage beats 10.52 % of golang submissions (20.1 MB)

/*
 * @lc app=leetcode.cn id=111 lang=golang
 *
 * [111] 二叉树的最小深度
 *
 * https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/description/
 *
 * algorithms
 * Easy (45.35%)
 * Likes:    407
 * Dislikes: 0
 * Total Accepted:    157.2K
 * Total Submissions: 346.1K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给定一个二叉树，找出其最小深度。
 *
 * 最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
 *
 * 说明：叶子节点是指没有子节点的节点。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：root = [3,9,20,null,null,15,7]
 * 输出：2
 *
 *
 * 示例 2：
 *
 *
 * 输入：root = [2,null,3,null,4,null,5,null,6]
 * 输出：5
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树中节点数的范围在 [0, 10^5] 内
 * -1000
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
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var min int
	var dfs func(*TreeNode, int)
	dfs = func(tn *TreeNode, level int) {
		if tn.Left == nil && tn.Right == nil {
			if min == 0 {
				min = level
				return
			}
			if level < min {
				min = level
			}
			return
		}
		if tn.Left != nil {
			dfs(tn.Left, level+1)
		}
		if tn.Right != nil {
			dfs(tn.Right, level+1)
		}
	}
	dfs(root, 1)
	return min
}

// @lc code=end
