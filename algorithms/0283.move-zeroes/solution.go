package main

import (
	"fmt"
	"reflect"
)

func moveZeroes(nums []int) {
	left := 0
	for i := range nums {
		if nums[i] != 0 {
			nums[i], nums[left] = nums[left], nums[i]
			left++
		}
	}
}

type Example struct {
	nums []int
	ans  []int
}

func main() {
	examples := []Example{
		{[]int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
		{[]int{0}, []int{0}},
	}
	for i, e := range examples {
		moveZeroes(e.nums)
		if reflect.DeepEqual(e.nums, e.ans) {
			fmt.Println("PASS: CASE", i)
		} else {
			fmt.Println("FAIL: CASE", i)
		}
	}
}
