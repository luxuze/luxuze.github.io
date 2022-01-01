package leetcode

// 104/104 cases passed (36 ms)
// Your runtime beats 33.04 % of golang submissions
// Your memory usage beats 24.6 % of golang submissions (5.4 MB)
/*
 * @lc app=leetcode.cn id=387 lang=golang
 *
 * [387] 字符串中的第一个唯一字符
 *
 * https://leetcode-cn.com/problems/first-unique-character-in-a-string/description/
 *
 * algorithms
 * Easy (53.90%)
 * Likes:    484
 * Dislikes: 0
 * Total Accepted:    249.4K
 * Total Submissions: 462.7K
 * Testcase Example:  '"leetcode"'
 *
 * 给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。
 *
 *
 *
 * 示例：
 *
 * s = "leetcode"
 * 返回 0
 *
 * s = "loveleetcode"
 * 返回 2
 *
 *
 *
 *
 * 提示：你可以假定该字符串只包含小写字母。
 *
 */

// @lc code=start
func firstUniqChar(s string) int {
	idxMap := make(map[rune]int)
	for _, v := range s {
		idxMap[v] += 1
	}
	for i, v := range s {
		if idxMap[v] == 1 {
			return i
		}
	}
	return -1
}

// @lc code=end
