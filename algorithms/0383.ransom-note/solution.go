package main

import (
	"fmt"
	"reflect"
)

func canConstruct(ransomNote string, magazine string) bool {
	var r [26]int
	for _, v := range magazine {
		r[v-'a']++
	}
	for _, v := range ransomNote {
		if r[v-'a'] == 0 {
			return false
		} else {
			r[v-'a']--
		}
	}
	return true
}

type Example struct {
	ransomNote string
	magazine   string
	ans        bool
}

func main() {
	examples := []Example{
		{"a", "b", false},
		{"aa", "ab", false},
		{"aa", "aab", true},
	}
	for i, e := range examples {
		ans := canConstruct(e.ransomNote, e.magazine)
		if reflect.DeepEqual(ans, e.ans) {
			fmt.Println("PASS: CASE", i)
		} else {
			fmt.Println("FAIL: CASE", i)
		}
	}
}
