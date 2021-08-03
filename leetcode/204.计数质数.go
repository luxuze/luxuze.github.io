package leetcode

// 21/21 cases passed (624 ms)
// Your runtime beats 5.2 % of golang submissions
// Your memory usage beats 99.58 % of golang submissions (2 MB)
/*
 * @lc app=leetcode.cn id=204 lang=golang
 *
 * [204] 计数质数
 *
 * https://leetcode-cn.com/problems/count-primes/description/
 *
 * algorithms
 * Easy (38.04%)
 * Likes:    721
 * Dislikes: 0
 * Total Accepted:    158.6K
 * Total Submissions: 417.4K
 * Testcase Example:  '10'
 *
 * 统计所有小于非负整数 n 的质数的数量。
 *
 *
 *
 * 示例 1：
 *
 * 输入：n = 10
 * 输出：4
 * 解释：小于 10 的质数一共有 4 个, 它们是 2, 3, 5, 7 。
 *
 *
 * 示例 2：
 *
 * 输入：n = 0
 * 输出：0
 *
 *
 * 示例 3：
 *
 * 输入：n = 1
 * 输出：0
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0 <= n <= 5 * 10^6
 *
 *
 */

// @lc code=start
// func countPrimes(n int) int {
// 	var ans int
// 	isPrime := func(x int) bool {
// 		for i := 2; i*i <= x; i++ {
// 			if x%i == 0 {
// 				return false
// 			}
// 		}
// 		return true
// 	}
// 	for i := 2; i < n; i++ {
// 		if isPrime(i) {
// 			ans++
// 		}
// 	}
// 	return ans
// }

func countPrimes(n int) (ans int) {
	isPrime := make([]bool, n)
	for i := range isPrime {
		isPrime[i] = true
	}
	for i := 2; i < n; i++ {
		if isPrime[i] {
			ans++
			for j := 2 * i; j < n; j += i {
				isPrime[j] = false
			}
		}
	}
	return
}

// @lc code=end
