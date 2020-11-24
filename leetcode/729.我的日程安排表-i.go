package leetcode

// 108/108 cases passed (104 ms)
// Your runtime beats 68.75 % of golang submissions
// Your memory usage beats 56.25 % of golang submissions (7.3 MB)
/*
 * @lc app=leetcode.cn id=729 lang=golang
 *
 * [729] 我的日程安排表 I
 *
 * https://leetcode-cn.com/problems/my-calendar-i/description/
 *
 * algorithms
 * Medium (49.23%)
 * Likes:    58
 * Dislikes: 0
 * Total Accepted:    4.7K
 * Total Submissions: 9.5K
 * Testcase Example:  '["MyCalendar","book","book","book"]\n[[],[10,20],[15,25],[20,30]]'
 *
 * 实现一个 MyCalendar 类来存放你的日程安排。如果要添加的时间内没有其他安排，则可以存储这个新的日程安排。
 *
 * MyCalendar 有一个 book(int start, int end)方法。它意味着在 start 到 end
 * 时间内增加一个日程安排，注意，这里的时间是半开区间，即 [start, end), 实数 x 的范围为，  start <= x < end。
 *
 * 当两个日程安排有一些时间上的交叉时（例如两个日程安排都在同一时间内），就会产生重复预订。
 *
 * 每次调用 MyCalendar.book方法时，如果可以将日程安排成功添加到日历中而不会导致重复预订，返回 true。否则，返回 false
 * 并且不要将该日程安排添加到日历中。
 *
 * 请按照以下步骤调用 MyCalendar 类: MyCalendar cal = new MyCalendar();
 * MyCalendar.book(start, end)
 *
 * 示例 1:
 *
 * MyCalendar();
 * MyCalendar.book(10, 20); // returns true
 * MyCalendar.book(15, 25); // returns false
 * MyCalendar.book(20, 30); // returns true
 * 解释:
 * 第一个日程安排可以添加到日历中.  第二个日程安排不能添加到日历中，因为时间 15 已经被第一个日程安排预定了。
 * 第三个日程安排可以添加到日历中，因为第一个日程安排并不包含时间 20 。
 *
 *
 * 说明:
 *
 *
 * 每个测试用例，调用 MyCalendar.book 函数最多不超过 100次。
 * 调用函数 MyCalendar.book(start, end)时， start 和 end 的取值范围为 [0, 10^9]。
 *
 *
 */

// @lc code=start

type myCalendarItem struct {
	Start, End int
}
type MyCalendar struct {
	Calendar []myCalendarItem
}

func MyCalendar_Constructor() MyCalendar {
	return MyCalendar{}
}

func (this *MyCalendar) Book(start int, end int) bool {
	for i := range this.Calendar {
		if this.Calendar[i].Start <= start && start < this.Calendar[i].End {
			return false
		}
		if this.Calendar[i].Start < end && start < this.Calendar[i].End {
			return false
		}
	}
	this.Calendar = append(this.Calendar, myCalendarItem{start, end})
	return true
}

/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
// @lc code=end
