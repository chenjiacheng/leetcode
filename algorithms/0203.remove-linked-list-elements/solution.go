package main

import (
	"fmt"
	"reflect"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode{Next: head} // 虚拟头结点
	cur := dummyHead
	for cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return dummyHead.Next
}

type Example struct {
	head *ListNode
	val  int
	ans  *ListNode
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
		{buildList([]int{1, 2, 6, 3, 4, 5, 6}), 6, buildList([]int{1, 2, 3, 4, 5})},
		{buildList([]int{}), 1, buildList([]int{})},
		{buildList([]int{7, 7, 7, 7}), 7, buildList([]int{})},
	}
	for i, e := range examples {
		ans := removeElements(e.head, e.val)
		if reflect.DeepEqual(ans, e.ans) {
			fmt.Println("PASS: CASE", i)
		} else {
			fmt.Println("FAIL: CASE", i)
		}
	}
}
