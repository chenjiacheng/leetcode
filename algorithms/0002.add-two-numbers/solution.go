package main

import (
	"fmt"
	"reflect"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var r, p *ListNode
	carry := 0
	for l1 != nil || l2 != nil {
		sum := 0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		sum += carry
		sum, carry = sum%10, sum/10
		if r == nil {
			r = &ListNode{Val: sum}
			p = r
		} else {
			p.Next = &ListNode{Val: sum}
			p = p.Next
		}
	}

	if carry > 0 {
		p.Next = &ListNode{Val: carry}
	}

	return r
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
		{buildList([]int{2, 4, 3}), buildList([]int{5, 6, 4}), buildList([]int{7, 0, 8})},
		{buildList([]int{0}), buildList([]int{0}), buildList([]int{0})},
		{buildList([]int{9, 9, 9, 9, 9, 9, 9}), buildList([]int{9, 9, 9, 9}), buildList([]int{8, 9, 9, 9, 0, 0, 0, 1})},
	}
	for i, e := range examples {
		ans := addTwoNumbers(e.l1, e.l2)
		if reflect.DeepEqual(ans, e.ans) {
			fmt.Println("PASS: CASE", i)
		} else {
			fmt.Println("FAIL: CASE", i)
		}
	}
}
