package main

import (
	"fmt"
	"reflect"
)

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	mp := make(map[int]int)
	for _, v1 := range nums1 {
		for _, v2 := range nums2 {
			mp[v1+v2]++
		}
	}

	ans := 0
	for _, v3 := range nums3 {
		for _, v4 := range nums4 {
			ans += mp[-v3-v4]
		}
	}
	return ans
}

type Example struct {
	nums1 []int
	nums2 []int
	nums3 []int
	nums4 []int
	ans   int
}

func main() {
	examples := []Example{
		{[]int{1, 2}, []int{-2, -1}, []int{-1, 2}, []int{0, 2}, 2},
		{[]int{0}, []int{0}, []int{0}, []int{0}, 1},
	}
	for i, e := range examples {
		ans := fourSumCount(e.nums1, e.nums2, e.nums3, e.nums4)
		if reflect.DeepEqual(ans, e.ans) {
			fmt.Println("PASS: CASE", i)
		} else {
			fmt.Println("FAIL: CASE", i)
		}
	}
}
