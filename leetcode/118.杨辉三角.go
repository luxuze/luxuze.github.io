package leetcode

// 15/15 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 64.68 % of golang submissions (2 MB)
/*
 * @lc app=leetcode.cn id=118 lang=golang
 *
 * [118] 杨辉三角
 *
 * https://leetcode-cn.com/problems/pascals-triangle/description/
 *
 * algorithms
 * Easy (67.85%)
 * Likes:    379
 * Dislikes: 0
 * Total Accepted:    117.8K
 * Total Submissions: 173.5K
 * Testcase Example:  '5'
 *
 * 给定一个非负整数 numRows，生成杨辉三角的前 numRows 行。
 *
 *
 *
 * 在杨辉三角中，每个数是它左上方和右上方的数的和。
 *
 * 示例:
 *
 * 输入: 5
 * 输出:
 * [
 * ⁠    [1],
 * ⁠   [1,1],
 * ⁠  [1,2,1],
 * ⁠ [1,3,3,1],
 * ⁠[1,4,6,4,1]
 * ]
 *
 */

// @lc code=start
func generate(numRows int) [][]int {
	answer := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		answer[i] = make([]int, i+1)
	}
	for i := range answer {
		for j := range answer[i] {
			if j == 0 || j == len(answer[i])-1 {
				answer[i][j] = 1
			} else {
				answer[i][j] = answer[i-1][j-1] + answer[i-1][j]
			}
		}
	}
	return answer
}

// @lc code=end
