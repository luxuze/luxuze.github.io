package leetcode

// 1017/1017 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 99.94 % of golang submissions (2.1 MB)
import "math"

/*
 * @lc app=leetcode.cn id=69 lang=golang
 *
 * [69] x 的平方根
 *
 * https://leetcode-cn.com/problems/sqrtx/description/
 *
 * algorithms
 * Easy (39.29%)
 * Likes:    730
 * Dislikes: 0
 * Total Accepted:    347.2K
 * Total Submissions: 883.8K
 * Testcase Example:  '4'
 *
 * 实现 int sqrt(int x) 函数。
 *
 * 计算并返回 x 的平方根，其中 x 是非负整数。
 *
 * 由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。
 *
 * 示例 1:
 *
 * 输入: 4
 * 输出: 2
 *
 *
 * 示例 2:
 *
 * 输入: 8
 * 输出: 2
 * 说明: 8 的平方根是 2.82842...,
 * 由于返回类型是整数，小数部分将被舍去。
 *
 *
 */

// @lc code=start
func mySqrt(x int) int {
	return int(math.Sqrt(float64(x)))
}

// @lc code=end
