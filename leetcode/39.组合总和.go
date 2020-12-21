package leetcode

// 170/170 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 18.36 % of golang submissions (4.1 MB)
/*
 * @lc app=leetcode.cn id=39 lang=golang
 *
 * [39] 组合总和
 *
 * https://leetcode-cn.com/problems/combination-sum/description/
 *
 * algorithms
 * Medium (71.73%)
 * Likes:    1096
 * Dislikes: 0
 * Total Accepted:    190.3K
 * Total Submissions: 265.2K
 * Testcase Example:  '[2,3,6,7]\n7'
 *
 * 给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
 *
 * candidates 中的数字可以无限制重复被选取。
 *
 * 说明：
 *
 *
 * 所有数字（包括 target）都是正整数。
 * 解集不能包含重复的组合。
 *
 *
 * 示例 1：
 *
 * 输入：candidates = [2,3,6,7], target = 7,
 * 所求解集为：
 * [
 * ⁠ [7],
 * ⁠ [2,2,3]
 * ]
 *
 *
 * 示例 2：
 *
 * 输入：candidates = [2,3,5], target = 8,
 * 所求解集为：
 * [
 * [2,2,2,2],
 * [2,3,3],
 * [3,5]
 * ]
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= candidates.length <= 30
 * 1 <= candidates[i] <= 200
 * candidate 中的每个元素都是独一无二的。
 * 1 <= target <= 500
 *
 *
 */

// @lc code=start
func combinationSum(candidates []int, target int) (answer [][]int) {
	var answerItem []int
	var dfs func(t, i int)
	dfs = func(t, i int) {
		if i == len(candidates) {
			return
		}
		if t == 0 {
			item := make([]int, len(answerItem))
			copy(item, answerItem)
			answer = append([][]int{item}, answer...)
			return
		}
		dfs(t, i+1)
		if t >= candidates[i] {
			answerItem = append(answerItem, candidates[i])
			dfs(t-candidates[i], i)
			answerItem = answerItem[:len(answerItem)-1]
		}
	}
	dfs(target, 0)
	return
}

// @lc code=end
