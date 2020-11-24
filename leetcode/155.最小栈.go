package leetcode

// 18/18 cases passed (56 ms)
// Your runtime beats 5.56 % of golang submissions
// Your memory usage beats 22.38 % of golang submissions (9.5 MB)
/*
 * @lc app=leetcode.cn id=155 lang=golang
 *
 * [155] 最小栈
 *
 * https://leetcode-cn.com/problems/min-stack/description/
 *
 * algorithms
 * Easy (55.68%)
 * Likes:    735
 * Dislikes: 0
 * Total Accepted:    179.3K
 * Total Submissions: 321.8K
 * Testcase Example:  '["MinStack","push","push","push","getMin","pop","top","getMin"]\n' +
  '[[],[-2],[0],[-3],[],[],[],[]]'
 *
 * 设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。
 *
 *
 * push(x) —— 将元素 x 推入栈中。
 * pop() —— 删除栈顶的元素。
 * top() —— 获取栈顶元素。
 * getMin() —— 检索栈中的最小元素。
 *
 *
 *
 *
 * 示例:
 *
 * 输入：
 * ["MinStack","push","push","push","getMin","pop","top","getMin"]
 * [[],[-2],[0],[-3],[],[],[],[]]
 *
 * 输出：
 * [null,null,null,null,-3,null,0,-2]
 *
 * 解释：
 * MinStack minStack = new MinStack();
 * minStack.push(-2);
 * minStack.push(0);
 * minStack.push(-3);
 * minStack.getMin();   --> 返回 -3.
 * minStack.pop();
 * minStack.top();      --> 返回 0.
 * minStack.getMin();   --> 返回 -2.
 *
 *
 *
 *
 * 提示：
 *
 *
 * pop、top 和 getMin 操作总是在 非空栈 上调用。
 *
 *
*/

// @lc code=start
type MinStack struct {
	S []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	this.S = append(this.S, x)
}

func (this *MinStack) Pop() {
	if len(this.S) == 0 {
		return
	}
	this.S = this.S[:len(this.S)-1]
}

func (this *MinStack) Top() int {
	if len(this.S) == 0 {
		return 0
	}
	return this.S[len(this.S)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.S) == 0 {
		return 0
	}
	min := this.S[0]
	for i := range this.S {
		if this.S[i] <= min {
			min = this.S[i]
		}
	}
	return min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @lc code=end
