package leetcode

// 45/45 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 8.05 % of golang submissions (2.6 MB)

// 45/45 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 28.74 % of golang submissions (2.5 MB)

/*
 * @lc app=leetcode.cn id=783 lang=golang
 *
 * [783] 二叉搜索树节点最小距离
 *
 * https://leetcode-cn.com/problems/minimum-distance-between-bst-nodes/description/
 *
 * algorithms
 * Easy (55.81%)
 * Likes:    102
 * Dislikes: 0
 * Total Accepted:    22.3K
 * Total Submissions: 40K
 * Testcase Example:  '[4,2,6,1,3,null,null]'
 *
 * 给定一个二叉搜索树的根节点 root，返回树中任意两节点的差的最小值。
 *
 *
 *
 * 示例：
 *
 * 输入: root = [4,2,6,1,3,null,null]
 * 输出: 1
 * 解释:
 * 注意，root是树节点对象(TreeNode object)，而不是数组。
 *
 * 给定的树 [4,2,6,1,3,null,null] 可表示为下图:
 *
 * ⁠         4
 * ⁠       /   \
 * ⁠     2      6
 * ⁠    / \
 * ⁠   1   3
 *
 * 最小的差值是 1, 它是节点1和节点2的差值, 也是节点3和节点2的差值。
 *
 *
 *
 * 注意：
 *
 *
 * 二叉树的大小范围在 2 到 100。
 * 二叉树总是有效的，每个节点的值都是整数，且不重复。
 * 本题与 530：https://leetcode-cn.com/problems/minimum-absolute-difference-in-bst/
 * 相同
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
// func minDiffInBST(root *TreeNode) int {
// 	var (
// 		list   []int
// 		answer int = 100
// 		min        = func(a, b int) int {
// 			if a < b {
// 				return a
// 			}
// 			return b
// 		}
// 		dfs func(*TreeNode)
// 	)

// 	dfs = func(tn *TreeNode) {
// 		if tn == nil {
// 			return
// 		}
// 		list = append(list, tn.Val)
// 		dfs(tn.Left)
// 		dfs(tn.Right)
// 	}
// 	dfs(root)
// 	sort.SliceStable(list, func(i, j int) bool { return list[i] < list[j] })
// 	fmt.Println(list)
// 	for i := 0; i < len(list)-1; i++ {
// 		answer = min(answer, list[i+1]-list[i])
// 	}
// 	return answer
// }
func minDiffInBST(root *TreeNode) int {
	var (
		answer int = 100
		temp   int
		min    = func(a, b int) int {
			if a < b {
				return a
			}
			return b
		}
		dfs func(*TreeNode)
	)
	dfs = func(tn *TreeNode) {
		if tn == nil {
			return
		}
		dfs(tn.Left)
		if temp != 0 {
			answer = min(answer, tn.Val-temp)
		}
		temp = tn.Val
		dfs(tn.Right)
	}
	dfs(root)

	return answer
}

// @lc code=end
