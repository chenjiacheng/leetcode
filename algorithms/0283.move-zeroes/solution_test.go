package leetcode0283

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		args args
		want []int
	}{
		{
			args: args{
				nums: []int{0, 1, 0, 3, 12},
			},
			want: []int{1, 3, 12, 0, 0},
		},
		{
			args: args{
				nums: []int{0},
			},
			want: []int{0},
		},
	}

	for _, test := range tests {
		moveZeroes(test.args.nums)
		assert.Equal(t, test.want, test.args.nums)
	}
}
