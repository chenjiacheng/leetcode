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

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	if root.Val > val {
		root.Left = insertIntoBST(root.Left, val)
	}
	if root.Val < val {
		root.Right = insertIntoBST(root.Right, val)
	}
	return root
}

type Example struct {
	root *TreeNode
	val  int
	ans  *TreeNode
}

func main() {
	examples := []Example{
		{&TreeNode{Val: 4, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 7}}, 5, &TreeNode{Val: 4, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 7, Left: &TreeNode{Val: 5}}}},
		{&TreeNode{Val: 40, Left: &TreeNode{Val: 20, Left: &TreeNode{Val: 10}, Right: &TreeNode{Val: 30}}, Right: &TreeNode{Val: 60, Left: &TreeNode{Val: 50}, Right: &TreeNode{Val: 70}}}, 25, &TreeNode{Val: 40, Left: &TreeNode{Val: 20, Left: &TreeNode{Val: 10}, Right: &TreeNode{Val: 30, Left: &TreeNode{Val: 25}}}, Right: &TreeNode{Val: 60, Left: &TreeNode{Val: 50}, Right: &TreeNode{Val: 70}}}},
	}
	for i, e := range examples {
		ans := insertIntoBST(e.root, e.val)
		if reflect.DeepEqual(ans, e.ans) {
			fmt.Println("PASS: CASE", i)
		} else {
			fmt.Println("FAIL: CASE", i)
		}
	}
}
