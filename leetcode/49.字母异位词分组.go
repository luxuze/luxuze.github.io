package leetcode

import (
	"sort"
)

// 114/114 cases passed (28 ms)
// Your runtime beats 75.77 % of golang submissions
// Your memory usage beats 34.71 % of golang submissions (8.4 MB)
/*
 * @lc app=leetcode.cn id=49 lang=golang
 *
 * [49] 字母异位词分组
 *
 * https://leetcode-cn.com/problems/group-anagrams/description/
 *
 * algorithms
 * Medium (65.92%)
 * Likes:    758
 * Dislikes: 0
 * Total Accepted:    196.4K
 * Total Submissions: 298K
 * Testcase Example:  '["eat","tea","tan","ate","nat","bat"]'
 *
 * 给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。
 *
 * 示例:
 *
 * 输入: ["eat", "tea", "tan", "ate", "nat", "bat"]
 * 输出:
 * [
 * ⁠ ["ate","eat","tea"],
 * ⁠ ["nat","tan"],
 * ⁠ ["bat"]
 * ]
 *
 * 说明：
 *
 *
 * 所有输入均为小写字母。
 * 不考虑答案输出的顺序。
 *
 *
 */

// @lc code=start
func groupAnagrams(strs []string) (ans [][]string) {
	mp := make(map[string][]string)
	for _, str := range strs {
		sb := []byte(str)
		sort.Slice(sb, func(i, j int) bool {
			return sb[i] < sb[j]
		})
		key := string(sb)
		mp[key] = append(mp[key], str)
	}
	for _, val := range mp {
		ans = append(ans, val)
	}
	return
}

// @lc code=end
