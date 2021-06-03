package leetcode

// 25/25 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 96.03 % of golang submissions (2 MB)
/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 *
 * https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/description/
 *
 * algorithms
 * Medium (56.02%)
 * Likes:    1341
 * Dislikes: 0
 * Total Accepted:    276.6K
 * Total Submissions: 487.6K
 * Testcase Example:  '"23"'
 *
 * 给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。
 *
 * 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
 *
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：digits = "23"
 * 输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]
 *
 *
 * 示例 2：
 *
 *
 * 输入：digits = ""
 * 输出：[]
 *
 *
 * 示例 3：
 *
 *
 * 输入：digits = "2"
 * 输出：["a","b","c"]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0
 * digits[i] 是范围 ['2', '9'] 的一个数字。
 *
 *
 */

// @lc code=start
func letterCombinations(digits string) (ans []string) {
	if digits == "" {
		return
	}
	mp := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
	length := len(digits)
	var dfs func(index int, ans_item string)
	dfs = func(index int, ans_item string) {
		if index == length {
			ans = append(ans, ans_item)
			return
		}
		for _, letter := range mp[digits[index]] {
			dfs(index+1, ans_item+string(letter))
		}
	}
	dfs(0, "")
	return
}

// @lc code=end
