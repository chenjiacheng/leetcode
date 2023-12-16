package main

import (
	"fmt"
	"reflect"
)

func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	// 定义边界
	up, right, down, left := 0, n-1, m-1, 0
	var ans []int
	for {
		// 遍历上边
		for i := left; i <= right; i++ {
			ans = append(ans, matrix[up][i])
		}
		up++
		if up > down {
			break
		}

		// 遍历右边
		for i := up; i <= down; i++ {
			ans = append(ans, matrix[i][right])
		}
		right--
		if right < left {
			break
		}

		// 遍历下边
		for i := right; i >= left; i-- {
			ans = append(ans, matrix[down][i])
		}
		down--
		if down < up {
			break
		}

		// 遍历左边
		for i := down; i >= up; i-- {
			ans = append(ans, matrix[i][left])
		}
		left++
		if left > right {
			break
		}
	}
	return ans
}

type Example struct {
	matrix [][]int
	ans    []int
}

func main() {
	examples := []Example{
		{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, []int{1, 2, 3, 6, 9, 8, 7, 4, 5}},
		{[][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}, []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}},
	}
	for i, e := range examples {
		ans := spiralOrder(e.matrix)
		if reflect.DeepEqual(ans, e.ans) {
			fmt.Println("PASS: CASE", i)
		} else {
			fmt.Println("FAIL: CASE", i)
		}
	}
}
