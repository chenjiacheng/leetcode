package main

import (
	"fmt"
	"reflect"
	"sort"
)

var (
	path []int
	res  [][]int
)

func subsetsWithDup(nums []int) [][]int {
	res, path = make([][]int, 0), make([]int, 0, len(nums))
	sort.Ints(nums)
	dfs(nums, 0)
	return res
}

func dfs(nums []int, start int) {
	tmp := make([]int, len(path))
	copy(tmp, path)
	res = append(res, tmp)

	for i := start; i < len(nums); i++ {
		if i != start && nums[i] == nums[i-1] {
			continue
		}
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
		{[]int{1, 2, 2}, [][]int{{}, {1}, {1, 2}, {1, 2, 2}, {2}, {2, 2}}},
		{[]int{0}, [][]int{{}, {0}}},
	}
	for i, e := range examples {
		ans := subsetsWithDup(e.nums)
		if reflect.DeepEqual(ans, e.ans) {
			fmt.Println("PASS: CASE", i)
		} else {
			fmt.Println("FAIL: CASE", i)
		}
	}
}
