package leetcode0020

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		args args
		want bool
	}{
		{
			args: args{
				s: "()",
			},
			want: true,
		},
		{
			args: args{
				s: "()[]{}",
			},
			want: true,
		},
		{
			args: args{
				s: "(]",
			},
			want: false,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.want, isValid(test.args.s))
	}
}
