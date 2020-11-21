package leetcode

// 109/109 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 100 % of golang submissions (2 MB)
/*
 * @lc app=leetcode.cn id=66 lang=golang
 *
 * [66] 加一
 */

// @lc code=start
func plusOne(digits []int) []int {
	digits[len(digits)-1]++
	for i := len(digits) - 1; 0 < i; i-- {
		if digits[i] != 10 {
			break
		}
		digits[i] = 0
		digits[i-1]++
	}
	if digits[0] == 10 {
		return append([]int{1, 0}, digits[1:]...)
	}
	return digits
}

// @lc code=end
