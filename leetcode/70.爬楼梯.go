package leetcode

// 45/45 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 17.05 % of golang submissions (1.9 MB)
/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 *
 * https://leetcode-cn.com/problems/climbing-stairs/description/
 *
 * algorithms
 * Easy (52.05%)
 * Likes:    1683
 * Dislikes: 0
 * Total Accepted:    459.8K
 * Total Submissions: 883.4K
 * Testcase Example:  '2'
 *
 * 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
 *
 * 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
 *
 * 注意：给定 n 是一个正整数。
 *
 * 示例 1：
 *
 * 输入： 2
 * 输出： 2
 * 解释： 有两种方法可以爬到楼顶。
 * 1.  1 阶 + 1 阶
 * 2.  2 阶
 *
 * 示例 2：
 *
 * 输入： 3
 * 输出： 3
 * 解释： 有三种方法可以爬到楼顶。
 * 1.  1 阶 + 1 阶 + 1 阶
 * 2.  1 阶 + 2 阶
 * 3.  2 阶 + 1 阶
 *
 *
 */

// @lc code=start
// func climbStairs(n int) (ans int) {

// 	var dfs func(x int)
// 	dfs = func(x int) {

// 		if x >= n {
// 			if x == n {
// 				ans++
// 			}
// 			return
// 		}

// 		dfs(x + 1)
// 		dfs(x + 2)
// 	}
// 	dfs(1)
// 	dfs(2)

// 	return
// }
func climbStairs(n int) (ans int) {
	dp := make([]int, n+2)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 2

	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

// @lc code=end
