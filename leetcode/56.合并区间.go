package leetcode

// 168/168 cases passed (20 ms)
// Your runtime beats 12.58 % of golang submissions
// Your memory usage beats 54.24 % of golang submissions (4.6 MB)
import (
	"sort"
)

/*
 * @lc app=leetcode.cn id=56 lang=golang
 *
 * [56] 合并区间
 *
 * https://leetcode-cn.com/problems/merge-intervals/description/
 *
 * algorithms
 * Medium (43.46%)
 * Likes:    705
 * Dislikes: 0
 * Total Accepted:    166.9K
 * Total Submissions: 384K
 * Testcase Example:  '[[1,3],[2,6],[8,10],[15,18]]'
 *
 * 给出一个区间的集合，请合并所有重叠的区间。
 *
 *
 *
 * 示例 1:
 *
 * 输入: intervals = [[1,3],[2,6],[8,10],[15,18]]
 * 输出: [[1,6],[8,10],[15,18]]
 * 解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
 *
 *
 * 示例 2:
 *
 * 输入: intervals = [[1,4],[4,5]]
 * 输出: [[1,5]]
 * 解释: 区间 [1,4] 和 [4,5] 可被视为重叠区间。
 *
 * 注意：输入类型已于2019年4月15日更改。 请重置默认代码定义以获取新方法签名。
 *
 *
 *
 * 提示：
 *
 *
 * intervals[i][0] <= intervals[i][1]
 *
 *
 */

// @lc code=start
func merge(intervals [][]int) (ans [][]int) {
	sort.SliceStable(
		intervals,
		func(i, j int) bool {
			return intervals[i][0] < intervals[j][0]
		})
	tmp := intervals[0]

	for i := 1; i < len(intervals); i++ {
		if tmp[1] >= intervals[i][0] && tmp[1] <= intervals[i][1] {
			tmp[1] = intervals[i][1]
			continue
		}
		if tmp[1] >= intervals[i][0] && tmp[1] >= intervals[i][1] {
			continue
		}

		ans = append(ans, tmp)
		tmp = intervals[i]
	}
	if tmp == nil {
		return ans
	}

	return append(ans, tmp)
}

// @lc code=end
