package leetcode

// 78/78 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 99.85 % of golang submissions (2.2 MB)
/*
 * @lc app=leetcode.cn id=28 lang=golang
 *
 * [28] 实现 strStr()
 *
 * https://leetcode-cn.com/problems/implement-strstr/description/
 *
 * algorithms
 * Easy (39.59%)
 * Likes:    623
 * Dislikes: 0
 * Total Accepted:    259.1K
 * Total Submissions: 654.5K
 * Testcase Example:  '"hello"\n"ll"'
 *
 * 实现 strStr() 函数。
 *
 * 给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置
 * (从0开始)。如果不存在，则返回  -1。
 *
 * 示例 1:
 *
 * 输入: haystack = "hello", needle = "ll"
 * 输出: 2
 *
 *
 * 示例 2:
 *
 * 输入: haystack = "aaaaa", needle = "bba"
 * 输出: -1
 *
 *
 * 说明:
 *
 * 当 needle 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。
 *
 * 对于本题而言，当 needle 是空字符串时我们应当返回 0 。这与C语言的 strstr() 以及 Java的 indexOf() 定义相符。
 *
 */

// @lc code=start
func strStr(haystack string, needle string) int {

	lenHaystack, lenNeedle := len(haystack), len(needle)

	if lenNeedle == 0 {
		return 0
	}
	if lenHaystack == 0 ||
		lenHaystack < lenNeedle {
		return -1
	}
	if lenHaystack == lenNeedle {
		if haystack != needle {
			return -1
		}
		return 0
	}

	for index := 0; index < lenHaystack-lenNeedle+1; index++ {
		for j := 0; j < lenNeedle; j++ {
			if haystack[index+j] != needle[j] {
				break
			}
			if j == lenNeedle-1 {
				return index
			}
		}
	}
	return -1
}

// @lc code=end
