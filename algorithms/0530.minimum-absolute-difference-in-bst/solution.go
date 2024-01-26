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

func getMinimumDifference(root *TreeNode) int {
	ans := math.MaxInt64
	var prev *TreeNode // 保存上一个指针
	var travel func(node *TreeNode)
	travel = func(node *TreeNode) {
		if node == nil {
			return
		}
		travel(node.Left)
		if prev != nil && node.Val-prev.Val < ans {
			ans = node.Val - prev.Val
		}
		prev = node
		travel(node.Right)
	}
	travel(root)
	return ans
}

type Example struct {
	root *TreeNode
	ans  int
}

func main() {
	examples := []Example{
		{&TreeNode{Val: 4, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 6}}, 1},
		{&TreeNode{Val: 1, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 48, Left: &TreeNode{Val: 12}, Right: &TreeNode{Val: 49}}}, 1},
	}
	for i, e := range examples {
		ans := getMinimumDifference(e.root)
		if reflect.DeepEqual(ans, e.ans) {
			fmt.Println("PASS: CASE", i)
		} else {
			fmt.Println("FAIL: CASE", i)
		}
	}
}
