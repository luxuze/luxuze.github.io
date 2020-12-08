package leetcode

// 32/32 cases passed (12 ms)
// Your runtime beats 42.88 % of golang submissions
// Your memory usage beats 64.96 % of golang submissions (3.4 MB)
import "sort"

/*
 * @lc app=leetcode.cn id=215 lang=golang
 *
 * [215] 数组中的第K个最大元素
 *
 * https://leetcode-cn.com/problems/kth-largest-element-in-an-array/description/
 *
 * algorithms
 * Medium (64.51%)
 * Likes:    814
 * Dislikes: 0
 * Total Accepted:    231.1K
 * Total Submissions: 358.2K
 * Testcase Example:  '[3,2,1,5,6,4]\n2'
 *
 * 在未排序的数组中找到第 k 个最大的元素。请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
 *
 * 示例 1:
 *
 * 输入: [3,2,1,5,6,4] 和 k = 2
 * 输出: 5
 *
 *
 * 示例 2:
 *
 * 输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
 * 输出: 4
 *
 * 说明:
 *
 * 你可以假设 k 总是有效的，且 1 ≤ k ≤ 数组的长度。
 *
 */

// @lc code=start
func findKthLargest(nums []int, k int) int {
	sort.SliceStable(nums, func(a, b int) bool { return nums[a] > nums[b] })
	return nums[k-1]
}

// @lc code=end
