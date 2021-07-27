package leetcode

// 208/208 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 99.13 % of golang submissions (2.3 MB)
import "strconv"

/*
 * @lc app=leetcode.cn id=257 lang=golang
 *
 * [257] 二叉树的所有路径
 *
 * https://leetcode-cn.com/problems/binary-tree-paths/description/
 *
 * algorithms
 * Easy (67.76%)
 * Likes:    543
 * Dislikes: 0
 * Total Accepted:    127.9K
 * Total Submissions: 188.6K
 * Testcase Example:  '[1,2,3,null,5]'
 *
 * 给定一个二叉树，返回所有从根节点到叶子节点的路径。
 *
 * 说明: 叶子节点是指没有子节点的节点。
 *
 * 示例:
 *
 * 输入:
 *
 * ⁠  1
 * ⁠/   \
 * 2     3
 * ⁠\
 * ⁠ 5
 *
 * 输出: ["1->2->5", "1->3"]
 *
 * 解释: 所有根节点到叶子节点的路径为: 1->2->5, 1->3
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
func binaryTreePaths(root *TreeNode) (ans []string) {
	m := "->"
	var dfs func(tn *TreeNode, path string)
	dfs = func(tn *TreeNode, path string) {
		if tn == nil {
			return
		}
		item := strconv.Itoa(tn.Val)
		if path != "" {
			item = m + item
		}
		path += item
		if tn.Left == nil && tn.Right == nil {
			ans = append(ans, path)
			return
		}
		if tn.Left != nil {
			dfs(tn.Left, path)
		}
		if tn.Right != nil {
			dfs(tn.Right, path)
		}
	}
	dfs(root, "")
	return
}

// @lc code=end
