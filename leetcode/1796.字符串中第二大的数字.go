package leetcode

// 301/301 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 27.27 % of golang submissions (2.3 MB)
import "sort"

/*
 * @lc app=leetcode.cn id=1796 lang=golang
 *
 * [1796] 字符串中第二大的数字
 *
 * https://leetcode-cn.com/problems/second-largest-digit-in-a-string/description/
 *
 * algorithms
 * Easy (48.37%)
 * Likes:    2
 * Dislikes: 0
 * Total Accepted:    5.7K
 * Total Submissions: 11.7K
 * Testcase Example:  '"dfa12321afd"'
 *
 * 给你一个混合字符串 s ，请你返回 s 中 第二大 的数字，如果不存在第二大的数字，请你返回 -1 。
 *
 * 混合字符串 由小写英文字母和数字组成。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "dfa12321afd"
 * 输出：2
 * 解释：出现在 s 中的数字包括 [1, 2, 3] 。第二大的数字是 2 。
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "abc1111"
 * 输出：-1
 * 解释：出现在 s 中的数字只包含 [1] 。没有第二大的数字。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * s 只包含小写英文字母和（或）数字。
 *
 *
 */

// @lc code=start
func secondHighest(s string) int {
	var arr []int
	mp := make(map[byte]struct{})

	for i := range s {
		if '0' <= s[i] && s[i] <= '9' {
			if _, ok := mp[s[i]]; ok {
				continue
			}
			mp[s[i]] = struct{}{}
			arr = append(arr, int(s[i]-'0'))
		}
	}
	if len(arr) < 2 {
		return -1
	}
	sort.Ints(arr)

	return arr[len(arr)-2]
}

// @lc code=end
