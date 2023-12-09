package main

import (
	"fmt"
	"reflect"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	length := getLength(head)

	// 删除头节点
	if length-n == 0 {
		head = head.Next
		return head
	}

	dummy := head
	i := 1
	for dummy != nil {
		if length-i == n {
			dummy.Next = dummy.Next.Next
			break
		}
		dummy = dummy.Next
		i++
	}

	return head
}

func getLength(l *ListNode) int {
	length := 1
	for l.Next != nil {
		length++
		l = l.Next
	}

	return length
}

type Example struct {
	head *ListNode
	n    int
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
		{buildList([]int{1, 2, 3, 4, 5}), 2, buildList([]int{1, 2, 3, 5})},
		{buildList([]int{1}), 1, buildList([]int{})},
		{buildList([]int{1, 2}), 1, buildList([]int{1})},
		{buildList([]int{1, 2}), 2, buildList([]int{2})},
	}
	for i, e := range examples {
		ans := removeNthFromEnd(e.head, e.n)
		if reflect.DeepEqual(ans, e.ans) {
			fmt.Println("PASS: CASE", i)
		} else {
			fmt.Println("FAIL: CASE", i)
		}
	}
}
