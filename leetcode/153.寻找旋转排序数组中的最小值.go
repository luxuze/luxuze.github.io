package leetcode

// 147/147 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 7.41 % of golang submissions (2.5 MB)
/*
 * @lc app=leetcode.cn id=153 lang=golang
 *
 * [153] 寻找旋转排序数组中的最小值
 *
 * https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array/description/
 *
 * algorithms
 * Medium (51.98%)
 * Likes:    306
 * Dislikes: 0
 * Total Accepted:    91.9K
 * Total Submissions: 176.8K
 * Testcase Example:  '[3,4,5,1,2]'
 *
 * 假设按照升序排序的数组在预先未知的某个点上进行了旋转。例如，数组 [0,1,2,4,5,6,7]  可能变为 [4,5,6,7,0,1,2] 。
 *
 * 请找出其中最小的元素。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [3,4,5,1,2]
 * 输出：1
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [4,5,6,7,0,1,2]
 * 输出：0
 *
 *
 * 示例 3：
 *
 *
 * 输入：nums = [1]
 * 输出：1
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * -5000
 * nums 中的所有整数都是 唯一 的
 * nums 原来是一个升序排序的数组，但在预先未知的某个点上进行了旋转
 *
 *
 */

// @lc code=start
func findMin(nums []int) int {
	if !(len(nums) > 0) {
		return 0
	}
	if nums[0] < nums[len(nums)-1] {
		return nums[0]
	}
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			return nums[i+1]
		}
	}
	return nums[len(nums)-1]
}

// @lc code=end
