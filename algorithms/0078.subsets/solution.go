package main

import (
	"fmt"
	"reflect"
)

var (
	path []int
	res  [][]int
)

func subsets(nums []int) [][]int {
	res, path = make([][]int, 0), make([]int, 0, len(nums))
	dfs(nums, 0)
	return res
}
func dfs(nums []int, start int) {
	tmp := make([]int, len(path))
	copy(tmp, path)
	res = append(res, tmp)

	for i := start; i < len(nums); i++ {
		path = append(path, nums[i])
		dfs(nums, i+1)
		path = path[:len(path)-1]
	}
}

type Example struct {
	nums []int
	ans  [][]int
}

func main() {
	examples := []Example{
		{[]int{1, 2, 3}, [][]int{{}, {1}, {1, 2}, {1, 2, 3}, {1, 3}, {2}, {2, 3}, {3}}},
		{[]int{0}, [][]int{{}, {0}}},
	}
	for i, e := range examples {
		ans := subsets(e.nums)
		if reflect.DeepEqual(ans, e.ans) {
			fmt.Println("PASS: CASE", i)
		} else {
			fmt.Println("FAIL: CASE", i)
		}
	}
}
