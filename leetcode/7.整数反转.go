package leetcode

// 1032/1032 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 42.86 % of golang submissions (2.2 MB)
/*
 * @lc app=leetcode.cn id=7 lang=golang
 *
 * [7] 整数反转
 */

// @lc code=start
func reverse(x int) int {
	//math.MinInt32 = -2147483648
	//math.MaxInt32 = 2147483647
	var result int
	for x != 0 {
		result = result*10 + x%10
		if result > 2147483647 || result < -2147483648 {
			return 0
		}
		x /= 10
	}
	return result
}

// @lc code=end
