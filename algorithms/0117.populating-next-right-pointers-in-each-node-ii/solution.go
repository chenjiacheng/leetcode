package main

import (
	"fmt"
	"reflect"
)

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	queue := []*Node{root}
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
			if i < l-1 {
				queue[i].Next = queue[i+1]
			}
		}
		queue = queue[l:]
	}
	return root
}

type Example struct {
	root *Node
	ans  *Node
}

func main() {
	examples := []Example{
		{nil, nil},
	}
	for i, e := range examples {
		ans := connect(e.root)
		if reflect.DeepEqual(ans, e.ans) {
			fmt.Println("PASS: CASE", i)
		} else {
			fmt.Println("FAIL: CASE", i)
		}
	}
}
