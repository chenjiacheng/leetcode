package main

import (
	"fmt"
	"reflect"
)

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	if root == nil {
		return nil
	}
	var ans [][]int
	queue := []*Node{root}
	for len(queue) > 0 {
		l := len(queue)
		tmp := make([]int, 0, l)
		for i := 0; i < l; i++ {
			queue = append(queue, queue[i].Children...)
			tmp = append(tmp, queue[i].Val)
		}
		queue = queue[l:]
		ans = append(ans, tmp)
	}
	return ans
}

type Example struct {
	root *Node
	ans  [][]int
}

func main() {
	examples := []Example{
		{&Node{Val: 1, Children: []*Node{{Val: 3, Children: []*Node{{Val: 5}, {Val: 6}}}, {Val: 4}, {Val: 4}}}, [][]int{{1}, {3, 2, 4}, {5, 6}}},
		{&Node{Val: 1, Children: []*Node{{Val: 2}, {Val: 3, Children: []*Node{{Val: 6}, {Val: 7, Children: []*Node{{Val: 11, Children: []*Node{{Val: 14}}}}}}}, {Val: 4, Children: []*Node{{Val: 8, Children: []*Node{{Val: 12}}}}}, {Val: 5, Children: []*Node{{Val: 9, Children: []*Node{{Val: 13}}}, {Val: 10}}}}}, [][]int{{1}, {2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13}, {14}}},
	}
	for i, e := range examples {
		ans := levelOrder(e.root)
		if reflect.DeepEqual(ans, e.ans) {
			fmt.Println("PASS: CASE", i)
		} else {
			fmt.Println("FAIL: CASE", i)
		}
	}
}
