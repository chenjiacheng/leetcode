package main

import (
	"fmt"
	"reflect"
)

func twoSum(nums []int, target int) []int {
	myMap := make(map[int]int)
	for i, num := range nums {
		res, ok := myMap[target-num]
		if ok {
			return []int{res, i}
		} else {
			myMap[num] = i
		}
	}
	return []int{}
}

type Example struct {
	nums   []int
	target int
	ans    []int
}

func main() {
	examples := []Example{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
	}
	for i, e := range examples {
		ans := twoSum(e.nums, e.target)
		if reflect.DeepEqual(ans, e.ans) {
			fmt.Println("PASS: CASE", i)
		} else {
			fmt.Println("FAIL: CASE", i)
		}
	}
}
