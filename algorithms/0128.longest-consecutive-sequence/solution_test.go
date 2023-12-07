package leetcode0128

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestConsecutive(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{
				nums: []int{100, 4, 200, 1, 3, 2},
			},
			want: 4,
		},
		{
			args: args{
				nums: []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
			},
			want: 9,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.want, longestConsecutive(test.args.nums))
	}
}
