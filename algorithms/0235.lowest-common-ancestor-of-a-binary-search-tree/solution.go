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
	if root.Val > p.Val && root.Val > q.Val {
		leftNode := lowestCommonAncestor(root.Left, p, q)
		if leftNode != nil {
			return leftNode
		}
	}
	if root.Val < p.Val && root.Val < q.Val {
		rightNode := lowestCommonAncestor(root.Right, p, q)
		if rightNode != nil {
			return rightNode
		}
	}
	return root
}

type Example struct {
	root *TreeNode
	p    *TreeNode
	q    *TreeNode
	ans  *TreeNode
}

func main() {
	node0 := &TreeNode{Val: 4, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 5}}
	node1 := &TreeNode{Val: 2, Left: &TreeNode{Val: 0}, Right: node0}
	node2 := &TreeNode{Val: 8, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 9}}
	node3 := &TreeNode{Val: 6, Left: node1, Right: node2}
	examples := []Example{
		{node3, node1, node2, node3},
		{node3, node1, node0, node1},
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
