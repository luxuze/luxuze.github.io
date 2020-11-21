package leetcode

// 8/8 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 14.29 % of golang submissions (2.8 MB)
/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 */

// @lc code=start
func generateParenthesis(n int) (result []string) {
	dfs(n, n, "", &result)
	return
}

func dfs(left int, right int, path string, result *[]string) {
	if left == 0 && right == 0 {
		*result = append(*result, path)
		return
	}
	if left > 0 {
		dfs(left-1, right, path+"(", result)
	}
	if right > left {
		dfs(left, right-1, path+")", result)
	}
}

// @lc code=end
