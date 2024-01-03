package main

import (
	"fmt"
	"reflect"
)

func isAnagram(s string, t string) bool {
	var r [26]int
	for _, v := range s {
		r[v-'a']++
	}
	for _, v := range t {
		r[v-'a']--
	}
	return r == [26]int{}
}

type Example struct {
	s   string
	t   string
	ans bool
}

func main() {
	examples := []Example{
		{"anagram", "nagaram", true},
		{"rat", "car", false},
	}
	for i, e := range examples {
		ans := isAnagram(e.s, e.t)
		if reflect.DeepEqual(ans, e.ans) {
			fmt.Println("PASS: CASE", i)
		} else {
			fmt.Println("FAIL: CASE", i)
		}
	}
}
