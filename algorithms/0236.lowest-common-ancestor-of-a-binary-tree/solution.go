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

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == p || root == q {
		return root
	}
	leftNode := lowestCommonAncestor(root.Left, p, q)
	rightNode := lowestCommonAncestor(root.Right, p, q)
	if leftNode != nil && rightNode != nil {
		return root
	}
	if leftNode == nil && rightNode != nil {
		return rightNode
	}
	if leftNode != nil && rightNode == nil {
		return leftNode
	}
	return nil
}

type Example struct {
	root *TreeNode
	p    *TreeNode
	q    *TreeNode
	ans  *TreeNode
}

func main() {
	node0 := &TreeNode{Val: 4}
	node1 := &TreeNode{Val: 5, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 7}, Right: node0}}
	node2 := &TreeNode{Val: 1, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 8}}
	node3 := &TreeNode{Val: 3, Left: node1, Right: node2}
	node4 := &TreeNode{Val: 2}
	node5 := &TreeNode{Val: 1, Left: node4}
	examples := []Example{
		{node3, node1, node2, node3},
		{node3, node1, node0, node1},
		{node4, node4, node5, node4},
	}
	for i, e := range examples {
		ans := lowestCommonAncestor(e.root, e.p, e.q)
		if reflect.DeepEqual(ans, e.ans) {
			fmt.Println("PASS: CASE", i)
		} else {
			fmt.Println("FAIL: CASE", i)
		}
	}
}
