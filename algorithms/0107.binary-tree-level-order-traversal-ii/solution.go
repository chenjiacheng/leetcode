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

func levelOrderBottom(root *TreeNode) [][]int {
	tmp := levelOrder(root)
	var ans [][]int
	for i := len(tmp) - 1; i >= 0; i-- {
		ans = append(ans, tmp[i])
	}
	return ans
}

// BFS（广度优先搜索）
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var ans [][]int
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		l := len(queue)
		tmp := make([]int, 0, l)
		for i := 0; i < l; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
			tmp = append(tmp, queue[i].Val)
		}
		queue = queue[l:]
		ans = append(ans, tmp)
	}
	return ans
}

type Example struct {
	root *TreeNode
	ans  [][]int
}

func main() {
	examples := []Example{
		{&TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}, [][]int{{15, 7}, {9, 20}, {3}}},
		{&TreeNode{Val: 1}, [][]int{{1}}},
		{nil, nil},
	}
	for i, e := range examples {
		ans := levelOrderBottom(e.root)
		if reflect.DeepEqual(ans, e.ans) {
			fmt.Println("PASS: CASE", i)
		} else {
			fmt.Println("FAIL: CASE", i)
		}
	}
}
