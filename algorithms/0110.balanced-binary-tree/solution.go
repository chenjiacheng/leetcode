package main

import (
	"fmt"
	"math"
	"reflect"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	return getHeight(root) != -1
}

func getHeight(node *TreeNode) int {
	if node == nil {
		return 0
	}
	l := getHeight(node.Left)
	if l == -1 {
		return -1
	}
	r := getHeight(node.Right)
	if r == -1 {
		return -1
	}
	if math.Abs(float64(l-r)) > 1 {
		return -1
	}
	return max(l, r) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type Example struct {
	root *TreeNode
	ans  bool
}

func main() {
	examples := []Example{
		{&TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}, true},
		{&TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 2}}, false},
		{nil, true},
	}
	for i, e := range examples {
		ans := isBalanced(e.root)
		if reflect.DeepEqual(ans, e.ans) {
			fmt.Println("PASS: CASE", i)
		} else {
			fmt.Println("FAIL: CASE", i)
		}
	}
}
