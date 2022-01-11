package leetcode

// 27/27 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 12.98 % of golang submissions (1.9 MB)
import "strings"

/*
 * @lc app=leetcode.cn id=434 lang=golang
 *
 * [434] 字符串中的单词数
 *
 * https://leetcode-cn.com/problems/number-of-segments-in-a-string/description/
 *
 * algorithms
 * Easy (39.81%)
 * Likes:    164
 * Dislikes: 0
 * Total Accepted:    67.2K
 * Total Submissions: 168.8K
 * Testcase Example:  '"Hello, my name is John"'
 *
 * 统计字符串中的单词个数，这里的单词指的是连续的不是空格的字符。
 *
 * 请注意，你可以假定字符串里不包括任何不可打印的字符。
 *
 * 示例:
 *
 * 输入: "Hello, my name is John"
 * 输出: 5
 * 解释: 这里的单词是指连续的不是空格的字符，所以 "Hello," 算作 1 个单词。
 *
 *
 */

// @lc code=start
func countSegments(s string) int {

	var ans int
	l := strings.Split(s, " ")
	for _, v := range l {
		if v != "" {
			ans++
		}
	}
	return ans
}

// @lc code=end
