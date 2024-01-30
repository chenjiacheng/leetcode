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

func bstToGst(root *TreeNode) *TreeNode {
	pre := 0
	var travel func(node *TreeNode)
	travel = func(node *TreeNode) {
		if node == nil {
			return
		}
		travel(node.Right)
		node.Val += pre
		pre = node.Val
		travel(node.Left)
	}
	travel(root)
	return root
}

type Example struct {
	root *TreeNode
	ans  *TreeNode
}

func main() {
	examples := []Example{
		{&TreeNode{Val: 4, Left: &TreeNode{Val: 1, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}}, Right: &TreeNode{Val: 6, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 7, Right: &TreeNode{Val: 8}}}}, &TreeNode{Val: 30, Left: &TreeNode{Val: 36, Left: &TreeNode{Val: 36}, Right: &TreeNode{Val: 35, Right: &TreeNode{Val: 33}}}, Right: &TreeNode{Val: 21, Left: &TreeNode{Val: 26}, Right: &TreeNode{Val: 15, Right: &TreeNode{Val: 8}}}}},
		{&TreeNode{Val: 0, Right: &TreeNode{Val: 1}}, &TreeNode{Val: 1, Right: &TreeNode{Val: 1}}},
		{&TreeNode{Val: 1, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 2}}, &TreeNode{Val: 3, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 2}}},
		{&TreeNode{Val: 3, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}}, Right: &TreeNode{Val: 4}}, &TreeNode{Val: 7, Left: &TreeNode{Val: 9, Left: &TreeNode{Val: 10}}, Right: &TreeNode{Val: 4}}},
	}
	for i, e := range examples {
		ans := bstToGst(e.root)
		if reflect.DeepEqual(ans, e.ans) {
			fmt.Println("PASS: CASE", i)
		} else {
			fmt.Println("FAIL: CASE", i)
		}
	}
}
