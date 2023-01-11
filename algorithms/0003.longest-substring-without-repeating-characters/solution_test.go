package leetcode0003

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{
				s: "abcabcbb",
			},
			want: 3,
		},
		{
			args: args{
				s: "bbbbb",
			},
			want: 1,
		},
		{
			args: args{
				s: "pwwkew",
			},
			want: 3,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.want, lengthOfLongestSubstring(test.args.s))
	}
}
