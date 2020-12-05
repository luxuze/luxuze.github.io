package leetcode

// 68/68 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 82.13 % of golang submissions (2 MB)
/*
 * @lc app=leetcode.cn id=145 lang=golang
 *
 * [145] 二叉树的后序遍历
 *
 * https://leetcode-cn.com/problems/binary-tree-postorder-traversal/description/
 *
 * algorithms
 * Medium (73.66%)
 * Likes:    483
 * Dislikes: 0
 * Total Accepted:    168K
 * Total Submissions: 228K
 * Testcase Example:  '[1,null,2,3]'
 *
 * 给定一个二叉树，返回它的 后序 遍历。
 *
 * 示例:
 *
 * 输入: [1,null,2,3]
 * ⁠  1
 * ⁠   \
 * ⁠    2
 * ⁠   /
 * ⁠  3
 *
 * 输出: [3,2,1]
 *
 * 进阶: 递归算法很简单，你可以通过迭代算法完成吗？
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
func postorderTraversal(root *TreeNode) []int {
	var answer []int
	var dfs func(tn *TreeNode)
	dfs = func(tn *TreeNode) {
		if tn == nil {
			return
		}
		dfs(tn.Left)
		dfs(tn.Right)
		answer = append(answer, tn.Val)
	}
	dfs(root)
	return answer
}

// @lc code=end
