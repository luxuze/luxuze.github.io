package leetcode

// 15/15 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 100 % of golang submissions (2.4 MB)
/*
 * @lc app=leetcode.cn id=413 lang=golang
 *
 * [413] 等差数列划分
 */

// @lc code=start
func numberOfArithmeticSlices(A []int) int {
	var sum int
	for i, j := 0, 1; j < len(A); j++ {
		if A[i+1]-A[i] == A[j]-A[j-1] {
			if j == len(A)-1 {
				return sum + numOfChildren(j-i+1)
			}
			continue
		}
		sum += numOfChildren(j - i)
		i = j - 1
	}
	return sum
}

func numOfChildren(length int) int {
	if length < 3 {
		return 0
	}
	if length == 3 {
		return 1
	}
	return (length - 1) * (length - 2) / 2
}

// @lc code=end
