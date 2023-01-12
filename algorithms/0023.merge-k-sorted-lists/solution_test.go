package leetcode0023

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	type args struct {
		lists []*ListNode
	}
	tests := []struct {
		args args
		want *ListNode
	}{
		{
			args: args{
				lists: []*ListNode{
					newListNode([]int{1, 4, 5}),
					newListNode([]int{1, 3, 4}),
					newListNode([]int{2, 6}),
				},
			},
			want: newListNode([]int{1, 1, 2, 3, 4, 4, 5, 6}),
		},
		{
			args: args{
				lists: []*ListNode{},
			},
			want: newListNode([]int{}),
		},
		{
			args: args{
				lists: []*ListNode{newListNode([]int{})},
			},
			want: newListNode([]int{}),
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.want, mergeKLists(test.args.lists))
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
