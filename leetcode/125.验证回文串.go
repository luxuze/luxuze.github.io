package leetcode

// 481/481 cases passed (4 ms)
// Your runtime beats 72.63 % of golang submissions
// Your memory usage beats 78.59 % of golang submissions (2.8 MB)
import "strings"

/*
 * @lc app=leetcode.cn id=125 lang=golang
 *
 * [125] 验证回文串
 *
 * https://leetcode-cn.com/problems/valid-palindrome/description/
 *
 * algorithms
 * Easy (46.63%)
 * Likes:    300
 * Dislikes: 0
 * Total Accepted:    180.7K
 * Total Submissions: 387.5K
 * Testcase Example:  '"A man, a plan, a canal: Panama"'
 *
 * 给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
 *
 * 说明：本题中，我们将空字符串定义为有效的回文串。
 *
 * 示例 1:
 *
 * 输入: "A man, a plan, a canal: Panama"
 * 输出: true
 *
 *
 * 示例 2:
 *
 * 输入: "race a car"
 * 输出: false
 *
 *
 */

// @lc code=start
func isPalindrome(s string) bool {
	if s == "" {
		return true
	}
	check := func(b byte) bool {
		if (b >= '0' && b <= '9') || (b >= 'a' && b <= 'z') {
			return true
		}
		return false
	}
	s = strings.ToLower(s)
	for i, j := 0, len(s)-1; i < j; {
		if !check(s[i]) {
			i++
			continue
		}
		if !check(s[j]) {
			j--
			continue
		}
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

// @lc code=end
