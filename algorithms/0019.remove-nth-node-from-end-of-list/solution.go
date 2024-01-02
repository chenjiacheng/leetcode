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
	dummyHead := &ListNode{Next: head}
	slow := dummyHead // 慢指针
	fast := dummyHead // 快指针

	// 快指针先走 n 步
	for n > 0 && fast != nil {
		fast = fast.Next
		n--
	}

	fast = fast.Next // 最后多走一步，为了实现慢指针停留上前一位
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next

	return dummyHead.Next
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
