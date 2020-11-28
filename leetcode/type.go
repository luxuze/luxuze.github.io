package leetcode

// ListNode 链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

// TreeNode 二叉树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Node N叉树
type Node struct {
	Val      int
	Children []*Node
}
