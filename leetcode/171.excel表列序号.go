package leetcode

// 1002/1002 cases passed (4 ms)
// Your runtime beats 40.14 % of golang submissions
// Your memory usage beats 100 % of golang submissions (2.1 MB)
/*
 * @lc app=leetcode.cn id=171 lang=golang
 *
 * [171] Excel表列序号
 *
 * https://leetcode-cn.com/problems/excel-sheet-column-number/description/
 *
 * algorithms
 * Easy (69.59%)
 * Likes:    279
 * Dislikes: 0
 * Total Accepted:    101.4K
 * Total Submissions: 141.4K
 * Testcase Example:  '"A"'
 *
 * 给你一个字符串 columnTitle ，表示 Excel 表格中的列名称。返回该列名称对应的列序号。
 *
 *
 *
 * 例如，
 *
 *
 * ⁠   A -> 1
 * ⁠   B -> 2
 * ⁠   C -> 3
 * ⁠   ...
 * ⁠   Z -> 26
 * ⁠   AA -> 27
 * ⁠   AB -> 28
 * ⁠   ...
 *
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: columnTitle = "A"
 * 输出: 1
 *
 *
 * 示例 2:
 *
 *
 * 输入: columnTitle = "AB"
 * 输出: 28
 *
 *
 * 示例 3:
 *
 *
 * 输入: columnTitle = "ZY"
 * 输出: 701
 *
 * 示例 4:
 *
 *
 * 输入: columnTitle = "FXSHRXW"
 * 输出: 2147483647
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= columnTitle.length <= 7
 * columnTitle 仅由大写英文组成
 * columnTitle 在范围 ["A", "FXSHRXW"] 内
 *
 *
 */

// @lc code=start
func titleToNumber(columnTitle string) (ans int) {
	for i, j := len(columnTitle)-1, 1; i >= 0; i-- {
		d := int(columnTitle[i]-'A') + 1
		ans += d * j
		j *= 26
	}
	return
}

// @lc code=end
