package leetcode

// 11509/11509 cases passed (28 ms)
// Your runtime beats 24.53 % of golang submissions
// Your memory usage beats 52 % of golang submissions (5.2 MB)
import "math"

/*
 * @lc app=leetcode.cn id=9 lang=golang
 *
 * [9] 回文数
 */

// @lc code=start
func lc9isPalindrome(x int) bool {
	if math.Signbit(float64(x)) {
		return false
	}
	rev := 0
	remainder := 0
	argz := x
	for argz > 0 {
		rev *= 10
		remainder = argz % 10
		rev += remainder
		argz /= 10
	}
	if rev == x {
		return true
	}
	return false
}

// @lc code=end
