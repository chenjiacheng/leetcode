package main

import (
	"fmt"
	"reflect"
)

func partition(s string) [][]string {
	if len(s) == 0 {
		return [][]string{}
	}
	var c []string
	var res [][]string
	findPalindrome(s, 0, c, &res)
	return res
}

func findPalindrome(s string, index int, c []string, res *[][]string) {
	if index == len(s) {
		b := make([]string, len(c))
		copy(b, c)
		*res = append(*res, b)
		return
	}
	for i := index; i < len(s); i++ {
		str := s[index : i+1]
		if isPalindrome(str) {
			c = append(c, str)
			findPalindrome(s, i+1, c, res)
			c = c[:len(c)-1]
		}
	}
}

func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

type Example struct {
	s   string
	ans [][]string
}

func main() {
	examples := []Example{
		{"aab", [][]string{{"a", "a", "b"}, {"aa", "b"}}},
		{"a", [][]string{{"a"}}},
	}
	for i, e := range examples {
		ans := partition(e.s)
		if reflect.DeepEqual(ans, e.ans) {
			fmt.Println("PASS: CASE", i)
		} else {
			fmt.Println("FAIL: CASE", i)
		}
	}
}
