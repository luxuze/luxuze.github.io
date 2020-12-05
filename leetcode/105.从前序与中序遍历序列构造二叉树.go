package leetcode

// 203/203 cases passed (8 ms)
// Your runtime beats 60.19 % of golang submissions
// Your memory usage beats 58.73 % of golang submissions (4.2 MB)
/*
 * @lc app=leetcode.cn id=105 lang=golang
 *
 * [105] 从前序与中序遍历序列构造二叉树
 *
 * https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/description/
 *
 * algorithms
 * Medium (68.57%)
 * Likes:    782
 * Dislikes: 0
 * Total Accepted:    133.3K
 * Total Submissions: 194.4K
 * Testcase Example:  '[3,9,20,15,7]\n[9,3,15,20,7]'
 *
 * 根据一棵树的前序遍历与中序遍历构造二叉树。
 *
 * 注意:
 * 你可以假设树中没有重复的元素。
 *
 * 例如，给出
 *
 * 前序遍历 preorder = [3,9,20,15,7]
 * 中序遍历 inorder = [9,3,15,20,7]
 *
 * 返回如下的二叉树：
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
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
func lc105BuildTree(preorder []int, inorder []int) *TreeNode {
	// preorder = [root, [left...], [right...]]
	// inorder = [[left...], root, [right...]]
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{preorder[0], nil, nil}
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}
	root.Left = lc105BuildTree(preorder[1:len(inorder[:i])+1], inorder[:i])
	root.Right = lc105BuildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])
	return root
}

// @lc code=end
