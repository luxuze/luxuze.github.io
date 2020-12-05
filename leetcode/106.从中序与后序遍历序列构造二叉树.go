package leetcode

// 203/203 cases passed (4 ms)
// Your runtime beats 94.44 % of golang submissions
// Your memory usage beats 86.29 % of golang submissions (4.1 MB)
/*
 * @lc app=leetcode.cn id=106 lang=golang
 *
 * [106] 从中序与后序遍历序列构造二叉树
 *
 * https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/description/
 *
 * algorithms
 * Medium (70.78%)
 * Likes:    411
 * Dislikes: 0
 * Total Accepted:    79.5K
 * Total Submissions: 112.3K
 * Testcase Example:  '[9,3,15,20,7]\n[9,15,7,20,3]'
 *
 * 根据一棵树的中序遍历与后序遍历构造二叉树。
 *
 * 注意:
 * 你可以假设树中没有重复的元素。
 *
 * 例如，给出
 *
 * 中序遍历 inorder = [9,3,15,20,7]
 * 后序遍历 postorder = [9,15,7,20,3]
 *
 * 返回如下的二叉树：
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
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
func buildTree(inorder []int, postorder []int) *TreeNode {
	// inorder [[left...], root, [right...]]
	// postorder [[left...], [right...], root]
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}
	root := &TreeNode{postorder[len(postorder)-1], nil, nil}
	var rootIndex int
	for ; rootIndex < len(inorder); rootIndex++ {
		if inorder[rootIndex] == root.Val {
			break
		}
	}
	root.Left = buildTree(inorder[:rootIndex], postorder[:rootIndex])
	root.Right = buildTree(inorder[rootIndex+1:], postorder[rootIndex:len(postorder)-1])
	return root
}

// @lc code=end
