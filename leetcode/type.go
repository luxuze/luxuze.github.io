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
	var (
		answer string
		layers = []*TreeNode{root}
	)

	for layers != nil {
		var temp []*TreeNode
		for i := range layers {
			if layers[i] == nil {
				answer = fmt.Sprintf("%s,null", answer)
				continue
			}
			answer = fmt.Sprintf("%s,%d", answer, layers[i].Val)
			temp = append(temp, layers[i].Left, layers[i].Right)
		}
		layers = temp
	}

	return answer[1:]
}

// Deserialize Deserialize your encoded data to tree.
func Deserialize(data string) *TreeNode {
	var (
		atoiTreeNode = func(s string) *TreeNode {
			if s == "null" {
				return nil
			}
			i, _ := strconv.Atoi(s)
			return &TreeNode{Val: i}
		}

		list   = strings.Split(data, ",")
		root   = atoiTreeNode(list[0])
		tnList = []*TreeNode{root}
	)

	list = list[1:]

	for len(list) > 0 {
		tnList[0].Left = atoiTreeNode(list[0])
		tnList[0].Right = atoiTreeNode(list[1])
		if tnList[0].Left != nil {
			tnList = append(tnList, tnList[0].Left)
		}
		if tnList[0].Right != nil {
			tnList = append(tnList, tnList[0].Right)
		}
		list = list[2:]
		tnList = tnList[1:]
	}

	return root
}
