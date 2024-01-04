package main

import (
	"fmt"
	"reflect"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return nil
	}

	var ans [][]int

	// 使用 sort 包排序
	sort.Ints(nums)

	for j := range nums {
		if j > 0 && nums[j] == nums[j-1] {
			continue // 去重
		}

		for i := j + 1; i < len(nums)-2; i++ {
			if i > j+1 && nums[i] == nums[i-1] {
				continue // 去重
			}
			left, right := i+1, len(nums)-1

			for right > left {
				sum := nums[j] + nums[i] + nums[left] + nums[right]
				if sum > target {
					right--
				} else if sum < target {
					left++
				} else {
					ans = append(ans, []int{nums[j], nums[i], nums[left], nums[right]})
					for left < right && nums[left] == nums[left+1] {
						left++ // 去重
					}
					for left < right && nums[right] == nums[right-1] {
						right-- // 去重
					}
					left++
					right--
				}
			}
		}
	}

	return ans
}

type Example struct {
	nums   []int
	target int
	ans    [][]int
}

func main() {
	examples := []Example{
		{[]int{1, 0, -1, 0, -2, 2}, 0, [][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}}},
		{[]int{2, 2, 2, 2, 2}, 8, [][]int{{2, 2, 2, 2}}},
	}
	for i, e := range examples {
		ans := fourSum(e.nums, e.target)
		if reflect.DeepEqual(ans, e.ans) {
			fmt.Println("PASS: CASE", i)
		} else {
			fmt.Println("FAIL: CASE", i)
		}
	}
}
