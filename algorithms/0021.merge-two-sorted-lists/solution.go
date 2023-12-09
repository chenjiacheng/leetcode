package main

import (
	"fmt"
	"reflect"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	var node *ListNode

	if l1.Val >= l2.Val {
		node = l2
		node.Next = mergeTwoLists(l1, l2.Next)
	} else {
		node = l1
		node.Next = mergeTwoLists(l1.Next, l2)
	}

	return node
}

type Example struct {
	l1  *ListNode
	l2  *ListNode
	ans *ListNode
}

func buildList(nums []int) *ListNode {
	dummy := &ListNode{Val: 0}
	cur := dummy
	for _, num := range nums {
		cur.Next = &ListNode{Val: num}
		cur = cur.Next
	}
	return dummy.Next
}

func main() {
	examples := []Example{
		{buildList([]int{1, 2, 4}), buildList([]int{1, 3, 4}), buildList([]int{1, 1, 2, 3, 4, 4})},
		{buildList([]int{}), buildList([]int{}), buildList([]int{})},
		{buildList([]int{}), buildList([]int{0}), buildList([]int{0})},
	}
	for i, e := range examples {
		ans := mergeTwoLists(e.l1, e.l2)
		if reflect.DeepEqual(ans, e.ans) {
			fmt.Println("PASS: CASE", i)
		} else {
			fmt.Println("FAIL: CASE", i)
		}
	}
}
