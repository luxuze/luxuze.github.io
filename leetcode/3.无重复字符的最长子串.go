package leetcode

// 987/987 cases passed (32 ms)
// Your runtime beats 14.8 % of golang submissions
// Your memory usage beats 80.65 % of golang submissions (2.6 MB)
import (
	"strings"
)

/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	var result int
	for left := 0; left < len(s); left++ {
		for right := left; right < len(s); right++ {
			if strings.IndexByte(s[left:right], s[right]) != -1 {
				break
			}
			if result < right-left {
				result = right - left
			}
		}
	}
	return result + 1
}

// @lc code=end
