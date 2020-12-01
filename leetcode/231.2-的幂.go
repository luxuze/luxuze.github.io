package leetcode

// 1108/1108 cases passed (4 ms)
// Your runtime beats 39.27 % of golang submissions
// Your memory usage beats 48.62 % of golang submissions (2.2 MB)
/*
 * @lc app=leetcode.cn id=231 lang=golang
 *
 * [231] 2的幂
 *
 * https://leetcode-cn.com/problems/power-of-two/description/
 *
 * algorithms
 * Easy (48.82%)
 * Likes:    263
 * Dislikes: 0
 * Total Accepted:    88.4K
 * Total Submissions: 181.1K
 * Testcase Example:  '1'
 *
 * 给定一个整数，编写一个函数来判断它是否是 2 的幂次方。
 *
 * 示例 1:
 *
 * 输入: 1
 * 输出: true
 * 解释: 2^0 = 1
 *
 * 示例 2:
 *
 * 输入: 16
 * 输出: true
 * 解释: 2^4 = 16
 *
 * 示例 3:
 *
 * 输入: 218
 * 输出: false
 *
 */

// @lc code=start
func isPowerOfTwo(n int) bool {
	var v int = 1
	for v <= n {
		if v == n {
			return true
		}
		v = v * 2
	}
	return false
}

// @lc code=end
