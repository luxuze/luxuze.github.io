package leetcode

// 36/36 cases passed (8 ms)
// Your runtime beats 17.78 % of golang submissions
// Your memory usage beats 100 % of golang submissions (5.5 MB)
import (
	"sort"
	"strings"
)

/*
 * @lc app=leetcode.cn id=893 lang=golang
 *
 * [893] 特殊等价字符串组
 */

// @lc code=start
func numSpecialEquivGroups(A []string) int {
	var mp = make(map[string]int)
	for ai := range A {
		var arr1 []string
		var arr2 []string
		for aii := range A[ai] {
			if aii%2 == 0 {
				arr1 = append(arr1, string(A[ai][aii]))
			} else {
				arr2 = append(arr2, string(A[ai][aii]))
			}
		}
		sort.Slice(arr1, func(i int, j int) bool { return arr1[i] < arr1[j] })
		sort.Slice(arr2, func(i int, j int) bool { return arr2[i] < arr2[j] })
		mp[strings.Join(arr1, "")+strings.Join(arr2, "")] = 1
	}
	return len(mp)
}

// @lc code=end
