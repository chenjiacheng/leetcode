package main

import (
	"fmt"
	"reflect"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	dummyHead := &ListNode{Next: head}
	cur := dummyHead
	for cur.Next != nil && cur.Next.Next != nil {
		tmp := cur.Next
		tmp1 := cur.Next.Next.Next

		cur.Next = cur.Next.Next
		cur.Next.Next = tmp
		cur.Next.Next.Next = tmp1

		cur = cur.Next.Next
	}
	return dummyHead.Next
}

type Example struct {
	head *ListNode
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
		{buildList([]int{1, 2, 3, 4}), buildList([]int{2, 1, 4, 3})},
		{buildList([]int{}), buildList([]int{})},
		{buildList([]int{1}), buildList([]int{1})},
	}
	for i, e := range examples {
		ans := swapPairs(e.head)
		if reflect.DeepEqual(ans, e.ans) {
			fmt.Println("PASS: CASE", i)
		} else {
			fmt.Println("FAIL: CASE", i)
		}
	}
}
