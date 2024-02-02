package main

import (
	"fmt"
	"reflect"
)

func combinationSum3(k int, n int) [][]int {
	if k == 0 {
		return [][]int{}
	}
	var c []int
	var res [][]int
	findCombinationSum3(k, n, 1, c, &res)
	return res
}

func findCombinationSum3(k, target, index int, c []int, res *[][]int) {
	if target == 0 {
		if len(c) == k {
			b := make([]int, len(c))
			copy(b, c)
			*res = append(*res, b)
		}
	}
	for i := index; i < 10; i++ {
		if target >= i {
			c = append(c, i)
			findCombinationSum3(k, target-i, i+1, c, res)
			c = c[:len(c)-1]
		}
	}
}

type Example struct {
	k   int
	n   int
	ans [][]int
}

func main() {
	examples := []Example{
		{3, 7, [][]int{{1, 2, 4}}},
		{3, 9, [][]int{{1, 2, 6}, {1, 3, 5}, {2, 3, 4}}},
	}
	for i, e := range examples {
		ans := combinationSum3(e.k, e.n)
		if reflect.DeepEqual(ans, e.ans) {
			fmt.Println("PASS: CASE", i)
		} else {
			fmt.Println("FAIL: CASE", i)
		}
	}
}
