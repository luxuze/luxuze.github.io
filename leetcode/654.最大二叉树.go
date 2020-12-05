package leetcode

// 107/107 cases passed (20 ms)
// Your runtime beats 68.82 % of golang submissions
// Your memory usage beats 89.58 % of golang submissions (6.7 MB)
/*
 * @lc app=leetcode.cn id=654 lang=golang
 *
 * [654] 最大二叉树
 *
 * https://leetcode-cn.com/problems/maximum-binary-tree/description/
 *
 * algorithms
 * Medium (81.39%)
 * Likes:    218
 * Dislikes: 0
 * Total Accepted:    26.4K
 * Total Submissions: 32.5K
 * Testcase Example:  '[3,2,1,6,0,5]'
 *
 * 给定一个不含重复元素的整数数组。一个以此数组构建的最大二叉树定义如下：
 *
 *
 * 二叉树的根是数组中的最大元素。
 * 左子树是通过数组中最大值左边部分构造出的最大二叉树。
 * 右子树是通过数组中最大值右边部分构造出的最大二叉树。
 *
 *
 * 通过给定的数组构建最大二叉树，并且输出这个树的根节点。
 *
 *
 *
 * 示例 ：
 *
 * 输入：[3,2,1,6,0,5]
 * 输出：返回下面这棵树的根节点：
 *
 * ⁠     6
 * ⁠   /   \
 * ⁠  3     5
 * ⁠   \    /
 * ⁠    2  0
 * ⁠      \
 * ⁠       1
 *
 *
 *
 *
 * 提示：
 *
 *
 * 给定的数组的大小在 [1, 1000] 之间。
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
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	maxIndex := -1
	maxVal := -1
	for i := 0; i < len(nums); i++ {
		if nums[i] > maxVal {
			maxVal = nums[i]
			maxIndex = i
		}
	}
	root := &TreeNode{maxVal, nil, nil}
	root.Left = constructMaximumBinaryTree(nums[0:maxIndex])
	root.Right = constructMaximumBinaryTree(nums[maxIndex+1:])
	return root
}

// @lc code=end
