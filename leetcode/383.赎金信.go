package leetcode

// 126/126 cases passed (12 ms)
// Your runtime beats 40.9 % of golang submissions
// Your memory usage beats 27.79 % of golang submissions (3.8 MB)
/*
 * @lc app=leetcode.cn id=383 lang=golang
 *
 * [383] 赎金信
 *
 * https://leetcode-cn.com/problems/ransom-note/description/
 *
 * algorithms
 * Easy (64.32%)
 * Likes:    263
 * Dislikes: 0
 * Total Accepted:    121.2K
 * Total Submissions: 188.4K
 * Testcase Example:  '"a"\n"b"'
 *
 * 给你两个字符串：ransomNote 和 magazine ，判断 ransomNote 能不能由 magazine 里面的字符构成。
 *
 * 如果可以，返回 true ；否则返回 false 。
 *
 * magazine 中的每个字符只能在 ransomNote 中使用一次。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：ransomNote = "a", magazine = "b"
 * 输出：false
 *
 *
 * 示例 2：
 *
 *
 * 输入：ransomNote = "aa", magazine = "ab"
 * 输出：false
 *
 *
 * 示例 3：
 *
 *
 * 输入：ransomNote = "aa", magazine = "aab"
 * 输出：true
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= ransomNote.length, magazine.length <= 10^5
 * ransomNote 和 magazine 由小写英文字母组成
 *
 *
 */

// @lc code=start
func canConstruct(ransomNote string, magazine string) bool {
	mp := make(map[rune]int)
	for _, v := range magazine {
		mp[v]++
	}
	for _, v := range ransomNote {
		mp[v]--
	}
	for _, v := range mp {
		if v < 0 {
			return false
		}
	}
	return true
}

// @lc code=end
