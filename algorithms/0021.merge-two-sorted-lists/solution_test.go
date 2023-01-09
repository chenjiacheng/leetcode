package leetcode0021

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		args args
		want *ListNode
	}{
		{
			args: args{
				l1: newListNode([]int{1, 2, 4}),
				l2: newListNode([]int{1, 3, 4}),
			},
			want: newListNode([]int{1, 1, 2, 3, 4, 4}),
		},
		{
			args: args{
				l1: newListNode([]int{}),
				l2: newListNode([]int{}),
			},
			want: newListNode([]int{}),
		},
		{
			args: args{
				l1: newListNode([]int{}),
				l2: newListNode([]int{0}),
			},
			want: newListNode([]int{0}),
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.want, mergeTwoLists(test.args.l1, test.args.l2))
	}
}

func newListNode(arr []int) *ListNode {
	var head ListNode
	var pre ListNode
	for _, num := range arr {
		node := ListNode{Val: num, Next: nil}
		if head.Next == nil {
			head.Next = &node
		}
		if pre.Next == nil {
			pre.Next = &node
		} else {
			pre.Next.Next = &node
			pre = *pre.Next
		}
	}
	return head.Next
}
