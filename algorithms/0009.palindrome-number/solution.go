package main

import (
	"fmt"
	"reflect"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	num := x
	cur := 0
	for x > 0 {
		cur = cur*10 + x%10
		x /= 10
	}
	return num == cur
}

type Example struct {
	x   int
	ans bool
}

func main() {
	examples := []Example{
		{121, true},
		{-121, false},
		{10, false},
	}
	for i, e := range examples {
		ans := isPalindrome(e.x)
		if reflect.DeepEqual(ans, e.ans) {
			fmt.Println("PASS: CASE", i)
		} else {
			fmt.Println("FAIL: CASE", i)
		}
	}
}
