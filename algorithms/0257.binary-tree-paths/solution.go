package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	ans := make([]string, 0)
	var travel func(node *TreeNode, path string)
	travel = func(node *TreeNode, path string) {
		if node != nil {
			path += strconv.Itoa(node.Val)
			// 如果是叶子节点，就记录结果
			if node.Left == nil && node.Right == nil {
				ans = append(ans, path)
			}
			path += "->"
			travel(node.Left, path)
			travel(node.Right, path)
		}
	}
	travel(root, "")
	return ans
}

type Example struct {
	root *TreeNode
	ans  []string
}

func main() {
	examples := []Example{
		{&TreeNode{Val: 1, Left: &TreeNode{Val: 2, Right: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3}}, []string{"1->2->5", "1->3"}},
		{&TreeNode{Val: 1}, []string{"1"}},
	}
	for i, e := range examples {
		ans := binaryTreePaths(e.root)
		if reflect.DeepEqual(ans, e.ans) {
			fmt.Println("PASS: CASE", i)
		} else {
			fmt.Println("FAIL: CASE", i)
		}
	}
}
