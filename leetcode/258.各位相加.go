package leetcode

// 1101/1101 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 100 % of golang submissions (2.1 MB)
/*
 * @lc app=leetcode.cn id=258 lang=golang
 *
 * [258] 各位相加
 *
 * https://leetcode-cn.com/problems/add-digits/description/
 *
 * algorithms
 * Easy (69.05%)
 * Likes:    351
 * Dislikes: 0
 * Total Accepted:    73.8K
 * Total Submissions: 106.9K
 * Testcase Example:  '38'
 *
 * 给定一个非负整数 num，反复将各个位上的数字相加，直到结果为一位数。
 *
 * 示例:
 *
 * 输入: 38
 * 输出: 2
 * 解释: 各位相加的过程为：3 + 8 = 11, 1 + 1 = 2。 由于 2 是一位数，所以返回 2。
 *
 *
 * 进阶:
 * 你可以不使用循环或者递归，且在 O(1) 时间复杂度内解决这个问题吗？
 *
 */

// @lc code=start
func addDigits(num int) int {
	return (num-1)%9 + 1
}

// @lc code=end
