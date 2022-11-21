package leetcode0019

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	type args struct {
		head []int
		n    int
	}
	tests := []struct {
		args args
		want []int
	}{
		{
			args: args{
				head: []int{1, 2, 3, 4, 5},
				n:    2,
			},
			want: []int{1, 2, 3, 5},
		},
		{
			args: args{
				head: []int{1},
				n:    1,
			},
			want: []int{},
		},
		{
			args: args{
				head: []int{1, 2},
				n:    1,
			},
			want: []int{1},
		},
		{
			args: args{
				head: []int{1, 2},
				n:    2,
			},
			want: []int{2},
		},
	}

	for _, test := range tests {
		assert.Equal(t, newListNode(test.want), removeNthFromEnd(newListNode(test.args.head), test.args.n))
		// removeNthFromEnd(newListNode(test.args.head), test.args.n)
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
