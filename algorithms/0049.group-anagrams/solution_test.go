package leetcode0049

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		args args
		want [][]string
	}{
		{
			args: args{
				strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			},
			want: [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}},
		},
		{
			args: args{
				strs: []string{""},
			},
			want: [][]string{{""}},
		},
		{
			args: args{
				strs: []string{"a"},
			},
			want: [][]string{{"a"}},
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.want, groupAnagrams(test.args.strs))
	}
}
