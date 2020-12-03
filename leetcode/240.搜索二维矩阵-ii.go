package leetcode

// 128/128 cases passed (28 ms)
// Your runtime beats 84.68 % of golang submissions
// Your memory usage beats 69.05 % of golang submissions (6.6 MB)
/*
 * @lc app=leetcode.cn id=240 lang=golang
 *
 * [240] 搜索二维矩阵 II
 *
 * https://leetcode-cn.com/problems/search-a-2d-matrix-ii/description/
 *
 * algorithms
 * Medium (42.65%)
 * Likes:    489
 * Dislikes: 0
 * Total Accepted:    90.4K
 * Total Submissions: 211.8K
 * Testcase Example:  '[[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]]\n' +
  '5'
 *
 * 编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target 。该矩阵具有以下特性：
 *
 *
 * 每行的元素从左到右升序排列。
 * 每列的元素从上到下升序排列。
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：matrix =
 * [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]],
 * target = 5
 * 输出：true
 *
 *
 * 示例 2：
 *
 *
 * 输入：matrix =
 * [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]],
 * target = 20
 * 输出：false
 *
 *
 *
 *
 * 提示：
 *
 *
 * m == matrix.length
 * n == matrix[i].length
 * 1
 * -10^9
 * 每行的所有元素从左到右升序排列
 * 每列的所有元素从上到下升序排列
 * -10^9
 *
 *
*/

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	for y := 0; y < len(matrix); y++ {
		if matrix[y][0] > target {
			return false
		}
		for x := 0; x < len(matrix[0]); x++ {
			if matrix[y][x] > target {
				break
			}
			if matrix[y][x] == target {
				return true
			}
		}
	}
	return false
}

// @lc code=end
