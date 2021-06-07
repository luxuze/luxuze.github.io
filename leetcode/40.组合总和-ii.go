package leetcode

// 174/174 cases passed (508 ms)
// Your runtime beats 5.07 % of golang submissions
// Your memory usage beats 5.19 % of golang submissions (8.6 MB)
import (
	"sort"
)

/*
 * @lc app=leetcode.cn id=40 lang=golang
 *
 * [40] 组合总和 II
 *
 * https://leetcode-cn.com/problems/combination-sum-ii/description/
 *
 * algorithms
 * Medium (64.04%)
 * Likes:    599
 * Dislikes: 0
 * Total Accepted:    163.6K
 * Total Submissions: 257.2K
 * Testcase Example:  '[10,1,2,7,6,1,5]\n8'
 *
 * 给定一个数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
 *
 * candidates 中的每个数字在每个组合中只能使用一次。
 *
 * 说明：
 *
 *
 * 所有数字（包括目标数）都是正整数。
 * 解集不能包含重复的组合。
 *
 *
 * 示例 1:
 *
 * 输入: candidates = [10,1,2,7,6,1,5], target = 8,
 * 所求解集为:
 * [
 * ⁠ [1, 7],
 * ⁠ [1, 2, 5],
 * ⁠ [2, 6],
 * ⁠ [1, 1, 6]
 * ]
 *
 *
 * 示例 2:
 *
 * 输入: candidates = [2,5,2,1,2], target = 5,
 * 所求解集为:
 * [
 * [1,2,2],
 * [5]
 * ]
 *
 */

// @lc code=start
func combinationSum2(candidates []int, target int) (ans [][]int) {
	if len(candidates) == 0 {
		return
	}
	sort.Ints(candidates)

	// FUCK !!!
	exist := func(a [][]int, b []int) bool {
		for _, ai := range a {
			if len(ai) == len(b) {
				var exist bool = true
				for i := range ai {
					if ai[i] != b[i] {
						exist = false
						break
					}
				}
				if exist {
					return true
				}
			}
		}
		return false
	}

	var dfs func(index, sum int, ans_item []int)
	dfs = func(index, sum int, ans_item []int) {
		if sum == target {
			if !exist(ans, ans_item) {
				ans_item_copy := make([]int, len(ans_item))
				copy(ans_item_copy, ans_item)
				ans = append(ans, ans_item_copy)
			}
			return
		}
		if index == len(candidates) || sum > target {
			return
		}
		// 剪枝减不掉, 操

		// 不要当前值
		dfs(index+1, sum, ans_item)

		// 要当前值
		ans_item = append(ans_item, candidates[index])
		sum += candidates[index]

		dfs(index+1, sum, ans_item)
	}

	dfs(0, 0, nil)
	return
}

// @lc code=end
