package leetcode

// 31/31 cases passed (4 ms)
// Your runtime beats 92.35 % of golang submissions
// Your memory usage beats 100 % of golang submissions (2.9 MB)
import "sort"

/*
 * @lc app=leetcode.cn id=414 lang=golang
 *
 * [414] 第三大的数
 *
 * https://leetcode-cn.com/problems/third-maximum-number/description/
 *
 * algorithms
 * Easy (39.41%)
 * Likes:    328
 * Dislikes: 0
 * Total Accepted:    95.2K
 * Total Submissions: 241.6K
 * Testcase Example:  '[3,2,1]'
 *
 * 给你一个非空数组，返回此数组中 第三大的数 。如果不存在，则返回数组中最大的数。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：[3, 2, 1]
 * 输出：1
 * 解释：第三大的数是 1 。
 *
 * 示例 2：
 *
 *
 * 输入：[1, 2]
 * 输出：2
 * 解释：第三大的数不存在, 所以返回最大的数 2 。
 *
 *
 * 示例 3：
 *
 *
 * 输入：[2, 2, 3, 1]
 * 输出：1
 * 解释：注意，要求返回第三大的数，是指在所有不同数字中排第三大的数。
 * 此例中存在两个值为 2 的数，它们都排第二。在所有不同数字中排第三大的数为 1 。
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * -2^31
 *
 *
 *
 *
 * 进阶：你能设计一个时间复杂度 O(n) 的解决方案吗？
 *
 */

// @lc code=start
func thirdMax(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)
	var idx int
	cursor := len(nums) - 1
	v := nums[cursor]
	for ; cursor >= 0; cursor-- {
		if v != nums[cursor] {
			v = nums[cursor]
			idx++
		}
		if idx == 2 {
			break
		}
	}
	if idx == 2 {
		return v
	}
	return nums[len(nums)-1]
}

// @lc code=end
