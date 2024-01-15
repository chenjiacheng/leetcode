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

func inorderTraversal(root *TreeNode) []int {
	var ans []int
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		traversal(node.Left)        // 左
		ans = append(ans, node.Val) // 中
		traversal(node.Right)       // 右
	}
	traversal(root)
	return ans
}

type Example struct {
	root *TreeNode
	ans  []int
}

func main() {
	examples := []Example{
		{&TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}, []int{1, 3, 2}},
		{nil, nil},
		{&TreeNode{Val: 1}, []int{1}},
	}
	for i, e := range examples {
		ans := inorderTraversal(e.root)
		if reflect.DeepEqual(ans, e.ans) {
			fmt.Println("PASS: CASE", i)
		} else {
			fmt.Println("FAIL: CASE", i)
		}
	}
}
