package leetcode

// 95/95 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 27.46 % of golang submissions (2.1 MB)
/*
 * @lc app=leetcode.cn id=409 lang=golang
 *
 * [409] 最长回文串
 *
 * https://leetcode-cn.com/problems/longest-palindrome/description/
 *
 * algorithms
 * Easy (55.58%)
 * Likes:    370
 * Dislikes: 0
 * Total Accepted:    104.9K
 * Total Submissions: 188.8K
 * Testcase Example:  '"abccccdd"'
 *
 * 给定一个包含大写字母和小写字母的字符串，找到通过这些字母构造成的最长的回文串。
 *
 * 在构造过程中，请注意区分大小写。比如 "Aa" 不能当做一个回文字符串。
 *
 * 注意:
 * 假设字符串的长度不会超过 1010。
 *
 * 示例 1:
 *
 *
 * 输入:
 * "abccccdd"
 *
 * 输出:
 * 7
 *
 * 解释:
 * 我们可以构造的最长的回文串是"dccaccd", 它的长度是 7。
 *
 *
 */

// @lc code=start
func longestPalindrome(s string) (cnt int) {
	mp := make(map[rune]int)
	for _, v := range s {
		mp[v]++
	}
	for _, v := range mp {
		cnt += v / 2 * 2
	}
	if cnt < len(s) {
		cnt++
	}
	return
}

// @lc code=end
