package main

import (
	"fmt"
	"reflect"
)

func removeDuplicates(nums []int) int {
	k := 0
	for i := range nums {
		if i == 0 || nums[i] != nums[i-1] {
			nums[k] = nums[i]
			k++
		}
	}

	return k
}

type Example struct {
	nums []int
	ans  int
}

func main() {
	examples := []Example{
		{[]int{1, 1, 2}, 2},
		{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5},
	}
	for i, e := range examples {
		ans := removeDuplicates(e.nums)
		if reflect.DeepEqual(ans, e.ans) {
			fmt.Println("PASS: CASE", i)
		} else {
			fmt.Println("FAIL: CASE", i)
		}
	}
}
