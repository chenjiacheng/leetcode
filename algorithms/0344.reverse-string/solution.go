package main

import (
	"fmt"
	"reflect"
)

func reverseString(s []byte) {
	l := len(s)
	for i := 0; i < l; i++ {
		s[i], s[l-1] = s[l-1], s[i]
		l--
	}
}

type Example struct {
	s   []byte
	ans []byte
}

func main() {
	examples := []Example{
		{[]byte{'h', 'e', 'l', 'l', 'o'}, []byte{'o', 'l', 'l', 'e', 'h'}},
		{[]byte{'H', 'a', 'n', 'n', 'a', 'h'}, []byte{'h', 'a', 'n', 'n', 'a', 'H'}},
	}
	for i, e := range examples {
		reverseString(e.s)
		if reflect.DeepEqual(e.s, e.ans) {
			fmt.Println("PASS: CASE", i)
		} else {
			fmt.Println("FAIL: CASE", i)
		}
	}
}
