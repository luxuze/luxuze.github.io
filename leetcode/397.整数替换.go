package leetcode

// 47/47 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 100 % of golang submissions (1.9 MB)
import "math"

/*
 * @lc app=leetcode.cn id=397 lang=golang
 *
 * [397] 整数替换
 */

// @lc code=start
func integerReplacement(n int) int {
	if n < 3 {
		return n - 1
	}
	if n == math.MinInt32 {
		return 32
	}
	if n%2 == 0 {
		return integerReplacement(n/2) + 1
	}
	i1 := integerReplacement(n+1) + 1
	i2 := integerReplacement(n-1) + 1
	if i1 > i2 {
		return i2
	}
	return i1
}

// @lc code=end
