package leetcode

// 478/478 cases passed (36 ms)
// Your runtime beats 94.67 % of golang submissions
// Your memory usage beats 87.83 % of golang submissions (6.5 MB)
/*
 * @lc app=leetcode.cn id=344 lang=golang
 *
 * [344] 反转字符串
 *
 * https://leetcode-cn.com/problems/reverse-string/description/
 *
 * algorithms
 * Easy (73.74%)
 * Likes:    327
 * Dislikes: 0
 * Total Accepted:    226.8K
 * Total Submissions: 307.4K
 * Testcase Example:  '["h","e","l","l","o"]'
 *
 * 编写一个函数，其作用是将输入的字符串反转过来。输入字符串以字符数组 char[] 的形式给出。
 *
 * 不要给另外的数组分配额外的空间，你必须原地修改输入数组、使用 O(1) 的额外空间解决这一问题。
 *
 * 你可以假设数组中的所有字符都是 ASCII 码表中的可打印字符。
 *
 *
 *
 * 示例 1：
 *
 * 输入：["h","e","l","l","o"]
 * 输出：["o","l","l","e","h"]
 *
 *
 * 示例 2：
 *
 * 输入：["H","a","n","n","a","h"]
 * 输出：["h","a","n","n","a","H"]
 *
 */

// @lc code=start
func reverseString(s []byte) {
	var tmp byte
	var lenS int = len(s) - 1
	if lenS == -1 {
		return
	}
	for i := 0; i < (lenS/2)+1; i++ {
		tmp = s[i]
		s[i] = s[lenS-i]
		s[lenS-i] = tmp
	}
}

// @lc code=end
