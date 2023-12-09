package main

import (
	"fmt"
	"reflect"
)

func longestConsecutive(nums []int) int {
	// 把值存储到hash表
	mp := map[int]bool{}
	for _, num := range nums {
		mp[num] = true
	}

	ans := 0 // 最长长度
	for num := range mp {
		// 如果是起点就计算长度
		if _, ok := mp[num-1]; !ok {
			x := num + 1
			for mp[x] {
				x++
			}
			if x-num > ans {
				ans = x - num
			}
		}
	}

	return ans
}

type Example struct {
	nums []int
	ans  int
}

func main() {
	examples := []Example{
		{[]int{100, 4, 200, 1, 3, 2}, 4},
		{[]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}, 9},
	}
	for i, e := range examples {
		ans := longestConsecutive(e.nums)
		if reflect.DeepEqual(ans, e.ans) {
			fmt.Println("PASS: CASE", i)
		} else {
			fmt.Println("FAIL: CASE", i)
		}
	}
}
