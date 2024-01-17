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

func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return nil
	}
	var ans []float64
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		l := len(queue)
		s := 0
		tmp := make([]int, 0, l)
		for i := 0; i < l; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
			tmp = append(tmp, queue[i].Val)
			s += queue[i].Val
		}
		// 计算每层元素的平均值，添加到结果集中
		ans = append(ans, float64(s)/float64(l))
		queue = queue[l:]
	}
	return ans
}

type Example struct {
	root *TreeNode
	ans  []float64
}

func main() {
	examples := []Example{
		{&TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}, []float64{3.00000, 14.50000, 11.00000}},
		{&TreeNode{Val: 3, Left: &TreeNode{Val: 9, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}, Right: &TreeNode{Val: 20}}, []float64{3.00000, 14.50000, 11.00000}},
	}
	for i, e := range examples {
		ans := averageOfLevels(e.root)
		if reflect.DeepEqual(ans, e.ans) {
			fmt.Println("PASS: CASE", i)
		} else {
			fmt.Println("FAIL: CASE", i)
		}
	}
}
