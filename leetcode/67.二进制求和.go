package leetcode

// 294/294 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 65.55 % of golang submissions (2.6 MB)
/*
 * @lc app=leetcode.cn id=67 lang=golang
 *
 * [67] 二进制求和
 *
 * https://leetcode-cn.com/problems/add-binary/description/
 *
 * algorithms
 * Easy (54.33%)
 * Likes:    523
 * Dislikes: 0
 * Total Accepted:    137.4K
 * Total Submissions: 252.8K
 * Testcase Example:  '"11"\n"1"'
 *
 * 给你两个二进制字符串，返回它们的和（用二进制表示）。
 *
 * 输入为 非空 字符串且只包含数字 1 和 0。
 *
 *
 *
 * 示例 1:
 *
 * 输入: a = "11", b = "1"
 * 输出: "100"
 *
 * 示例 2:
 *
 * 输入: a = "1010", b = "1011"
 * 输出: "10101"
 *
 *
 *
 * 提示：
 *
 *
 * 每个字符串仅由字符 '0' 或 '1' 组成。
 * 1 <= a.length, b.length <= 10^4
 * 字符串如果不是 "0" ，就都不含前导零。
 *
 *
 */

// @lc code=start
func addBinary(a string, b string) (answer string) {
	var max int
	var carry byte = '0'
	var cha int
	lenA, lenB := len(a), len(b)
	if lenA < lenB {
		max = lenB
		cha = lenB - lenA
	} else {
		max = lenA
		cha = lenA - lenB
	}

	get := func(s string, i int) byte {
		if len(s) == max {
			return s[max-i-1]
		}
		if i < max-cha {
			return s[max-i-cha-1]
		}
		return '0'
	}

	for i := 0; i < max; i++ {
		val := get(a, i) + get(b, i) + carry - '0'*3
		switch val {
		case 0:
			answer = "0" + answer
		case 1:
			answer = "1" + answer
			carry = '0'
		case 2:
			answer = "0" + answer
			carry = '1'
		case 3:
			answer = "1" + answer
			carry = '1'
		}
	}
	switch carry {
	case '1':
		return "1" + answer
	case '2':
		return "10" + answer
	}
	return
}

// @lc code=end
