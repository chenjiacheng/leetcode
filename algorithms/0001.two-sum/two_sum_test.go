package leetcode0001

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTwoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		args args
		want []int
	}{
		{
			args: args{
				nums:   []int{2, 7, 11, 15},
				target: 9,
			},
			want: []int{0, 1},
		},
		{
			args: args{
				nums:   []int{2, 2, 11, 15},
				target: 4,
			},
			want: []int{0, 1},
		},
		{
			args: args{
				nums:   []int{2},
				target: 0,
			},
			want: []int{},
		},
		{
			args: args{
				nums:   []int{},
				target: 0,
			},
			want: []int{},
		},
		{
			args: args{
				nums:   []int{2, 3, 11, 15},
				target: 17,
			},
			want: []int{0, 3},
		},
		{
			args: args{
				nums:   []int{11, 15, -1},
				target: 10,
			},
			want: []int{0, 2},
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.want, twoSum(test.args.nums, test.args.target))
	}
}
