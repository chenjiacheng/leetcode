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

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var ans []int
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
			// 取每层的最后一个元素，添加到结果集中
			if i == l-1 {
				ans = append(ans, queue[i].Val)
			}
		}
		queue = queue[l:]
	}
	return ans
}

type Example struct {
	root *TreeNode
	ans  []int
}

func main() {
	examples := []Example{
		{&TreeNode{Val: 1, Left: &TreeNode{Val: 2, Right: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 4}}}, []int{1, 3, 4}},
		{&TreeNode{Val: 1, Right: &TreeNode{Val: 3}}, []int{1, 3}},
	}
	for i, e := range examples {
		ans := rightSideView(e.root)
		if reflect.DeepEqual(ans, e.ans) {
			fmt.Println("PASS: CASE", i)
		} else {
			fmt.Println("FAIL: CASE", i)
		}
	}
}
