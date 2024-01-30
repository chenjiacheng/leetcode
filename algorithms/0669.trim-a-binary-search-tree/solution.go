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

func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left = trimBST(root.Left, low, high)
	root.Right = trimBST(root.Right, low, high)
	/*找到的值有以下情况：
	1.没找到删的点
	2.叶子结点，左右为空
	3.左不空右空
	4.左空右不空
	5.左右不空*/
	if root.Val < low || root.Val > high {
		if root.Left == nil && root.Right == nil {
			return nil
		} else if root.Left == nil && root.Right != nil {
			return root.Right
		} else if root.Left != nil && root.Right == nil {
			return root.Left
		} else {
			cur := root.Right
			for cur.Left != nil {
				cur = cur.Left
			}
			cur.Left = root.Left
			return root.Right
		}
	}
	return root
}

type Example struct {
	root *TreeNode
	low  int
	high int
	ans  *TreeNode
}

func main() {
	examples := []Example{
		{&TreeNode{Val: 1, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 2}}, 1, 2, &TreeNode{Val: 1, Right: &TreeNode{Val: 2}}},
		{&TreeNode{Val: 3, Left: &TreeNode{Val: 0, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}}}, Right: &TreeNode{Val: 4}}, 1, 3, &TreeNode{Val: 3, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}}}},
	}
	for i, e := range examples {
		ans := trimBST(e.root, e.low, e.high)
		if reflect.DeepEqual(ans, e.ans) {
			fmt.Println("PASS: CASE", i)
		} else {
			fmt.Println("FAIL: CASE", i)
		}
	}
}
