package leetcode

// 28/28 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 100 % of golang submissions (1.9 MB)
import "fmt"

/*
 * @lc app=leetcode.cn id=228 lang=golang
 *
 * [228] 汇总区间
 *
 * https://leetcode-cn.com/problems/summary-ranges/description/
 *
 * algorithms
 * Easy (57.60%)
 * Likes:    191
 * Dislikes: 0
 * Total Accepted:    62.5K
 * Total Submissions: 108.6K
 * Testcase Example:  '[0,1,2,4,5,7]'
 *
 * 给定一个无重复元素的有序整数数组 nums 。
 *
 * 返回 恰好覆盖数组中所有数字 的 最小有序 区间范围列表。也就是说，nums 的每个元素都恰好被某个区间范围所覆盖，并且不存在属于某个范围但不属于
 * nums 的数字 x 。
 *
 * 列表中的每个区间范围 [a,b] 应该按如下格式输出：
 *
 *
 * "a->b" ，如果 a != b
 * "a" ，如果 a == b
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [0,1,2,4,5,7]
 * 输出：["0->2","4->5","7"]
 * 解释：区间范围是：
 * [0,2] --> "0->2"
 * [4,5] --> "4->5"
 * [7,7] --> "7"
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [0,2,3,4,6,8,9]
 * 输出：["0","2->4","6","8->9"]
 * 解释：区间范围是：
 * [0,0] --> "0"
 * [2,4] --> "2->4"
 * [6,6] --> "6"
 * [8,9] --> "8->9"
 *
 *
 * 示例 3：
 *
 *
 * 输入：nums = []
 * 输出：[]
 *
 *
 * 示例 4：
 *
 *
 * 输入：nums = [-1]
 * 输出：["-1"]
 *
 *
 * 示例 5：
 *
 *
 * 输入：nums = [0]
 * 输出：["0"]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0
 * -2^31
 * nums 中的所有值都 互不相同
 * nums 按升序排列
 *
 *
 */

// @lc code=start
func summaryRanges(nums []int) (ans []string) {
	var i int
	for j := 0; j < len(nums); j++ {
		if j+1 == len(nums) || nums[i]-nums[j+1] != i-j-1 {
			item := fmt.Sprint(nums[i])
			if i != j {
				item = item + fmt.Sprintf("->%d", nums[j])
			}
			ans = append(ans, item)
			i = j + 1
		}
	}

	return
}

// @lc code=end
