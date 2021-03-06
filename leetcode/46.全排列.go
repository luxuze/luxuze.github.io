package leetcode

// 25/25 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 67.09 % of golang submissions (2.7 MB)
/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 *
 * https://leetcode-cn.com/problems/permutations/description/
 *
 * algorithms
 * Medium (77.67%)
 * Likes:    1181
 * Dislikes: 0
 * Total Accepted:    265.4K
 * Total Submissions: 341.7K
 * Testcase Example:  '[1,2,3]'
 *
 * 给定一个 没有重复 数字的序列，返回其所有可能的全排列。
 *
 * 示例:
 *
 * 输入: [1,2,3]
 * 输出:
 * [
 * ⁠ [1,2,3],
 * ⁠ [1,3,2],
 * ⁠ [2,1,3],
 * ⁠ [2,3,1],
 * ⁠ [3,1,2],
 * ⁠ [3,2,1]
 * ]
 *
 */

// @lc code=start
func permute(nums []int) (ans [][]int) {

	visited := make(map[int]bool, len(nums))

	var dfs func(ns []int)
	dfs = func(ns []int) {
		if len(ns) == len(nums) {
			ans = append(ans, ns)
			return
		}
		for i := range nums {
			if visited[nums[i]] {
				continue
			}
			visited[nums[i]] = true
			dfs(append(ns, nums[i]))
			visited[nums[i]] = false
		}
	}

	dfs(nil)
	return
}

// @lc code=end
