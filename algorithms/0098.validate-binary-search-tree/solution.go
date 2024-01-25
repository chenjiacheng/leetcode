package main

import (
	"fmt"
	"reflect"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	var prev *TreeNode // 保存上一个指针
	var travel func(node *TreeNode) bool
	travel = func(node *TreeNode) bool {
		if node == nil {
			return true
		}
		leftRes := travel(node.Left)
		// 当前值小于等于前一个节点的值，返回false
		if prev != nil && node.Val <= prev.Val {
			return false
		}
		prev = node
		rightRes := travel(node.Right)
		return leftRes && rightRes
	}
	return travel(root)
}

type Example struct {
	root *TreeNode
	ans  bool
}

func main() {
	examples := []Example{
		{&TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}, true},
		{&TreeNode{Val: 5, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 6}}}, false},
	}
	for i, e := range examples {
		ans := isValidBST(e.root)
		if reflect.DeepEqual(ans, e.ans) {
			fmt.Println("PASS: CASE", i)
		} else {
			fmt.Println("FAIL: CASE", i)
		}
	}
}
