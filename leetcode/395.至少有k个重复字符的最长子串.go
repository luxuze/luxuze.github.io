package leetcode

// 28/28 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 100 % of golang submissions (2.1 MB)
/*
 * @lc app=leetcode.cn id=395 lang=golang
 *
 * [395] 至少有K个重复字符的最长子串
 */

// @lc code=start
func longestSubstring(s string, k int) int {
	if len(s) < k {
		return 0
	}

	var sumMap = make(map[byte]int)
	for i := range s {
		sumMap[s[i]]++
	}

	for i := range s {
		if sumMap[s[i]] < k {
			left := longestSubstring(s[:i], k)
			for j := i; j < len(s); j++ {
				if sumMap[s[j]] >= k {
					right := longestSubstring(s[j:], k)
					if left > right {
						return left
					}
					return right
				}
			}
			return left
		}
	}
	return len(s)
}

// @lc code=end
