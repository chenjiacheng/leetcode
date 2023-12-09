package main

import (
	"fmt"
	"reflect"
)

func isValid(s string) bool {
	if s == "" {
		return true
	}

	if len(s)%2 == 1 {
		return false
	}

	pairs := map[byte]byte{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	var stack []byte

	for _, v := range s {
		switch v {
		// 如果是左边就入栈
		case '(', '{', '[':
			stack = append(stack, byte(v))
		// 如果是右边则从栈中取出一个做对比
		case ')', '}', ']':
			if len(stack) == 0 || pairs[stack[len(stack)-1]] != byte(v) {
				return false
			}
			stack = stack[:len(stack)-1]
		default:
			return false
		}
	}

	if len(stack) > 0 {
		return false
	}

	return true
}

type Example struct {
	s   string
	ans bool
}

func main() {
	examples := []Example{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
	}
	for i, e := range examples {
		ans := isValid(e.s)
		if reflect.DeepEqual(ans, e.ans) {
			fmt.Println("PASS: CASE", i)
		} else {
			fmt.Println("FAIL: CASE", i)
		}
	}
}
