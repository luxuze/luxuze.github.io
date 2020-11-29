package leetcode

// 76/76 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 48.44 % of golang submissions (2.6 MB)
/*
 * @lc app=leetcode.cn id=951 lang=golang
 *
 * [951] 翻转等价二叉树
 *
 * https://leetcode-cn.com/problems/flip-equivalent-binary-trees/description/
 *
 * algorithms
 * Medium (65.54%)
 * Likes:    70
 * Dislikes: 0
 * Total Accepted:    7K
 * Total Submissions: 10.8K
 * Testcase Example:  '[1,2,3,4,5,6,null,null,null,7,8]\n[1,3,2,null,6,4,5,null,null,null,null,8,7]'
 *
 * 我们可以为二叉树 T 定义一个翻转操作，如下所示：选择任意节点，然后交换它的左子树和右子树。
 *
 * 只要经过一定次数的翻转操作后，能使 X 等于 Y，我们就称二叉树 X 翻转等价于二叉树 Y。
 *
 * 编写一个判断两个二叉树是否是翻转等价的函数。这些树由根节点 root1 和 root2 给出。
 *
 *
 *
 * 示例：
 *
 * 输入：root1 = [1,2,3,4,5,6,null,null,null,7,8], root2 =
 * [1,3,2,null,6,4,5,null,null,null,null,8,7]
 * 输出：true
 * 解释：我们翻转值为 1，3 以及 5 的三个节点。
 *
 *
 *
 *
 *
 * 提示：
 *
 *
 * 每棵树最多有 100 个节点。
 * 每棵树中的每个值都是唯一的、在 [0, 99] 范围内的整数。
 *
 *
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
func flipEquiv(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil && root2 != nil || root1 != nil && root2 == nil {
		return false
	}
	return root1.Val == root2.Val &&
		(flipEquiv(root1.Left, root2.Left) && flipEquiv(root1.Right, root2.Right) ||
			flipEquiv(root1.Left, root2.Right) && flipEquiv(root1.Right, root2.Left))
}

// @lc code=end
