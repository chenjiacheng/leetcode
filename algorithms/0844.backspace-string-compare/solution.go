package main

import (
	"fmt"
	"reflect"
)

// 利用栈方式实现
func backspaceCompare(s string, t string) bool {
	var sStack []byte
	for i := range s {
		if s[i] != '#' {
			sStack = append(sStack, s[i])
		} else if len(sStack) > 0 {
			sStack = sStack[:len(sStack)-1]
		}
	}

	var tStack []byte
	for i := range t {
		if t[i] != '#' {
			tStack = append(tStack, t[i])
		} else if len(tStack) > 0 {
			tStack = tStack[:len(tStack)-1]
		}
	}

	return string(sStack) == string(tStack)
}

type Example struct {
	s   string
	t   string
	ans bool
}

func main() {
	examples := []Example{
		{"ab#c", "ad#c", true},
		{"ab##", "c#d#", true},
		{"a#c", "b", false},
	}
	for i, e := range examples {
		ans := backspaceCompare(e.s, e.t)
		if reflect.DeepEqual(ans, e.ans) {
			fmt.Println("PASS: CASE", i)
		} else {
			fmt.Println("FAIL: CASE", i)
		}
	}
}
