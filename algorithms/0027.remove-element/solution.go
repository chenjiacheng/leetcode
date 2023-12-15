package main

import (
	"fmt"
	"reflect"
)

func removeElement(nums []int, val int) int {
	k := 0
	for i := range nums {
		if nums[i] != val {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}

type Example struct {
	nums []int
	val  int
	ans  int
}

func main() {
	examples := []Example{
		{[]int{3, 2, 2, 3}, 3, 2},
		{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2, 5},
	}
	for i, e := range examples {
		ans := removeElement(e.nums, e.val)
		if reflect.DeepEqual(ans, e.ans) {
			fmt.Println("PASS: CASE", i)
		} else {
			fmt.Println("FAIL: CASE", i)
		}
	}
}
