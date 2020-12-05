package leetcode

// 34/34 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 5.49 % of golang submissions (6.7 MB)
/*
 * @lc app=leetcode.cn id=107 lang=golang
 *
 * [107] 二叉树的层次遍历 II
 *
 * https://leetcode-cn.com/problems/binary-tree-level-order-traversal-ii/description/
 *
 * algorithms
 * Easy (67.79%)
 * Likes:    373
 * Dislikes: 0
 * Total Accepted:    112.5K
 * Total Submissions: 166K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给定一个二叉树，返回其节点值自底向上的层次遍历。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）
 *
 * 例如：
 * 给定二叉树 [3,9,20,null,null,15,7],
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 *
 *
 * 返回其自底向上的层次遍历为：
 *
 * [
 * ⁠ [15,7],
 * ⁠ [9,20],
 * ⁠ [3]
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
func levelOrderBottom(root *TreeNode) (answer [][]int) {
	if root == nil {
		return
	}
	list := []*TreeNode{root}
	for len(list) > 0 {
		var listItem []int
		var newList []*TreeNode
		for i := range list {
			listItem = append(listItem, list[i].Val)
			if list[i].Left != nil {
				newList = append(newList, list[i].Left)
			}
			if list[i].Right != nil {
				newList = append(newList, list[i].Right)
			}
		}
		answer = append([][]int{listItem}, answer...)
		list = newList
	}
	return
}

// @lc code=end
