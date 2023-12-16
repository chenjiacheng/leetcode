package main

import (
	"fmt"
	"reflect"
)

func generateMatrix(n int) [][]int {
	ans := make([][]int, n)
	for i := range ans {
		ans[i] = make([]int, n)
	}
	// 定义边界
	up, right, down, left := 0, n-1, n-1, 0
	num := 1
	for {
		// 遍历上边
		for i := left; i <= right; i++ {
			ans[up][i] = num
			num++
		}
		up++
		if up > down {
			break
		}

		// 遍历右边
		for i := up; i <= down; i++ {
			ans[i][right] = num
			num++
		}
		right--
		if right < left {
			break
		}

		// 遍历下边
		for i := right; i >= left; i-- {
			ans[down][i] = num
			num++
		}
		down--
		if down < up {
			break
		}

		// 遍历左边
		for i := down; i >= up; i-- {
			ans[i][left] = num
			num++
		}
		left++
		if left > right {
			break
		}
	}
	return ans
}

type Example struct {
	n   int
	ans [][]int
}

func main() {
	examples := []Example{
		{3, [][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}}},
		{1, [][]int{{1}}},
	}
	for i, e := range examples {
		ans := generateMatrix(e.n)
		if reflect.DeepEqual(ans, e.ans) {
			fmt.Println("PASS: CASE", i)
		} else {
			fmt.Println("FAIL: CASE", i)
		}
	}
}
