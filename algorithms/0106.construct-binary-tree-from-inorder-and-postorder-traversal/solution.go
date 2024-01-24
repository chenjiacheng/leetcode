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

func buildTree(inorder []int, postorder []int) *TreeNode {
	pl := len(postorder)
	if pl == 0 {
		return nil
	}
	rootVal := postorder[pl-1]
	ans := &TreeNode{Val: rootVal}
	if pl == 1 {
		return ans
	}
	index := 0
	for ; index < len(inorder); index++ {
		if inorder[index] == rootVal {
			break
		}
	}
	ans.Left = buildTree(inorder[0:index], postorder[0:index])
	ans.Right = buildTree(inorder[index+1:], postorder[index:pl-1])
	return ans
}

type Example struct {
	inorder   []int
	postorder []int
	ans       *TreeNode
}

func main() {
	examples := []Example{
		{[]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3}, &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}},
		{[]int{-1}, []int{-1}, &TreeNode{Val: -1}},
	}
	for i, e := range examples {
		ans := buildTree(e.inorder, e.postorder)
		if reflect.DeepEqual(ans, e.ans) {
			fmt.Println("PASS: CASE", i)
		} else {
			fmt.Println("FAIL: CASE", i)
		}
	}
}
