package main

import (
	"fmt"
	"reflect"
	"sort"
)

func threeSum(nums []int) [][]int {
	var ans [][]int

	// 使用 sort 包排序
	sort.Ints(nums)

	for i := range nums {

		// nums[i] 去重处理
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, len(nums)-1

		for right > left {
			sum := nums[i] + nums[left] + nums[right]
			if sum > 0 {
				right--
			} else if sum < 0 {
				left++
			} else {
				// nums[left] 和 nums[right] 去重处理
				if !(left > i+1 && nums[left] == nums[left-1]) {
					ans = append(ans, []int{nums[i], nums[left], nums[right]})
				}
				right--
				left++
			}
		}
	}

	return ans
}

type Example struct {
	nums []int
	ans  [][]int
}

func main() {
	examples := []Example{
		{[]int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{nums: []int{0, 1, 1}},
		{[]int{0, 0, 0}, [][]int{{0, 0, 0}}},
	}
	for i, e := range examples {
		ans := threeSum(e.nums)
		if reflect.DeepEqual(ans, e.ans) {
			fmt.Println("PASS: CASE", i)
		} else {
			fmt.Println("FAIL: CASE", i)
		}
	}
}
