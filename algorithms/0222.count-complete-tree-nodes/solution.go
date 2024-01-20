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

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// 完全二叉树只有两种情况，情况一：就是满二叉树，情况二：最后一层叶子节点没有满。
	// 对于情况一，可以直接用 2^树深度 - 1 来计算，注意这里根节点深度为1。
	// 对于情况二，分别递归左孩子，和右孩子，递归到某一深度一定会有左孩子或者右孩子为满二叉树，然后依然可以按照情况1来计算。
	left := root.Left
	right := root.Right
	leftDepth, rightDepth := 0, 0
	// 求左子树深度
	for left != nil {
		left = left.Left
		leftDepth++
	}
	// 求右子树深度
	for right != nil {
		right = right.Right
		rightDepth++
	}
	// 如果左子树深度和右子树深度相等，则说明该树是满二叉树
	if leftDepth == rightDepth {
		return (2 << leftDepth) - 1 // 注意(2<<1) 相当于2^2，所以leftDepth初始为0
	}
	return countNodes(root.Left) + countNodes(root.Right) + 1
}

type Example struct {
	root *TreeNode
	ans  int
}

func main() {
	examples := []Example{
		{&TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 6}}}, 6},
		{nil, 0},
		{&TreeNode{Val: 1}, 1},
	}
	for i, e := range examples {
		ans := countNodes(e.root)
		if reflect.DeepEqual(ans, e.ans) {
			fmt.Println("PASS: CASE", i)
		} else {
			fmt.Println("FAIL: CASE", i)
		}
	}
}
