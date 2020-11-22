package leetcode

// 315/315 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 93.96 % of golang submissions (2.5 MB)
/*
 * @lc app=leetcode.cn id=415 lang=golang
 *
 * [415] 字符串相加
 *
 * https://leetcode-cn.com/problems/add-strings/description/
 *
 * algorithms
 * Easy (51.84%)
 * Likes:    280
 * Dislikes: 0
 * Total Accepted:    84.8K
 * Total Submissions: 163.6K
 * Testcase Example:  '"0"\n"0"'
 *
 * 给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和。
 *
 *
 *
 * 提示：
 *
 *
 * num1 和num2 的长度都小于 5100
 * num1 和num2 都只包含数字 0-9
 * num1 和num2 都不包含任何前导零
 * 你不能使用任何內建 BigInteger 库， 也不能直接将输入的字符串转换为整数形式
 *
 *
 */

// @lc code=start
func addStrings(num1 string, num2 string) string {
	len1, len2 := len(num1), len(num2)
	if len1 > len2 { //keep num1 is smaller
		len1, len2 = len2, len1
		num1, num2 = num2, num1
	}

	result := make([]byte, len2+1)
	var jw byte

	for i := len2 - 1; i >= 0; i-- {
		j := i - (len2 - len1)
		var item byte
		if j >= 0 {
			jw, item = addBytes(jw, num1[j], num2[i])
		} else {
			jw, item = addBytes(jw, '0', num2[i])
		}
		result[i+1] = item
	}

	if jw != 0 {
		result[0] = jw + '0'
		return string(result)
	}

	return string(result[1:])
}

func addBytes(jw, s1, s2 byte) (byte, byte) {
	sum := jw + s1 + s2 - '0' - '0'
	if sum >= 10 {
		return 1, sum - 10 + '0'
	}
	return 0, sum + '0'
}

// @lc code=end
