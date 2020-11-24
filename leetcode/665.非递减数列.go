package leetcode

// 334/334 cases passed (28 ms)
// Your runtime beats 87.94 % of golang submissions
// Your memory usage beats 78.57 % of golang submissions (6.4 MB)
/*
 * @lc app=leetcode.cn id=665 lang=golang
 *
 * [665] 非递减数列
 *
 * https://leetcode-cn.com/problems/non-decreasing-array/description/
 *
 * algorithms
 * Easy (23.10%)
 * Likes:    375
 * Dislikes: 0
 * Total Accepted:    30.7K
 * Total Submissions: 132.8K
 * Testcase Example:  '[4,2,3]'
 *
 * 给你一个长度为 n 的整数数组，请你判断在 最多 改变 1 个元素的情况下，该数组能否变成一个非递减数列。
 *
 * 我们是这样定义一个非递减数列的： 对于数组中所有的 i (0 <= i <= n-2)，总满足 nums[i] <= nums[i + 1]。
 *
 *
 *
 * 示例 1:
 *
 * 输入: nums = [4,2,3]
 * 输出: true
 * 解释: 你可以通过把第一个4变成1来使得它成为一个非递减数列。
 *
 *
 * 示例 2:
 *
 * 输入: nums = [4,2,1]
 * 输出: false
 * 解释: 你不能在只改变一个元素的情况下将其变为非递减数列。
 *
 *
 *
 *
 * 说明：
 *
 *
 * 1 <= n <= 10 ^ 4
 * - 10 ^ 5 <= nums[i] <= 10 ^ 5
 *
 *
 */

// @lc code=start
func checkPossibility(nums []int) bool {
	if len(nums) < 3 {
		return true
	}
	var flag bool
	if nums[0] > nums[1] {
		flag = true
		nums[0] = nums[1]
	}
	for i := 1; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			if flag {
				return false
			}
			flag = true

			if nums[i-1] <= nums[i+1] {
				nums[i] = nums[i-1]
			} else {
				nums[i+1] = nums[i]
			}
		}
	}
	return true
}

// @lc code=end
