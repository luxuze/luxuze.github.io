package leetcode

// 34/34 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 87.88 % of golang submissions (2.8 MB)
/*
 * @lc app=leetcode.cn id=102 lang=golang
 *
 * [102] 二叉树的层序遍历
 *
 * https://leetcode-cn.com/problems/binary-tree-level-order-traversal/description/
 *
 * algorithms
 * Medium (63.75%)
 * Likes:    708
 * Dislikes: 0
 * Total Accepted:    221.9K
 * Total Submissions: 348.1K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。
 *
 *
 *
 * 示例：
 * 二叉树：[3,9,20,null,null,15,7],
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 *
 *
 * 返回其层次遍历结果：
 *
 * [
 * ⁠ [3],
 * ⁠ [9,20],
 * ⁠ [15,7]
 * ]
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
func levelOrder(root *TreeNode) [][]int {
	answer := [][]int{}
	if root == nil {
		return answer
	}
	level := []*TreeNode{root}
	for i := 0; len(level) > 0; i++ {
		answer = append(answer, []int{})
		var tmp []*TreeNode
		for j := 0; j < len(level); j++ {
			answer[i] = append(answer[i], level[j].Val)
			if level[j].Left != nil {
				tmp = append(tmp, level[j].Left)
			}
			if level[j].Right != nil {
				tmp = append(tmp, level[j].Right)
			}
		}
		level = tmp
	}
	return answer
}

// @lc code=end
