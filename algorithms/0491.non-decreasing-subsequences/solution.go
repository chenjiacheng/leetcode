package main

import (
	"fmt"
	"reflect"
)

var (
	res  [][]int
	path []int
)

func findSubsequences(nums []int) [][]int {
	res, path = make([][]int, 0), make([]int, 0, len(nums))
	dfs(nums, 0)
	return res
}

func dfs(nums []int, start int) {
	if len(path) >= 2 {
		tmp := make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)
	}
	used := make(map[int]bool, len(nums)) // 初始化used字典，用以对同层元素去重
	for i := start; i < len(nums); i++ {
		if used[nums[i]] { // 去重
			continue
		}
		if len(path) == 0 || nums[i] >= path[len(path)-1] {
			path = append(path, nums[i])
			used[nums[i]] = true
			dfs(nums, i+1)
			path = path[:len(path)-1]
		}
	}
}

type Example struct {
	nums []int
	ans  [][]int
}

func main() {
	examples := []Example{
		{[]int{4, 6, 7, 7}, [][]int{{4, 6}, {4, 6, 7}, {4, 6, 7, 7}, {4, 7}, {4, 7, 7}, {6, 7}, {6, 7, 7}, {7, 7}}},
		{[]int{4, 4, 3, 2, 1}, [][]int{{4, 4}}},
	}
	for i, e := range examples {
		ans := findSubsequences(e.nums)
		if reflect.DeepEqual(ans, e.ans) {
			fmt.Println("PASS: CASE", i)
		} else {
			fmt.Println("FAIL: CASE", i)
		}
	}
}
