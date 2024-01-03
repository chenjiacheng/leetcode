package main

import (
	"fmt"
	"reflect"
)

func isHappy(n int) bool {
	r := map[int]bool{}
	for !r[n] {
		n, r[n] = step(n), true
	}
	return n == 1
}

func step(n int) int {
	sum := 0
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n /= 10
	}
	return sum
}

type Example struct {
	n   int
	ans bool
}

func main() {
	examples := []Example{
		{19, true},
		{2, false},
	}
	for i, e := range examples {
		ans := isHappy(e.n)
		if reflect.DeepEqual(ans, e.ans) {
			fmt.Println("PASS: CASE", i)
		} else {
			fmt.Println("FAIL: CASE", i)
		}
	}
}
