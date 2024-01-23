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

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	var ans bool
	var travel func(node *TreeNode, sum int)
	travel = func(node *TreeNode, sum int) {
		if node != nil {
			sum += node.Val
			// 如果是叶子节点，且结果相等就返回结果
			if node.Left == nil && node.Right == nil && sum == targetSum {
				ans = true
				return
			}
			travel(node.Left, sum)
			travel(node.Right, sum)
		}
	}
	travel(root, 0)
	return ans
}

type Example struct {
	root      *TreeNode
	targetSum int
	ans       bool
}

func main() {
	examples := []Example{
		{&TreeNode{Val: 5, Left: &TreeNode{Val: 4, Left: &TreeNode{Val: 11, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 2}}}, Right: &TreeNode{Val: 8, Left: &TreeNode{Val: 13}, Right: &TreeNode{Val: 4, Right: &TreeNode{Val: 1}}}}, 22, true},
		// {&TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}, 5, false},
		// {nil, 0, false},
	}
	for i, e := range examples {
		ans := hasPathSum(e.root, e.targetSum)
		if reflect.DeepEqual(ans, e.ans) {
			fmt.Println("PASS: CASE", i)
		} else {
			fmt.Println("FAIL: CASE", i)
		}
	}
}
