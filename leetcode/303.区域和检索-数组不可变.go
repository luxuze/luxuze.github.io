package leetcode

// 15/15 cases passed (92 ms)
// Your runtime beats 7.07 % of golang submissions
// Your memory usage beats 62.77 % of golang submissions (9.3 MB)
/*
 * @lc app=leetcode.cn id=303 lang=golang
 *
 * [303] 区域和检索 - 数组不可变
 *
 * https://leetcode-cn.com/problems/range-sum-query-immutable/description/
 *
 * algorithms
 * Easy (73.29%)
 * Likes:    404
 * Dislikes: 0
 * Total Accepted:    139.5K
 * Total Submissions: 190.4K
 * Testcase Example:  '["NumArray","sumRange","sumRange","sumRange"]\n' +
  '[[[-2,0,3,-5,2,-1]],[0,2],[2,5],[0,5]]'
 *
 * 给定一个整数数组  nums，求出数组从索引 i 到 j（i ≤ j）范围内元素的总和，包含 i、j 两点。
 *
 *
 *
 * 实现 NumArray 类：
 *
 *
 * NumArray(int[] nums) 使用数组 nums 初始化对象
 * int sumRange(int i, int j) 返回数组 nums 从索引 i 到 j（i ≤ j）范围内元素的总和，包含 i、j 两点（也就是
 * sum(nums[i], nums[i + 1], ... , nums[j])）
 *
 *
 *
 *
 * 示例：
 *
 *
 * 输入：
 * ["NumArray", "sumRange", "sumRange", "sumRange"]
 * [[[-2, 0, 3, -5, 2, -1]], [0, 2], [2, 5], [0, 5]]
 * 输出：
 * [null, 1, -1, -3]
 *
 * 解释：
 * NumArray numArray = new NumArray([-2, 0, 3, -5, 2, -1]);
 * numArray.sumRange(0, 2); // return 1 ((-2) + 0 + 3)
 * numArray.sumRange(2, 5); // return -1 (3 + (-5) + 2 + (-1))
 * numArray.sumRange(0, 5); // return -3 ((-2) + 0 + 3 + (-5) + 2 + (-1))
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0
 * -10^5
 * 0
 * 最多调用 10^4 次 sumRange 方法
 *
 *
 *
 *
*/

// @lc code=start
type NumArray struct {
	nums *[]int
}

func NumArray_Constructor(nums []int) NumArray {
	return NumArray{nums: &nums}
}

func (this *NumArray) SumRange(left int, right int) int {
	var ans int
	for left <= right {
		ans += (*this.nums)[left]
		left++
	}
	return ans
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
// @lc code=end
