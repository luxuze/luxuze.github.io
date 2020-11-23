package leetcode

// 123/123 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 100 % of golang submissions (2.3 MB)
/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] 最长公共前缀
 *
 * https://leetcode-cn.com/problems/longest-common-prefix/description/
 *
 * algorithms
 * Easy (38.92%)
 * Likes:    1355
 * Dislikes: 0
 * Total Accepted:    400.5K
 * Total Submissions: 1M
 * Testcase Example:  '["flower","flow","flight"]'
 *
 * 编写一个函数来查找字符串数组中的最长公共前缀。
 *
 * 如果不存在公共前缀，返回空字符串 ""。
 *
 * 示例 1:
 *
 * 输入: ["flower","flow","flight"]
 * 输出: "fl"
 *
 *
 * 示例 2:
 *
 * 输入: ["dog","racecar","car"]
 * 输出: ""
 * 解释: 输入不存在公共前缀。
 *
 *
 * 说明:
 *
 * 所有输入只包含小写字母 a-z 。
 *
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {
	switch len(strs) {
	case 0:
		return ""
	case 1:
		return strs[0]
	}
	var maxLength int
	for i := range strs {
		l := len(strs[i])
		if i == 0 {
			maxLength = l
		} else if l < maxLength {
			maxLength = l
		}
	}
	result := make([]byte, maxLength)
	for i := 0; i < maxLength; i++ {
		current := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if strs[j][i] != current {
				return string(result[:i])
			}
		}
		result[i] = strs[0][i]
	}
	return string(result)
}

// @lc code=end
