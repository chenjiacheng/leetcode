package main

import (
	"fmt"
	"reflect"
)

// 左闭右闭写法
/*func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		middle := (left + right) / 2
		if nums[middle] < target {
			left = middle + 1
		} else if nums[middle] > target {
			right = middle - 1
		} else {
			return middle
		}
	}
	return -1
}*/

// 左闭右开写法
func search(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		middle := (left + right) / 2
		if nums[middle] < target {
			left = middle + 1
		} else if nums[middle] > target {
			right = middle
		} else {
			return middle
		}
	}
	return -1
}

type Example struct {
	nums   []int
	target int
	ans    int
}

func main() {
	examples := []Example{
		{[]int{-1, 0, 3, 5, 9, 12}, 9, 4},
		{[]int{-1, 0, 3, 5, 9, 12}, 2, -1},
	}
	for i, e := range examples {
		ans := search(e.nums, e.target)
		if reflect.DeepEqual(ans, e.ans) {
			fmt.Println("PASS: CASE", i)
		} else {
			fmt.Println("FAIL: CASE", i)
		}
	}
}
