package leetcode

// 59/59 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 8.77 % of golang submissions (2.3 MB)
/*
 * @lc app=leetcode.cn id=88 lang=golang
 *
 * [88] 合并两个有序数组
 *
 * https://leetcode-cn.com/problems/merge-sorted-array/description/
 *
 * algorithms
 * Easy (48.96%)
 * Likes:    703
 * Dislikes: 0
 * Total Accepted:    237.9K
 * Total Submissions: 485.8K
 * Testcase Example:  '[1,2,3,0,0,0]\n3\n[2,5,6]\n3'
 *
 * 给你两个有序整数数组 nums1 和 nums2，请你将 nums2 合并到 nums1 中，使 nums1 成为一个有序数组。
 *
 *
 *
 * 说明：
 *
 *
 * 初始化 nums1 和 nums2 的元素数量分别为 m 和 n 。
 * 你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。
 *
 *
 *
 *
 * 示例：
 *
 *
 * 输入：
 * nums1 = [1,2,3,0,0,0], m = 3
 * nums2 = [2,5,6],       n = 3
 *
 * 输出：[1,2,2,3,5,6]
 *
 *
 *
 * 提示：
 *
 *
 * -10^9
 * nums1.length == m + n
 * nums2.length == n
 *
 *
 */

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int) {
	if m == 0 {
		copy(nums1, nums2)
		return
	}
	if n == 0 {
		return
	}
	tmp := make([]int, m)
	copy(tmp, nums1[:m])
	var index int
	for i, j := 0, 0; i < m || j < n; index++ {
		if i == m {
			nums1[index] = nums2[j]
			j++
			continue
		}
		if j == n {
			nums1[index] = tmp[i]
			i++
			continue
		}
		if tmp[i] < nums2[j] {
			nums1[index] = tmp[i]
			i++
		} else {
			nums1[index] = nums2[j]
			j++
		}
	}
}

// @lc code=end
