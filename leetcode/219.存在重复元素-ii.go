package leetcode

// 50/50 cases passed (104 ms)
// Your runtime beats 53.4 % of golang submissions
// Your memory usage beats 31.15 % of golang submissions (11 MB)
import "math"

/*
 * @lc app=leetcode.cn id=219 lang=golang
 *
 * [219] 存在重复元素 II
 *
 * https://leetcode-cn.com/problems/contains-duplicate-ii/description/
 *
 * algorithms
 * Easy (42.44%)
 * Likes:    291
 * Dislikes: 0
 * Total Accepted:    102.5K
 * Total Submissions: 241K
 * Testcase Example:  '[1,2,3,1]\n3'
 *
 * 给定一个整数数组和一个整数 k，判断数组中是否存在两个不同的索引 i 和 j，使得 nums [i] = nums [j]，并且 i 和 j 的差的
 * 绝对值 至多为 k。
 *
 *
 *
 * 示例 1:
 *
 * 输入: nums = [1,2,3,1], k = 3
 * 输出: true
 *
 * 示例 2:
 *
 * 输入: nums = [1,0,1,1], k = 1
 * 输出: true
 *
 * 示例 3:
 *
 * 输入: nums = [1,2,3,1,2,3], k = 2
 * 输出: false
 *
 */

// @lc code=start
func containsNearbyDuplicate(nums []int, k int) bool {
	mp := make(map[int]int)
	for i := range nums {
		if idx, ok := mp[nums[i]]; ok {
			if math.Abs(float64(idx-i)) <= float64(k) {
				return true
			}
		}
		mp[nums[i]] = i
	}
	return false
}

// @lc code=end
