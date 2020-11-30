package leetcode

import (
	"fmt"
	"strconv"
	"strings"
)

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

// Serialize Serializes a tree to a single string.
func Serialize(root *TreeNode) string {
	if root == nil {
		return "null"
	}
	return fmt.Sprintf("%d,%s,%s",
		root.Val, Serialize(root.Left), Serialize(root.Right))
}

// Deserialize Deserialize your encoded data to tree.
func Deserialize(data string) *TreeNode {
	var dfs func(*[]string) *TreeNode
	dfs = func(l *[]string) *TreeNode {
		rootVal := (*l)[0]
		*l = (*l)[1:]
		if rootVal == "null" {
			return nil
		}
		val, _ := strconv.Atoi(rootVal)
		root := &TreeNode{Val: val}
		root.Left = dfs(l)
		root.Right = dfs(l)
		return root
	}
	l := strings.Split(data, ",")
	return dfs(&l)
}
