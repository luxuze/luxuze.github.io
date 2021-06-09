package leetcode

// 34/34 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 23.75 % of golang submissions (2.1 MB)
/*
 * @lc app=leetcode.cn id=119 lang=golang
 *
 * [119] 杨辉三角 II
 *
 * https://leetcode-cn.com/problems/pascals-triangle-ii/description/
 *
 * algorithms
 * Easy (65.38%)
 * Likes:    293
 * Dislikes: 0
 * Total Accepted:    118K
 * Total Submissions: 180.5K
 * Testcase Example:  '3'
 *
 * 给定一个非负索引 k，其中 k ≤ 33，返回杨辉三角的第 k 行。
 *
 *
 *
 * 在杨辉三角中，每个数是它左上方和右上方的数的和。
 *
 * 示例:
 *
 * 输入: 3
 * 输出: [1,3,3,1]
 *
 *
 * 进阶：
 *
 * 你可以优化你的算法到 O(k) 空间复杂度吗？
 *
 */

// @lc code=start
func getRow(numRows int) []int {
	answer := make([][]int, numRows+1)
	for i := 0; i <= numRows; i++ {
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
	return answer[numRows]
}

// @lc code=end
