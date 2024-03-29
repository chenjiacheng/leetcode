package main

import (
	"fmt"
	"reflect"
)

func maxArea(height []int) int {
	left, right, ans := 0, len(height)-1, 0

	for left < right {

		res := 0

		// 当右边大于左边，需要移动左边，反之移动右边
		if height[right] > height[left] {
			res = height[left] * (right - left)
			left++
		} else {
			res = height[right] * (right - left)
			right--
		}

		if res > ans {
			ans = res
		}
	}

	return ans
}

type Example struct {
	height []int
	ans    int
}

func main() {
	examples := []Example{
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{[]int{1, 1}, 1},
	}
	for i, e := range examples {
		ans := maxArea(e.height)
		if reflect.DeepEqual(ans, e.ans) {
			fmt.Println("PASS: CASE", i)
		} else {
			fmt.Println("FAIL: CASE", i)
		}
	}
}
