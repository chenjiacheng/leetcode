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

func findMode(root *TreeNode) []int {
	ans := make([]int, 0)
	count := 1
	max := 1
	var prev *TreeNode
	var travel func(node *TreeNode)
	travel = func(node *TreeNode) {
		if node == nil {
			return
		}
		travel(node.Left)
		if prev != nil && prev.Val == node.Val {
			count++
		} else {
			count = 1
		}
		if count >= max {
			if count > max && len(ans) > 0 {
				ans = []int{node.Val}
			} else {
				ans = append(ans, node.Val)
			}
			max = count
		}
		prev = node
		travel(node.Right)
	}
	travel(root)
	return ans
}

type Example struct {
	root *TreeNode
	ans  []int
}

func main() {
	examples := []Example{
		{&TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 2}}}, []int{2}},
		{&TreeNode{Val: 0}, []int{0}},
	}
	for i, e := range examples {
		ans := findMode(e.root)
		if reflect.DeepEqual(ans, e.ans) {
			fmt.Println("PASS: CASE", i)
		} else {
			fmt.Println("FAIL: CASE", i)
		}
	}
}
