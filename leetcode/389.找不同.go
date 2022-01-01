package leetcode

// 54/54 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 69.76 % of golang submissions (2.1 MB)
/*
 * @lc app=leetcode.cn id=389 lang=golang
 *
 * [389] 找不同
 *
 * https://leetcode-cn.com/problems/find-the-difference/description/
 *
 * algorithms
 * Easy (69.01%)
 * Likes:    287
 * Dislikes: 0
 * Total Accepted:    108.7K
 * Total Submissions: 157.4K
 * Testcase Example:  '"abcd"\n"abcde"'
 *
 * 给定两个字符串 s 和 t，它们只包含小写字母。
 *
 * 字符串 t 由字符串 s 随机重排，然后在随机位置添加一个字母。
 *
 * 请找出在 t 中被添加的字母。
 *
 *
 *
 * 示例 1：
 *
 * 输入：s = "abcd", t = "abcde"
 * 输出："e"
 * 解释：'e' 是那个被添加的字母。
 *
 *
 * 示例 2：
 *
 * 输入：s = "", t = "y"
 * 输出："y"
 *
 *
 * 示例 3：
 *
 * 输入：s = "a", t = "aa"
 * 输出："a"
 *
 *
 * 示例 4：
 *
 * 输入：s = "ae", t = "aea"
 * 输出："a"
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0 <= s.length <= 1000
 * t.length == s.length + 1
 * s 和 t 只包含小写字母
 *
 *
 */

// @lc code=start
func findTheDifference(s string, t string) byte {
	var ans int64
	for _, v := range t {
		ans += int64(v)
	}
	for _, v := range s {
		ans -= int64(v)
	}
	return byte(ans)
}

// @lc code=end
