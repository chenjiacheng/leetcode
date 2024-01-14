package main

import (
	"fmt"
	"reflect"
	"sort"
)

func topKFrequent(nums []int, k int) []int {
	var ans []int
	mp := map[int]int{}
	for _, num := range nums {
		mp[num]++
	}
	for key, _ := range mp {
		ans = append(ans, key)
	}
	// 实现快排
	sort.Slice(ans, func(a, b int) bool {
		return mp[ans[a]] > mp[ans[b]]
	})
	return ans[:k]
}

type Example struct {
	nums []int
	k    int
	ans  []int
}

func main() {
	examples := []Example{
		{[]int{1, 1, 1, 2, 2, 3}, 2, []int{1, 2}},
		{[]int{1}, 1, []int{1}},
	}
	for i, e := range examples {
		ans := topKFrequent(e.nums, e.k)
		if reflect.DeepEqual(ans, e.ans) {
			fmt.Println("PASS: CASE", i)
		} else {
			fmt.Println("FAIL: CASE", i)
		}
	}
}
