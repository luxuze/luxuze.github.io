package leetcode

// 150/150 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 44.07 % of golang submissions (2.7 MB)
import "strings"

/*
 * @lc app=leetcode.cn id=331 lang=golang
 *
 * [331] 验证二叉树的前序序列化
 *
 * https://leetcode-cn.com/problems/verify-preorder-serialization-of-a-binary-tree/description/
 *
 * algorithms
 * Medium (45.94%)
 * Likes:    139
 * Dislikes: 0
 * Total Accepted:    10K
 * Total Submissions: 21.7K
 * Testcase Example:  '"9,3,4,#,#,1,#,#,2,#,6,#,#"'
 *
 * 序列化二叉树的一种方法是使用前序遍历。当我们遇到一个非空节点时，我们可以记录下这个节点的值。如果它是一个空节点，我们可以使用一个标记值记录，例如 #。
 *
 * ⁠    _9_
 * ⁠   /   \
 * ⁠  3     2
 * ⁠ / \   / \
 * ⁠4   1  #  6
 * / \ / \   / \
 * # # # #   # #
 *
 *
 * 例如，上面的二叉树可以被序列化为字符串 "9,3,4,#,#,1,#,#,2,#,6,#,#"，其中 # 代表一个空节点。
 *
 * 给定一串以逗号分隔的序列，验证它是否是正确的二叉树的前序序列化。编写一个在不重构树的条件下的可行算法。
 *
 * 每个以逗号分隔的字符或为一个整数或为一个表示 null 指针的 '#' 。
 *
 * 你可以认为输入格式总是有效的，例如它永远不会包含两个连续的逗号，比如 "1,,3" 。
 *
 * 示例 1:
 *
 * 输入: "9,3,4,#,#,1,#,#,2,#,6,#,#"
 * 输出: true
 *
 * 示例 2:
 *
 * 输入: "1,#"
 * 输出: false
 *
 *
 * 示例 3:
 *
 * 输入: "9,#,#,1"
 * 输出: false
 *
 */

// @lc code=start
func isValidSerialization(preorder string) bool {
	var flag int = 1
	tnList := strings.Split(preorder, ",")
	for i := range tnList {
		flag--
		if flag < 0 {
			return false
		}
		if tnList[i] != "#" {
			flag += 2
		}
	}
	return flag == 0
}

// @lc code=end
