package main

import (
	"fmt"
	"reflect"
)

func combine(n int, k int) [][]int {
	if n <= 0 || k <= 0 || k > n {
		return [][]int{}
	}
	var c []int
	var res [][]int
	generateCombinations(n, k, 1, c, &res)
	return res
}

func generateCombinations(n, k, start int, c []int, res *[][]int) {
	if len(c) == k {
		b := make([]int, len(c))
		copy(b, c)
		*res = append(*res, b)
		return
	}
	for i := start; i <= n-(k-len(c))+1; i++ {
		c = append(c, i)
		generateCombinations(n, k, i+1, c, res)
		c = c[:len(c)-1]
	}
	return
}

type Example struct {
	n   int
	k   int
	ans [][]int
}

func main() {
	examples := []Example{
		{4, 2, [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}}},
		{1, 1, [][]int{{1}}},
	}
	for i, e := range examples {
		ans := combine(e.n, e.k)
		if reflect.DeepEqual(ans, e.ans) {
			fmt.Println("PASS: CASE", i)
		} else {
			fmt.Println("FAIL: CASE", i)
		}
	}
}
