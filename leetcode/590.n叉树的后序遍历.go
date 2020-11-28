package leetcode

// 37/37 cases passed (8 ms)
// Your runtime beats 16.34 % of golang submissions
// Your memory usage beats 14.49 % of golang submissions (6.9 MB)
/*
 * @lc app=leetcode.cn id=590 lang=golang
 *
 * [590] N叉树的后序遍历
 *
 * https://leetcode-cn.com/problems/n-ary-tree-postorder-traversal/description/
 *
 * algorithms
 * Easy (75.09%)
 * Likes:    113
 * Dislikes: 0
 * Total Accepted:    40.8K
 * Total Submissions: 54.4K
 * Testcase Example:  '[1,null,3,2,4,null,5,6]'
 *
 * 给定一个 N 叉树，返回其节点值的后序遍历。
 *
 * 例如，给定一个 3叉树 :
 *
 *
 *
 *
 *
 *
 *
 * 返回其后序遍历: [5,6,3,2,4,1].
 *
 *
 *
 * 说明: 递归法很简单，你可以使用迭代法完成此题吗?
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func postorder(root *Node) (answer []int) {
	if root == nil {
		return
	}
	answer = append([]int{root.Val}, answer...)
	if root.Children == nil {
		return
	}
	for i := range root.Children {
		answer = append(postorder(root.Children[len(root.Children)-i-1]), answer...)
	}
	return
}

// @lc code=end
