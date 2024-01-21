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

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftValue := 0
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		leftValue += root.Left.Val
	}
	return leftValue + sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right)
}

type Example struct {
	root *TreeNode
	ans  int
}

func main() {
	examples := []Example{
		{&TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}, 24},
		{&TreeNode{Val: 1}, 0},
	}
	for i, e := range examples {
		ans := sumOfLeftLeaves(e.root)
		if reflect.DeepEqual(ans, e.ans) {
			fmt.Println("PASS: CASE", i)
		} else {
			fmt.Println("FAIL: CASE", i)
		}
	}
}
