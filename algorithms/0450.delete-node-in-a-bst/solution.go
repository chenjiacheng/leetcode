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

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	/*找到的值有以下情况：
	1.没找到删的点
	2.叶子结点，左右为空
	3.左不空右空
	4.左空右不空
	5.左右不空*/
	if root.Val == key {
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
	if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	}
	return root
}

type Example struct {
	root *TreeNode
	key  int
	ans  *TreeNode
}

func main() {
	examples := []Example{
		{&TreeNode{Val: 5, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 6, Right: &TreeNode{Val: 7}}}, 3, &TreeNode{Val: 5, Left: &TreeNode{Val: 4, Left: &TreeNode{Val: 2}}, Right: &TreeNode{Val: 6, Right: &TreeNode{Val: 7}}}},
		{&TreeNode{Val: 5, Left: &TreeNode{Val: 2, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 6, Right: &TreeNode{Val: 7}}}, 0, &TreeNode{Val: 5, Left: &TreeNode{Val: 2, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 6, Right: &TreeNode{Val: 7}}}},
		{nil, 0, nil},
	}
	for i, e := range examples {
		ans := deleteNode(e.root, e.key)
		if reflect.DeepEqual(ans, e.ans) {
			fmt.Println("PASS: CASE", i)
		} else {
			fmt.Println("FAIL: CASE", i)
		}
	}
}
