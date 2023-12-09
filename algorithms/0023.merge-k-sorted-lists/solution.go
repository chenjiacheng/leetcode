package main

import (
	"fmt"
	"reflect"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	length := len(lists)
	if length == 0 {
		return nil
	}

	left, right := 0, length-1
	for right > 0 {
		for left < right {
			// 头尾合并，并移动左和右指针
			lists[left] = mergeTwoLists(lists[left], lists[right])
			left++
			right--
		}
		left = 0
	}

	return lists[0]
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
	lists []*ListNode
	ans   *ListNode
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
		{[]*ListNode{buildList([]int{1, 4, 5}), buildList([]int{1, 3, 4}), buildList([]int{2, 6})}, buildList([]int{1, 1, 2, 3, 4, 4, 5, 6})},
		{[]*ListNode{}, buildList([]int{})},
		{[]*ListNode{buildList([]int{})}, buildList([]int{})},
	}
	for i, e := range examples {
		ans := mergeKLists(e.lists)
		if reflect.DeepEqual(ans, e.ans) {
			fmt.Println("PASS: CASE", i)
		} else {
			fmt.Println("FAIL: CASE", i)
		}
	}
}
