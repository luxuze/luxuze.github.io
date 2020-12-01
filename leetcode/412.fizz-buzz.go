package leetcode

// 8/8 cases passed (4 ms)
// Your runtime beats 92.67 % of golang submissions
// Your memory usage beats 79.6 % of golang submissions (3.4 MB)
import "strconv"

/*
 * @lc app=leetcode.cn id=412 lang=golang
 *
 * [412] Fizz Buzz
 *
 * https://leetcode-cn.com/problems/fizz-buzz/description/
 *
 * algorithms
 * Easy (65.23%)
 * Likes:    77
 * Dislikes: 0
 * Total Accepted:    50.6K
 * Total Submissions: 77.5K
 * Testcase Example:  '1'
 *
 * 写一个程序，输出从 1 到 n 数字的字符串表示。
 *
 * 1. 如果 n 是3的倍数，输出“Fizz”；
 *
 * 2. 如果 n 是5的倍数，输出“Buzz”；
 *
 * 3.如果 n 同时是3和5的倍数，输出 “FizzBuzz”。
 *
 * 示例：
 *
 * n = 15,
 *
 * 返回:
 * [
 * ⁠   "1",
 * ⁠   "2",
 * ⁠   "Fizz",
 * ⁠   "4",
 * ⁠   "Buzz",
 * ⁠   "Fizz",
 * ⁠   "7",
 * ⁠   "8",
 * ⁠   "Fizz",
 * ⁠   "Buzz",
 * ⁠   "11",
 * ⁠   "Fizz",
 * ⁠   "13",
 * ⁠   "14",
 * ⁠   "FizzBuzz"
 * ]
 *
 *
 */

// @lc code=start
func fizzBuzz(n int) []string {
	answer := make([]string, n)
	for i := range answer {
		var tmp string
		if (i+1)%3 == 0 {
			tmp += "Fizz"
		}
		if (i+1)%5 == 0 {
			tmp += "Buzz"
		}
		if tmp != "" {
			answer[i] = tmp
		} else {
			answer[i] = strconv.Itoa(i + 1)
		}
	}
	return answer
}

// @lc code=end
