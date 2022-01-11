package leetcode

// 23/23 cases passed (8 ms)
// Your runtime beats 91.85 % of golang submissions
// Your memory usage beats 49.53 % of golang submissions (6.2 MB)
/*
 * @lc app=leetcode.cn id=501 lang=golang
 *
 * [501] 二叉搜索树中的众数
 *
 * https://leetcode-cn.com/problems/find-mode-in-binary-search-tree/description/
 *
 * algorithms
 * Easy (51.85%)
 * Likes:    396
 * Dislikes: 0
 * Total Accepted:    75.1K
 * Total Submissions: 144.9K
 * Testcase Example:  '[1,null,2,2]'
 *
 * 给定一个有相同值的二叉搜索树（BST），找出 BST 中的所有众数（出现频率最高的元素）。
 *
 * 假定 BST 有如下定义：
 *
 *
 * 结点左子树中所含结点的值小于等于当前结点的值
 * 结点右子树中所含结点的值大于等于当前结点的值
 * 左子树和右子树都是二叉搜索树
 *
 *
 * 例如：
 * 给定 BST [1,null,2,2],
 *
 * ⁠  1
 * ⁠   \
 * ⁠    2
 * ⁠   /
 * ⁠  2
 *
 *
 * 返回[2].
 *
 * 提示：如果众数超过1个，不需考虑输出顺序
 *
 * 进阶：你可以不使用额外的空间吗？（假设由递归产生的隐式调用栈的开销不被计算在内）
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
func findMode(root *TreeNode) (ans []int) {
	var (
		dfs           func(tn *TreeNode)
		val, cnt, max int
	)
	dfs = func(tn *TreeNode) {
		if tn == nil {
			return
		}
		dfs(tn.Left)
		if val == tn.Val {
			cnt++
		} else {
			val, cnt = tn.Val, 1
		}
		if cnt > max {
			max = cnt
			ans = []int{val}
		} else if cnt == max {
			ans = append(ans, val)
		}
		dfs(tn.Right)
	}
	dfs(root)
	return
}

// @lc code=end
