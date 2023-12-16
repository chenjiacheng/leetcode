package main

import (
	"fmt"
	"reflect"
)

func minSubArrayLen(target int, nums []int) int {
	key, sum, ans := 0, 0, 0
	for i := range nums {
		// 如果当前值满足条件，直接返回1
		if nums[i] >= target {
			ans = 1
			break
		}
		sum += nums[i]
		for sum >= target {
			res := i - key + 1
			if res < ans || ans == 0 {
				ans = res
			}
			sum -= nums[key]
			key++
		}
	}
	return ans
}

type Example struct {
	target int
	nums   []int
	ans    int
}

func main() {
	examples := []Example{
		{7, []int{2, 3, 1, 2, 4, 3}, 2},
		{4, []int{1, 4, 4}, 1},
		{11, []int{1, 1, 1, 1, 1, 1, 1, 1}, 0},
	}
	for i, e := range examples {
		ans := minSubArrayLen(e.target, e.nums)
		if reflect.DeepEqual(ans, e.ans) {
			fmt.Println("PASS: CASE", i)
		} else {
			fmt.Println("FAIL: CASE", i)
		}
	}
}
