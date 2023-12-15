package main

import (
	"fmt"
	"reflect"
)

func sortedSquares(nums []int) []int {
	n := len(nums)
	l, r := 0, n-1
	ans := make([]int, n)
	for k := n - 1; k >= 0; k-- {
		if v1, v2 := nums[l]*nums[l], nums[r]*nums[r]; v1 > v2 {
			ans[k] = v1
			l++
		} else {
			ans[k] = v2
			r--
		}
	}
	return ans
}

type Example struct {
	nums []int
	ans  []int
}

func main() {
	examples := []Example{
		{[]int{-4, -1, 0, 3, 10}, []int{0, 1, 9, 16, 100}},
		{[]int{-7, -3, 2, 3, 11}, []int{4, 9, 9, 49, 121}},
		{[]int{-5, -3, -2, -1}, []int{1, 4, 9, 25}},
	}
	for i, e := range examples {
		ans := sortedSquares(e.nums)
		if reflect.DeepEqual(ans, e.ans) {
			fmt.Println("PASS: CASE", i)
		} else {
			fmt.Println("FAIL: CASE", i)
		}
	}
}
