package leetcode

import "strings"

// 36/36 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 23.01 % of golang submissions (2 MB)
/*
 * @lc app=leetcode.cn id=290 lang=golang
 *
 * [290] 单词规律
 *
 * https://leetcode-cn.com/problems/word-pattern/description/
 *
 * algorithms
 * Easy (45.53%)
 * Likes:    417
 * Dislikes: 0
 * Total Accepted:    90.2K
 * Total Submissions: 198.2K
 * Testcase Example:  '"abba"\n"dog cat cat dog"'
 *
 * 给定一种规律 pattern 和一个字符串 str ，判断 str 是否遵循相同的规律。
 *
 * 这里的 遵循 指完全匹配，例如， pattern 里的每个字母和字符串 str 中的每个非空单词之间存在着双向连接的对应规律。
 *
 * 示例1:
 *
 * 输入: pattern = "abba", str = "dog cat cat dog"
 * 输出: true
 *
 * 示例 2:
 *
 * 输入:pattern = "abba", str = "dog cat cat fish"
 * 输出: false
 *
 * 示例 3:
 *
 * 输入: pattern = "aaaa", str = "dog cat cat dog"
 * 输出: false
 *
 * 示例 4:
 *
 * 输入: pattern = "abba", str = "dog dog dog dog"
 * 输出: false
 *
 * 说明:
 * 你可以假设 pattern 只包含小写字母， str 包含了由单个空格分隔的小写字母。
 *
 */

// @lc code=start
func wordPattern(pattern string, s string) bool {
	mp := make(map[string]string)
	mp2 := make(map[string]string)
	ss := strings.Split(s, " ")
	if len(ss) != len(pattern) {
		return false
	}
	for i, v := range pattern {
		if v2, ok := mp[string(v)]; ok {
			if v2 != ss[i] {
				return false
			}
		} else {
			mp[string(v)] = ss[i]
		}
		if v2, ok := mp2[ss[i]]; ok {
			if v2 != string(v) {
				return false
			}
		} else {
			mp2[ss[i]] = string(v)
		}
	}
	return true
}

// @lc code=end
