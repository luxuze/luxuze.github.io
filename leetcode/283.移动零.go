package leetcode

// 74/74 cases passed (68 ms)
// Your runtime beats 5.82 % of golang submissions
// Your memory usage beats 96.97 % of golang submissions (6.6 MB)
/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 *
 * https://leetcode-cn.com/problems/move-zeroes/description/
 *
 * algorithms
 * Easy (64.04%)
 * Likes:    1373
 * Dislikes: 0
 * Total Accepted:    584.2K
 * Total Submissions: 912.2K
 * Testcase Example:  '[0,1,0,3,12]'
 *
 * 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
 *
 * 示例:
 *
 * 输入: [0,1,0,3,12]
 * 输出: [1,3,12,0,0]
 *
 * 说明:
 *
 *
 * 必须在原数组上操作，不能拷贝额外的数组。
 * 尽量减少操作次数。
 *
 *
 */

// @lc code=start
func moveZeroes(nums []int) {
	for i := range nums {
		j := i
		for j < len(nums)-1 && nums[j] == 0 {
			j++
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
}

// @lc code=end
