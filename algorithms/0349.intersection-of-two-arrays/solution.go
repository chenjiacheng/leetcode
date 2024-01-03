package main

import (
	"fmt"
	"reflect"
)

func intersection(nums1 []int, nums2 []int) []int {
	set := make(map[int]bool, 0)
	ans := make([]int, 0)
	for _, v := range nums1 {
		set[v] = true
	}
	for _, v := range nums2 {
		if set[v] {
			ans = append(ans, v)
			delete(set, v) // 为了去重所以要删掉
		}
	}
	return ans
}

type Example struct {
	nums1 []int
	nums2 []int
	ans   []int
}

func main() {
	examples := []Example{
		{[]int{1, 2, 2, 1}, []int{2, 2}, []int{2}},
		{[]int{4, 9, 5}, []int{9, 4, 9, 8, 4}, []int{9, 4}},
	}
	for i, e := range examples {
		ans := intersection(e.nums1, e.nums2)
		if reflect.DeepEqual(ans, e.ans) {
			fmt.Println("PASS: CASE", i)
		} else {
			fmt.Println("FAIL: CASE", i)
		}
	}
}
