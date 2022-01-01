package leetcode

// 55/55 cases passed (4 ms)
// Your runtime beats 76.74 % of golang submissions
// Your memory usage beats 51.94 % of golang submissions (3.1 MB)
/*
 * @lc app=leetcode.cn id=349 lang=golang
 *
 * [349] 两个数组的交集
 *
 * https://leetcode-cn.com/problems/intersection-of-two-arrays/description/
 *
 * algorithms
 * Easy (73.88%)
 * Likes:    459
 * Dislikes: 0
 * Total Accepted:    238.9K
 * Total Submissions: 323.4K
 * Testcase Example:  '[1,2,2,1]\n[2,2]'
 *
 * 给定两个数组，编写一个函数来计算它们的交集。
 *
 *
 *
 * 示例 1：
 *
 * 输入：nums1 = [1,2,2,1], nums2 = [2,2]
 * 输出：[2]
 *
 *
 * 示例 2：
 *
 * 输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
 * 输出：[9,4]
 *
 *
 *
 * 说明：
 *
 *
 * 输出结果中的每个元素一定是唯一的。
 * 我们可以不考虑输出结果的顺序。
 *
 *
 */

// @lc code=start
func intersection(nums1 []int, nums2 []int) (ans []int) {
	mp := make(map[int]struct{})
	mp2 := make(map[int]struct{})
	for _, v := range nums1 {
		if _, ok := mp[v]; !ok {
			mp[v] = struct{}{}
		}
	}
	for _, v := range nums2 {
		if _, ok := mp[v]; ok {
			if _, ok2 := mp2[v]; ok2 {
				continue
			}
			ans = append(ans, v)
			mp2[v] = struct{}{}
		}
	}
	return
}

// @lc code=end
