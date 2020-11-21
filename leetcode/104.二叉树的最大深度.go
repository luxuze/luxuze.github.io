package leetcode

// 39/39 cases passed (4 ms)
// Your runtime beats 92.34 % of golang submissions
// Your memory usage beats 16.67 % of golang submissions (4.5 MB)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
 * @lc app=leetcode.cn id=104 lang=golang
 *
 * [104] 二叉树的最大深度
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
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	que := []*TreeNode{root}
	depth := 1
	start, end := 0, 1

	for {
		numOfChildren := 0
		for i := start; i < end; i++ {
			node := que[i]
			if node.Left != nil {
				que = append(que, node.Left)
				numOfChildren++
			}
			if node.Right != nil {
				que = append(que, node.Right)
				numOfChildren++
			}
		}
		if numOfChildren == 0 {
			break
		}

		depth++
		start = end
		end = end + numOfChildren
	}

	return depth
}

// @lc code=end
