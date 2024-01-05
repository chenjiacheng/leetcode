package main

import (
	"fmt"
	"reflect"
)

func reverseStr(s string, k int) string {
	b := []byte(s)
	l := len(s)
	for i := 0; i < l; i += 2 * k {
		if i+k <= l {
			reverseString(b[i : i+k])
		} else {
			reverseString(b[i:l])
		}
	}
	return string(b)
}

func reverseString(s []byte) {
	l := len(s)
	for i := 0; i < l; i++ {
		s[i], s[l-1] = s[l-1], s[i]
		l--
	}
}

type Example struct {
	s   string
	k   int
	ans string
}

func main() {
	examples := []Example{
		{"abcdefg", 2, "bacdfeg"},
		{"abcd", 2, "bacd"},
	}
	for i, e := range examples {
		ans := reverseStr(e.s, e.k)
		if reflect.DeepEqual(ans, e.ans) {
			fmt.Println("PASS: CASE", i)
		} else {
			fmt.Println("FAIL: CASE", i)
		}
	}
}
