package leetcode

// 58/58 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 7.2 % of golang submissions (2.2 MB)
import "strings"

/*
 * @lc app=leetcode.cn id=58 lang=golang
 *
 * [58] 最后一个单词的长度
 *
 * https://leetcode-cn.com/problems/length-of-last-word/description/
 *
 * algorithms
 * Easy (34.73%)
 * Likes:    336
 * Dislikes: 0
 * Total Accepted:    203.4K
 * Total Submissions: 585.5K
 * Testcase Example:  '"Hello World"'
 *
 * 给你一个字符串 s，由若干单词组成，单词之间用单个或多个连续的空格字符隔开。返回字符串中最后一个单词的长度。如果不存在最后一个单词，请返回 0 。
 *
 * 单词 是指仅由字母组成、不包含任何空格字符的最大子字符串。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "Hello World"
 * 输出：5
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = " "
 * 输出：0
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 10^4
 * s 仅有英文字母和空格 ' ' 组成
 *
 *
 */

// @lc code=start
func lengthOfLastWord(s string) int {
	s = strings.TrimLeft(s, " ")
	s = strings.TrimRight(s, " ")
	arr := strings.Split(s, " ")
	if len(arr) == 0 {
		return 0
	}
	return len(arr[len(arr)-1])
}

// @lc code=end
