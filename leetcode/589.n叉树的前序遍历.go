package leetcode

// 37/37 cases passed (4 ms)
// Your runtime beats 67.24 % of golang submissions
// Your memory usage beats 21.55 % of golang submissions (6.1 MB)
/*
 * @lc app=leetcode.cn id=589 lang=golang
 *
 * [589] N叉树的前序遍历
 *
 * https://leetcode-cn.com/problems/n-ary-tree-preorder-traversal/description/
 *
 * algorithms
 * Easy (74.04%)
 * Likes:    116
 * Dislikes: 0
 * Total Accepted:    55.7K
 * Total Submissions: 75.2K
 * Testcase Example:  '[1,null,3,2,4,null,5,6]'
 *
 * 给定一个 N 叉树，返回其节点值的前序遍历。
 *
 * 例如，给定一个 3叉树 :
 *
 *
 *
 *
 *
 *
 *
 * 返回其前序遍历: [1,3,5,6,2,4]。
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

func preorder(root *Node) (answer []int) {
	if root == nil {
		return
	}
	answer = append(answer, root.Val)
	if root.Children == nil {
		return
	}
	for i := range root.Children {
		answer = append(answer, preorder(root.Children[i])...)
	}
	return
}

// @lc code=end
