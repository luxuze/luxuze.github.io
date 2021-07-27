package leetcode

// 1013/1013 cases passed (4 ms)
// Your runtime beats 34.04 % of golang submissions
// Your memory usage beats 100 % of golang submissions (2.1 MB)
/*
 * @lc app=leetcode.cn id=263 lang=golang
 *
 * [263] 丑数
 *
 * https://leetcode-cn.com/problems/ugly-number/description/
 *
 * algorithms
 * Easy (51.50%)
 * Likes:    266
 * Dislikes: 0
 * Total Accepted:    100.3K
 * Total Submissions: 194.7K
 * Testcase Example:  '6'
 *
 * 给你一个整数 n ，请你判断 n 是否为 丑数 。如果是，返回 true ；否则，返回 false 。
 *
 * 丑数 就是只包含质因数 2、3 和/或 5 的正整数。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：n = 6
 * 输出：true
 * 解释：6 = 2 × 3
 *
 * 示例 2：
 *
 *
 * 输入：n = 8
 * 输出：true
 * 解释：8 = 2 × 2 × 2
 *
 *
 * 示例 3：
 *
 *
 * 输入：n = 14
 * 输出：false
 * 解释：14 不是丑数，因为它包含了另外一个质因数 7 。
 *
 *
 * 示例 4：
 *
 *
 * 输入：n = 1
 * 输出：true
 * 解释：1 通常被视为丑数。
 *
 *
 *
 *
 * 提示：
 *
 *
 * -2^31
 *
 *
 */

// @lc code=start
func isUgly(n int) bool {
	if n <= 0 {
		return false
	}
	for _, v := range []int{2, 3, 5} {
		for n%v == 0 {
			n /= v

		}
	}

	return n == 1
}

// @lc code=end
