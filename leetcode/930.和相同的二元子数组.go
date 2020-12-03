package leetcode

// 59/59 cases passed (1272 ms)
// Your runtime beats 7.32 % of golang submissions
// Your memory usage beats 73.17 % of golang submissions (6.7 MB)
/*
 * @lc app=leetcode.cn id=930 lang=golang
 *
 * [930] 和相同的二元子数组
 *
 * https://leetcode-cn.com/problems/binary-subarrays-with-sum/description/
 *
 * algorithms
 * Medium (38.19%)
 * Likes:    74
 * Dislikes: 0
 * Total Accepted:    4.8K
 * Total Submissions: 12.6K
 * Testcase Example:  '[1,0,1,0,1]\n2'
 *
 * 在由若干 0 和 1  组成的数组 A 中，有多少个和为 S 的非空子数组。
 *
 *
 *
 * 示例：
 *
 * 输入：A = [1,0,1,0,1], S = 2
 * 输出：4
 * 解释：
 * 如下面黑体所示，有 4 个满足题目要求的子数组：
 * [1,0,1,0,1]
 * [1,0,1,0,1]
 * [1,0,1,0,1]
 * [1,0,1,0,1]
 *
 *
 *
 *
 * 提示：
 *
 *
 * A.length <= 30000
 * 0 <= S <= A.length
 * A[i] 为 0 或 1
 *
 *
 */

// @lc code=start
func numSubarraysWithSum(A []int, S int) (answer int) {
	for i := 0; i < len(A); i++ {
		var sum int
		for j := i; j < len(A) && sum <= S; j++ {
			sum += A[j]
			if sum == S {
				answer++
			}
		}
	}
	return
}

// @lc code=end
