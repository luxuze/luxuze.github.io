package leetcode

// 62/62 cases passed (4 ms)
// Your runtime beats 89.71 % of golang submissions
// Your memory usage beats 99.93 % of golang submissions (2.9 MB)
/*
 * @lc app=leetcode.cn id=35 lang=golang
 *
 * [35] 搜索插入位置
 *
 * https://leetcode-cn.com/problems/search-insert-position/description/
 *
 * algorithms
 * Easy (46.89%)
 * Likes:    768
 * Dislikes: 0
 * Total Accepted:    291.4K
 * Total Submissions: 621.5K
 * Testcase Example:  '[1,3,5,6]\n5'
 *
 * 给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
 *
 * 你可以假设数组中无重复元素。
 *
 * 示例 1:
 *
 * 输入: [1,3,5,6], 5
 * 输出: 2
 *
 *
 * 示例 2:
 *
 * 输入: [1,3,5,6], 2
 * 输出: 1
 *
 *
 * 示例 3:
 *
 * 输入: [1,3,5,6], 7
 * 输出: 4
 *
 *
 * 示例 4:
 *
 * 输入: [1,3,5,6], 0
 * 输出: 0
 *
 *
 */

// @lc code=start
func searchInsert(nums []int, target int) int {
	for i := range nums {
		if nums[i] >= target {
			return i
		}
	}
	return len(nums)
}

// @lc code=end
