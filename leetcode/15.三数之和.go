package leetcode

// 318/318 cases passed (44 ms)
// Your runtime beats 34.6 % of golang submissions
// Your memory usage beats 81.5 % of golang submissions (7.3 MB)
import "sort"

/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 *
 * https://leetcode-cn.com/problems/3sum/description/
 *
 * algorithms
 * Medium (30.28%)
 * Likes:    2830
 * Dislikes: 0
 * Total Accepted:    384K
 * Total Submissions: 1.3M
 * Testcase Example:  '[-1,0,1,2,-1,-4]'
 *
 * 给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0
 * ？请你找出所有满足条件且不重复的三元组。
 *
 * 注意：答案中不可以包含重复的三元组。
 *
 *
 *
 * 示例：
 *
 * 给定数组 nums = [-1, 0, 1, 2, -1, -4]，
 *
 * 满足要求的三元组集合为：
 * [
 * ⁠ [-1, 0, 1],
 * ⁠ [-1, -1, 2]
 * ]
 *
 *
 */

// @lc code=start
func threeSum(nums []int) (answer [][]int) {
	l := len(nums)
	sort.Ints(nums)
	for i := 0; i < l-2; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		target := -1 * nums[i]
		k := l - 1
		for j := i + 1; j < l; j++ {
			if j >= i+2 && nums[j] == nums[j-1] {
				continue
			}
			for ; k > j && nums[j]+nums[k] > target; k-- {
			}
			if k == j {
				break
			}
			if nums[j]+nums[k] == target {
				answer = append(answer, []int{nums[i], nums[j], nums[k]})
			}
		}
	}
	return
}

// @lc code=end
