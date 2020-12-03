package leetcode

// 85/85 cases passed (40 ms)
// Your runtime beats 36.43 % of golang submissions
// Your memory usage beats 11.72 % of golang submissions (7.2 MB)
import (
	"sort"
)

/*
 * @lc app=leetcode.cn id=378 lang=golang
 *
 * [378] 有序矩阵中第K小的元素
 *
 * https://leetcode-cn.com/problems/kth-smallest-element-in-a-sorted-matrix/description/
 *
 * algorithms
 * Medium (62.87%)
 * Likes:    486
 * Dislikes: 0
 * Total Accepted:    58K
 * Total Submissions: 92.2K
 * Testcase Example:  '[[1,5,9],[10,11,13],[12,13,15]]\n8'
 *
 * 给定一个 n x n 矩阵，其中每行和每列元素均按升序排序，找到矩阵中第 k 小的元素。
 * 请注意，它是排序后的第 k 小元素，而不是第 k 个不同的元素。
 *
 *
 *
 * 示例：
 *
 * matrix = [
 * ⁠  [ 1,  5,  9],
 * ⁠  [10, 11, 13],
 * ⁠  [12, 13, 15]
 * ],
 * k = 8,
 *
 * 返回 13。
 *
 *
 *
 *
 * 提示：
 * 你可以假设 k 的值永远是有效的，1 ≤ k ≤ n^2 。
 *
 */

// @lc code=start
func kthSmallest(matrix [][]int, k int) int {
	var list []int
	for i := range matrix {
		list = append(list, matrix[i]...)
	}
	sort.Ints(list)
	return list[k-1]
}

// @lc code=end
