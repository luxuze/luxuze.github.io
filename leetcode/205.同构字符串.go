package leetcode

// 43/43 cases passed (8 ms)
// Your runtime beats 47.43 % of golang submissions
// Your memory usage beats 71.69 % of golang submissions (2.6 MB)
/*
 * @lc app=leetcode.cn id=205 lang=golang
 *
 * [205] 同构字符串
 *
 * https://leetcode-cn.com/problems/isomorphic-strings/description/
 *
 * algorithms
 * Easy (50.07%)
 * Likes:    373
 * Dislikes: 0
 * Total Accepted:    106.4K
 * Total Submissions: 212.6K
 * Testcase Example:  '"egg"\n"add"'
 *
 * 给定两个字符串 s 和 t，判断它们是否是同构的。
 *
 * 如果 s 中的字符可以按某种映射关系替换得到 t ，那么这两个字符串是同构的。
 *
 *
 * 每个出现的字符都应当映射到另一个字符，同时不改变字符的顺序。不同字符不能映射到同一个字符上，相同字符只能映射到同一个字符上，字符可以映射到自己本身。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入：s = "egg", t = "add"
 * 输出：true
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "foo", t = "bar"
 * 输出：false
 *
 * 示例 3：
 *
 *
 * 输入：s = "paper", t = "title"
 * 输出：true
 *
 *
 *
 * 提示：
 *
 *
 * 可以假设 s 和 t 长度相同。
 *
 *
 */

// @lc code=start
func isIsomorphic(s string, t string) bool {
	length := len(s)
	if length != len(t) {
		return false
	}
	m1, m2 := make(map[byte]byte), make(map[byte]byte)
	for i := 0; i < length; i++ {
		a, b := s[i], t[i]
		if m1[a] > 0 && m1[a] != b || m2[b] > 0 && m2[b] != a {
			return false
		}
		m1[a] = b
		m2[b] = a
		continue
	}
	return true
}

// @lc code=end
