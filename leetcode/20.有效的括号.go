package leetcode

// 76/76 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 100 % of golang submissions (2 MB)
/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */

// @lc code=start
// func isValid(s string) bool {
// 	if s == "" {
// 		return true
// 	}
// 	re := regexp.MustCompile(`\(\)|\[\]|\{\}`)
// 	if !re.MatchString(s) {
// 		return false
// 	}
// 	return isValid(re.ReplaceAllString(s, ""))
// }

var mp = map[byte]byte{
	'(': ')',
	'{': '}',
	'[': ']',
}

func isValid(s string) bool {
	var stack []byte
	for i := range s {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			stack = append(stack, s[i])
			continue
		}
		if len(stack) == 0 {
			return false
		}
		if mp[stack[len(stack)-1]] == s[i] {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	if len(stack) != 0 {
		return false
	}
	return true
}

// @lc code=end
