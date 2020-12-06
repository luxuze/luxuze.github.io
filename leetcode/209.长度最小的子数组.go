package leetcode

// 15/15 cases passed (76 ms)
// Your runtime beats 14.47 % of golang submissions
// Your memory usage beats 27.39 % of golang submissions (3.8 MB)
/*
 * @lc app=leetcode.cn id=209 lang=golang
 *
 * [209] 长度最小的子数组
 *
 * https://leetcode-cn.com/problems/minimum-size-subarray-sum/description/
 *
 * algorithms
 * Medium (44.60%)
 * Likes:    507
 * Dislikes: 0
 * Total Accepted:    101K
 * Total Submissions: 226.4K
 * Testcase Example:  '7\n[2,3,1,2,4,3]'
 *
 * 给定一个含有 n 个正整数的数组和一个正整数 s ，找出该数组中满足其和 ≥ s 的长度最小的 连续
 * 子数组，并返回其长度。如果不存在符合条件的子数组，返回 0。
 *
 *
 *
 * 示例：
 *
 * 输入：s = 7, nums = [2,3,1,2,4,3]
 * 输出：2
 * 解释：子数组 [4,3] 是该条件下的长度最小的子数组。
 *
 *
 *
 *
 * 进阶：
 *
 *
 * 如果你已经完成了 O(n) 时间复杂度的解法, 请尝试 O(n log n) 时间复杂度的解法。
 *
 *
 */

// @lc code=start
func minSubArrayLen(s int, nums []int) int {
	answer := len(nums)
	if !(answer > 0) {
		return 0
	}
	var flag bool
	for i := range nums {
		var tmp int
		for j := i; j < len(nums); j++ {
			tmp += nums[j]
			if tmp >= s {
				flag = true
				if j-i < answer {
					answer = j - i + 1
				}
				break
			}
		}
	}
	if !flag {
		return 0
	}
	return answer
}

// @lc code=end
