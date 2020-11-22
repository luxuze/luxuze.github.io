package leetcode

// 52/52 cases passed (4 ms)
// Your runtime beats 96.8 % of golang submissions
// Your memory usage beats 72.3 % of golang submissions (3 MB)
/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 *
 * https://leetcode-cn.com/problems/two-sum/description/
 *
 * algorithms
 * Easy (49.32%)
 * Likes:    9670
 * Dislikes: 0
 * Total Accepted:    1.6M
 * Total Submissions: 3.2M
 * Testcase Example:  '[2,7,11,15]\n9'
 *
 * 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
 *
 * 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。
 *
 *
 *
 * 示例:
 *
 * 给定 nums = [2, 7, 11, 15], target = 9
 *
 * 因为 nums[0] + nums[1] = 2 + 7 = 9
 * 所以返回 [0, 1]
 *
 *
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	for i := range nums {
		i2 := findElementInSlice(nums, target-nums[i])
		if i2 != -1 && i2 != i {
			return []int{i, i2}
		}
	}
	return []int{}
}

func findElementInSlice(s []int, ele int) int {
	for i := range s {
		if s[i] == ele {
			return i
		}
	}
	return -1
}

// @lc code=end
