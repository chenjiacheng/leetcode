package leetcode0011

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxArea(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{
				height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			},
			want: 49,
		},
		{
			args: args{
				height: []int{1, 1},
			},
			want: 1,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.want, maxArea(test.args.height))
	}
}
