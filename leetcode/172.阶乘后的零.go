package leetcode

// 500/500 cases passed (4 ms)
// Your runtime beats 20.3 % of golang submissions
// Your memory usage beats 54.85 % of golang submissions (2 MB)
/*
 * @lc app=leetcode.cn id=172 lang=golang
 *
 * [172] 阶乘后的零
 *
 * https://leetcode-cn.com/problems/factorial-trailing-zeroes/description/
 *
 * algorithms
 * Easy (43.41%)
 * Likes:    486
 * Dislikes: 0
 * Total Accepted:    82.8K
 * Total Submissions: 190.5K
 * Testcase Example:  '3'
 *
 * 给定一个整数 n，返回 n! 结果尾数中零的数量。
 *
 * 示例 1:
 *
 * 输入: 3
 * 输出: 0
 * 解释: 3! = 6, 尾数中没有零。
 *
 * 示例 2:
 *
 * 输入: 5
 * 输出: 1
 * 解释: 5! = 120, 尾数中有 1 个零.
 *
 * 说明: 你算法的时间复杂度应为 O(log n) 。
 *
 */

// @lc code=start
func trailingZeroes(n int) (ans int) {
	for ; n > 1; n-- {
		_n := n
		for _n%5 == 0 {
			_n /= 5
			ans++
		}
	}
	return
}

// @lc code=end
